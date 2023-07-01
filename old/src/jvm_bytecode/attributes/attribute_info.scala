import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.TreeMap
import java.util.Vector

import scala.collection.JavaConversions._


trait AttributeOwner {
    def constants(): ConstantPool

    // only applicable if the owner is FieldInfo
    def fieldType(): FieldType = null

    // only applicable if the owner is MethodInfo
    def methodType(): MethodType = null
}

abstract class Attribute(o: AttributeOwner, attributeName: String) {
    var _owner: AttributeOwner = o
    var _name: ConstUtf8Info = null
    if (attributeName != null) {
        _name = _owner.constants().getUtf8(attributeName)
    }

    def name(): String = _name.value()

    def serialize(output: DataOutputStream)

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream)

    def debugString(indent: String): String
}

abstract class RawBytesAttribute(
        o: AttributeOwner,
        attributeName: String,
        b: Array[Byte]) extends Attribute(o, attributeName) {
    var _bytes: Array[Byte] = b

    def bytes(): Array[Byte] = _bytes

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(_bytes.length)
        output.write(_bytes)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        _name = n
        _bytes = new Array[Byte](attrLength)
        input.readFully(_bytes)
    }
}

class NoValueAttribute(
        o: AttributeOwner,
        attributeName: String) extends Attribute(o, attributeName) {

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(0)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 0) {
            throw new Exception("Unexpected attribute length")
        }
    }

    def debugString(indent: String): String = indent + name()
}

class StringValueAttribute(
        o: AttributeOwner,
        attributeName: String,
        v: String) extends Attribute(o, attributeName) {

    var _value: ConstUtf8Info = null
    if (v != null) {
        _value = _owner.constants().getUtf8(v)
    }

    def value(): String = _value.value()

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2)
        output.writeShort(_value.index)
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 2) {
            throw new Exception("Unexpected attribute length")
        }
        _value = _owner.constants().getUtf8ByIndex(
                input.readUnsignedShort())
    }

    def debugString(indent: String): String = {
        return indent + name() + ": " + value()
    }
}

class UnsupportedAttribute(
        o: AttributeOwner,
        attributeName: String,
        b: Array[Byte]) extends RawBytesAttribute(o, attributeName, b) {
    def this(o: AttributeOwner) = this(o, null, null)

    def debugString(indent: String): String = indent + name() + " (Unsupported)"
}

object SourceDebugExtensionAttribute {
    def modifiedUtf8(s: String): Array[Byte] = {
        if (s == null) {
            return null
        }
        var buffer = new ByteArrayOutputStream()
        (new DataOutputStream(buffer)).writeUTF(s)
        return buffer.toByteArray()
    }
}

class SourceDebugExtensionAttribute(
        o: AttributeOwner,
        s: String) extends RawBytesAttribute(
                o,
                "SourceDebugExtension",
                SourceDebugExtensionAttribute.modifiedUtf8(s)) {
    def this(o: AttributeOwner) = this(o, null)

    def debugString(indent: String): String = indent + name()
}

class SourceFileAttribute(
        o: AttributeOwner,
        n: String) extends StringValueAttribute(o, "SourceFile", n) {
    def this(o: AttributeOwner) = this(o, null)
}

// see page 118-123 for signature syntax
class SignatureAttribute(
        o: AttributeOwner,
        n: String) extends StringValueAttribute(o, "Signature", n) {
    def this(o: AttributeOwner) = this(o, null)
}

class DeprecatedAttribute(
        o: AttributeOwner) extends NoValueAttribute(o, "Deprecated") {
}

class SyntheticAttribute(
        o: AttributeOwner) extends NoValueAttribute(o, "Synthetic") {
}

class EnclosingMethodAttribute(
        o: AttributeOwner,
        className: String,
        methodName: String,
        methodType: MethodType) extends Attribute(
                o,
                "EnclosingMethod") {
    def this(o: AttributeOwner) = this(o, null, null, null)

    var _enclosingClass: ConstClassInfo = null
    if (className != null) {
        _enclosingClass = _owner.constants().getClass(className)
    }

    var _methodNameAndType: ConstNameAndTypeInfo = null
    var _enclosingMethodType: MethodType = null
    if (methodName != null) {  // assume method type is also not null
        _methodNameAndType = _owner.constants().getNameAndType(
                methodName,
                methodType.descriptorString())
        _enclosingMethodType = methodType
    }

    def enclosingClassName(): String = _enclosingClass.className()

    def enclosingMethodName(): String = {
        if (_methodNameAndType == null) {
            return null
        }
        return _methodNameAndType.name()
    }

    def enclosingMethodType(): MethodType = return _enclosingMethodType

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(4)
        output.writeShort(_enclosingClass.index)
        if (_methodNameAndType == null) {
            output.writeShort(0)
        } else {
            output.writeShort(_methodNameAndType.index)
        }
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 4) {
            throw new Exception("Unexpected attribute length")
        }
        _enclosingClass = _owner.constants().getClassByIndex(
                input.readUnsignedShort())

        val index = input.readUnsignedShort()
        if (index == 0) {
            _methodNameAndType = null
            _enclosingMethodType = null
        } else {
            _methodNameAndType = _owner.constants().getNameAndTypeByIndex(index)
            var parser = new DescriptorParser(
                    _methodNameAndType.descriptorString())
            _enclosingMethodType = parser.parseMethodDescriptor()
        }
    }

    def debugString(indent: String): String = {
        var v = "???"
        if (_methodNameAndType != null) {
            v = _methodNameAndType.debugValue()
        }
        return indent + name() + ": " + enclosingClassName() + "." + v
    }
}

class ExceptionsAttribute(
        o: AttributeOwner) extends Attribute(o, "Exceptions") {
    var _exceptions = new Vector[ConstClassInfo]()

    def exceptions(): Vector[String] = {
        var result = new Vector[String]()
        for (c <- _exceptions) {
            result.add(c.className())
        }
        return result
    }

    def add(exceptionName: String) {
        _exceptions.add(_owner.constants().getClass(exceptionName))
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2 + 2 * _exceptions.size())
        output.writeShort(_exceptions.size())
        for (c <- _exceptions) {
            output.writeShort(c.index)
        }
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        val numEntries = input.readUnsignedShort()
        if (attrLength != (2 + 2 * numEntries)) {
            throw new Exception("Unexpected attribute length")
        }

        if (!_exceptions.isEmpty()) {
            throw new Exception(
                    "deserializing into non-empty Exceptions attribute")
        }

        for (_ <- 1 to numEntries) {
            _exceptions.add(_owner.constants().getClassByIndex(
                    input.readUnsignedShort()))
        }
    }

    def debugString(indent: String): String = {
        var result: String = null
        for (c <- _exceptions) {
            if (result == null) {
                result = c.className()
            } else {
                result += ", " + c.className()
            }
        }

        return indent + name() + ": " + result
    }
}

class LineNumberTableAttribute(o: AttributeOwner)
        extends Attribute(o, "LineNumberTable") {
    // pc -> line #
    var table = new TreeMap[Int, Int]()

    def mergeFrom(other: LineNumberTableAttribute) {
        for (entry <- table.entrySet()) {
            table.put(entry.getKey(), entry.getValue())
        }
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2 + 4 * table.size())
        output.writeShort(table.size())
        for (entry <- table.entrySet()) {
            output.writeShort(entry.getKey())
            output.writeShort(entry.getValue())
        }
    }

    def deserialize(n: ConstUtf8Info, attrLength: Int, input: DataInputStream) {
        val numEntries = input.readUnsignedShort()
        for (_ <- 1 to numEntries) {
            val pc = input.readUnsignedShort()
            val line = input.readUnsignedShort()
            table.put(pc, line)
        }
    }

    def debugString(indent: String): String = {
        var result = indent + "LineNumberTable:\n"
        for (entry <- table.entrySet()) {
            result += indent + "  " + entry.getKey() +
                    ": line " + entry.getValue() + "\n"
        }
        return result
    }
}

