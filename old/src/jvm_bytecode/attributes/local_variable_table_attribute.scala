import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.TreeMap
import java.util.Vector

import scala.collection.JavaConversions._


class LocalVariableEntry(
        o: AttributeOwner,
        start: Int,
        end: Int,
        name: String,
        descriptor: FieldType,
        at: Int) {
    def this(o: AttributeOwner) = this(o, -1, -1, null, null, -1)
    var _owner = o

    var startPc = start  // inclusive
    var endPc = end  // exclusive

    var _fieldName: ConstUtf8Info = null
    if (name != null) {
        _fieldName = _owner.constants().getUtf8(name)
    }

    var _fieldType: FieldType = descriptor
    var _fieldDescriptor: ConstUtf8Info = null
    if (fieldType != null) {
        _fieldDescriptor = _owner.constants().getUtf8(
                fieldType.descriptorString())
    }

    var index = at

    def fieldName(): String = _fieldName.value()

    def fieldType(): FieldType = _fieldType

    def serialize(output: DataOutputStream) {
        output.writeShort(startPc)
        output.writeShort(endPc - startPc)
        output.writeShort(_fieldName.index)
        output.writeShort(_fieldDescriptor.index)
        output.writeShort(index)
    }

    def deserialize(input: DataInputStream) {
        startPc = input.readUnsignedShort()
        endPc = startPc + input.readUnsignedShort()
        _fieldName = _owner.constants().getUtf8ByIndex(
                input.readUnsignedShort())
        _fieldDescriptor = _owner.constants().getUtf8ByIndex(
                input.readUnsignedShort())
        index = input.readUnsignedShort()

        var parser = new DescriptorParser(_fieldDescriptor.value())
        _fieldType = parser.parseFieldDescriptor()
    }

    def debugString(indent: String): String = {
        return indent + "[" + startPc + ", " + endPc + ") index: " + index +
                " field: " + fieldName() + " " + _fieldDescriptor.value() + "\n"
    }
}

abstract class LocalVariableTableBaseAttribute(o: AttributeOwner, name: String)
        extends Attribute(o, name) {
    var table = new Vector[LocalVariableEntry]()

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2 + 10 * table.size())
        output.writeShort(table.size())
        for (entry <- table) {
            entry.serialize(output)
        }
    }

    def deserialize(
            name: ConstUtf8Info,
            attrLength: Int,
            input: DataInputStream) {
        val numEntries = input.readUnsignedShort()
        if (attrLength != (2 + 10 * numEntries)) {
            throw new Exception("Unexpected")
        }

        table = new Vector[LocalVariableEntry]()
        for (_ <- 1 to numEntries) {
            var entry = new LocalVariableEntry(_owner)
            entry.deserialize(input)
            table.add(entry)
        }
    }

    def debugString(indent: String): String = {
        var result = indent + "LocalVariableTable:\n"
        for (entry <- table) {
            result += entry.debugString(indent + "  ")
        }
        return result
    }
}

class LocalVariableTableAttribute(o: AttributeOwner)
        extends LocalVariableTableBaseAttribute(o, "LocalVariableTable") {
}

// TODO handle parameterize types
class LocalVariableTypeTableAttribute(o: AttributeOwner)
        extends LocalVariableTableBaseAttribute(o, "LocalVariableTypeTable") {
}
