import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


// see 154-156 for details
class BootstrapMethodEntry(o: AttributeOwner, mh: ConstMethodHandleInfo) {
    def this(o: AttributeOwner) = this(o, null)

    var _owner = o

    var _methodHandle: ConstMethodHandleInfo = mh
    var _arguments = new Vector[ConstInfo]()

    def methodHandle(): ConstMethodHandleInfo = _methodHandle
    def arguments(): Vector[ConstInfo] = _arguments

    def addArg(c: ConstInfo) {
        _checkArg(c)
        _arguments.add(c)
    }

    def _checkArg(c: ConstInfo) {
        c match {
            case _: ConstStringInfo => return
            case _: ConstClassInfo => return
            case _: ConstIntegerInfo => return
            case _: ConstLongInfo => return
            case _: ConstFloatInfo => return
            case _: ConstDoubleInfo => return
            case _: ConstMethodHandleInfo => return
            case _: ConstMethodTypeInfo => return
            case _ => throw new Exception(
                    "Invalid bootstrap method argument type: " + c.typeName())
        }
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_methodHandle.index)
        output.writeShort(_arguments.size())

        for (arg <- _arguments) {
            output.writeShort(arg.index)
        }
    }

    def deserialize(input: DataInputStream): Int = {
        if (!_arguments.isEmpty()) {
            throw new Exception(
                    "deserializing into non-emtpy boostrap method info")
        }

        _methodHandle = _owner.constants().getMethodHandleByIndex(
                input.readUnsignedShort())

        val argCount = input.readUnsignedShort()
        for (_ <- 1 to argCount) {
            val arg = _owner.constants().getByIndex(input.readUnsignedShort())
            _checkArg(arg)
            _arguments.add(arg)
        }

        return 4 + 2 * argCount
    }

    def debugString(indent: String): String = {
        var result = indent + _methodHandle.debugValue() + "\n"
        for (c <- _arguments) {
            result += indent + "  " + c.debugValue() + "\n"
        }

        return result
    }
}

// see 154-156 for details
class BootstrapMethodsAttribute(
        o: AttributeOwner) extends Attribute(o, "BootstrapMethods") {

    var _methods = new Vector[BootstrapMethodEntry]()

    def methods(): Vector[BootstrapMethodEntry] = _methods

    def add(mh: ConstMethodHandleInfo): BootstrapMethodEntry = {
        val m = new BootstrapMethodEntry(_owner, mh)
        _methods.add(m)
        return m
    }

    def serialize(output: DataOutputStream) {
        var buffer = new ByteArrayOutputStream()
        var bufferWriter = new DataOutputStream(buffer)

        for (m <- _methods) {
            m.serialize(bufferWriter)
        }

        val value = buffer.toByteArray()

        output.writeShort(_name.index)
        output.writeInt(2 + value.length)
        output.writeShort(_methods.size())
        output.write(value)
    }

    def deserialize(
            name: ConstUtf8Info,
            attrLength: Int,
            input: DataInputStream) {
        if (name.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + name.value())
        }

        val numMethods = input.readUnsignedShort()
        var totalRead = 2

        for (_ <- 1 to numMethods) {
            var m = new BootstrapMethodEntry(_owner)
            totalRead += m.deserialize(input)
            _methods.append(m)
        }
    }

    def debugString(indent: String): String = {
        var result = indent + name() + "\n"
        for (m <- _methods) {
            result += m.debugString(indent + "  ")
        }
        return result
    }
}
