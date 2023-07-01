import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutput
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


class ExceptionEntry(
        owner: AttributeOwner,
        spc: Int,
        epc: Int,
        hpc: Int,
        c: ConstClassInfo) {
    def this(owner: AttributeOwner) = this(owner, 0, 0, 0, null)

    var _owner = owner

    var startPc = spc  // inclusive
    var endPc = epc  // exclusive
    var handlerPc = hpc
    var classType: ConstClassInfo = c  // null means catch all

    // only used for deserialization
    var _tmpSection: CodeScope = null

    def className(): String = {
        if (classType == null) {
            return null
        }
        return classType.className()
    }

    def serialize(output: DataOutput) {
        output.writeShort(startPc)
        output.writeShort(endPc)
        output.writeShort(handlerPc)
        if (classType == null) {
            output.writeShort(0)
        } else {
            output.writeShort(classType.index)
        }
    }

    def deserialize(input: DataInputStream) {
        startPc = input.readUnsignedShort()
        endPc = input.readUnsignedShort()
        handlerPc = input.readUnsignedShort()

        val index = input.readUnsignedShort()
        if (index == 0) {
            classType = null
        } else {
            classType = _owner.constants().getClassByIndex(index)
        }
    }

    def debugString(indent: String): String = {
        var name = "(any)"
        if (classType != null) {
            name = classType.className()
        }

        var result = indent + "[" + startPc + ", " + endPc + ") -> " + handlerPc
        result += " for " + name + "\n"
        return result
    }
}

class CodeAttribute(o: AttributeOwner)
        extends Attribute(o, "Code")
        with AttributeOwner {

    var maxStack = 0
    var maxLocals = 0

    var code = new CodeScope(this, null, 0)

    var attributes = new CodeAttributes(this)

    def constants(): ConstantPool = _owner.constants()

    def serialize(output: DataOutputStream) {
        var buffer = new ByteArrayOutputStream()
        var codeWriter = new DataOutputStream(buffer)

        codeWriter.writeShort(maxStack)  // TODO compute max stack
        codeWriter.writeShort(maxLocals)  // TODO compute max locals

        code.serialize(codeWriter)

        attributes.lineNumberTable = code.generateLineNumberTable()

        attributes.serialize(codeWriter)

        // finally write the real result

        val bytes = buffer.toByteArray()

        output.writeShort(_name.index)
        output.writeInt(bytes.length)
        output.write(bytes)
    }

    def _populateLineNumber(operations: Vector[Operation]) {
        if (attributes.lineNumberTable != null) {
            val table = attributes.lineNumberTable.table
            for (op <- operations) {
                val entry = table.floorEntry(op.pc)
                if (entry != null) {
                    op.line = entry.getValue()
                }
            }
        }
    }

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream) {
        maxStack = input.readUnsignedShort()
        maxLocals = input.readUnsignedShort()

        val codeLength = input.readInt()
        var codeBytes = new Array[Byte](codeLength)
        input.readFully(codeBytes)

        var operations = Operation.deserialize(this, codeBytes)

        var exceptionEntries = new Vector[ExceptionEntry]()
        val numExceptionEntries = input.readUnsignedShort()
        for (_ <- 1 to numExceptionEntries) {
            var entry = new ExceptionEntry(_owner)
            entry.deserialize(input)
            exceptionEntries.add(entry)
        }

        attributes = new CodeAttributes(this)
        attributes.deserialize(input)

        _populateLineNumber(operations)

        code = CodeScopeReconstructor.reconstruct(
                this,
                exceptionEntries,
                operations)
    }

    def debugString(indent: String): String = {
        var result = indent + "Code:\n"
        result += indent + "  Max stack: " + maxStack + "\n"
        result += indent + "  Max locals: " + maxLocals + "\n"
        val subIndent = indent + "    "
        result += code.debugString(subIndent)
        result += indent + "  Exceptions:\n"
        var exceptions = new Vector[ExceptionEntry]()
        code._collectExceptionEntries(exceptions)
        for (entry <- exceptions) {
            result += entry.debugString(subIndent)
        }
        result += indent + "  Attributes:\n"
        result += attributes.debugString(subIndent)

        return result
    }
}
