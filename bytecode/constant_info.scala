import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.HashMap


object ConstInfo {
    // see table 4.4-A (page 79)
    val CLASS = 7
    val FIELD_REF = 9
    val METHOD_REF = 10
    val INTERFACE_METHOD_REF = 11
    val STRING = 8
    val INTEGER = 3
    val FLOAT = 4
    val LONG = 5
    val DOUBLE = 6
    val NAME_AND_TYPE = 12
    val UTF8 = 1
    val METHOD_HANDLE = 15
    val METHOD_TYPE = 16
    val INVOKE_DYNAMIC = 18

    // tag -> order
    var tagTopoOrder = new HashMap[Int, Int]()

    tagTopoOrder.put(UTF8, 0)
    tagTopoOrder.put(INTEGER, 1)
    tagTopoOrder.put(FLOAT, 2)
    tagTopoOrder.put(LONG, 3)
    tagTopoOrder.put(DOUBLE, 4)

/* TODO
    tagTopoOrder.put(CLASS, )
    tagTopoOrder.put(FIELD_REF, )
    tagTopoOrder.put(METHOD_REF, )
    tagTopoOrder.put(INTERFACE_METHOD_REF, )
    tagTopoOrder.put(STRING, )
    tagTopoOrder.put(NAME_AND_TYPE, )
    tagTopoOrder.put(METHOD_HANDLE, )
    tagTopoOrder.put(METHOD_TYPE, )
    tagTopoOrder.put(INVOKE_DYNAMIC, )
*/

    def tagOrder(tag: Int): Int = {
        if (tagTopoOrder.containsKey(tag)) {
            return tagTopoOrder.get(tag)
        }
        return 10000
    }
}

trait ConstInfo extends Comparable[ConstInfo] {
    var index = 0

    def indexSize(): Int = 1

    def tag(): Int

    def typeName(): String

    def serialize(output: DataOutputStream)

    def deserialize(parsedTag: Int, input: DataInputStream)

    def bindConstReferences(pool: ConstantPool)

    def compareTo(other: ConstInfo): Int = {
        if (tag() < other.tag()) {
            return -1
        }
        if (tag() > other.tag()) {
            return 1
        }
        return _compareTo(other)
    }

    def _compareTo(other: ConstInfo): Int
}

// TODO: Fix encoding / decoding.  jvm uses a non-standard "modified-utf8"
// encoding (see section 4.4.7 page 85-86).
class ConstUtf8Info extends ConstInfo {
    var value = ""

    def tag(): Int = ConstInfo.UTF8

    def typeName(): String = "Utf8"

    def serialize(output: DataOutputStream) {
        var buffer = new ByteArrayOutputStream()
        (new DataOutputStream(buffer)).writeUTF(value)
        val utf8Value = buffer.toByteArray()

        if (utf8Value.length > 65535) {
            throw new Exception("utf8 string too long")
        }

        output.writeByte(tag())
        output.writeShort(utf8Value.length)
        output.write(utf8Value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }

        val length = input.readUnsignedShort()

        var utf8Bytes = new Array[Byte](length)
        input.readFully(utf8Bytes)

        value = new String(utf8Bytes, "UTF-8")
    }

    def bindConstReferences(pool: ConstantPool) {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstUtf8Info => {
                return value.compareTo(other.value)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

class ConstClassInfo extends ConstInfo {
    def tag(): Int = ConstInfo.CLASS

    def typeName(): String = "Class"

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def bindConstReferences(pool: ConstantPool) {
        throw new Exception("TODO")
    }

    def _compareTo(other: ConstInfo): Int = {
        // TODO
        return 0
    }
}

class ConstStringInfo extends ConstInfo {
    def tag(): Int = ConstInfo.STRING

    def typeName(): String = "String"

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def bindConstReferences(pool: ConstantPool) {
        throw new Exception("TODO")
    }

    def _compareTo(other: ConstInfo): Int = {
        // TODO
        return 0
    }
}

class ConstLongInfo extends ConstInfo {
    def tag(): Int = ConstInfo.LONG

    def typeName(): String = "Long"

    override def indexSize(): Int = 2

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def bindConstReferences(pool: ConstantPool) {
        throw new Exception("TODO")
    }

    def _compareTo(other: ConstInfo): Int = {
        // TODO
        return 0
    }
}

class ConstDoubleInfo extends ConstInfo {
    def tag(): Int = ConstInfo.DOUBLE

    def typeName(): String = "Double"

    override def indexSize(): Int = 2

    def serialize(output: DataOutputStream) {
        throw new Exception("TODO")
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        throw new Exception("TODO")
    }

    def bindConstReferences(pool: ConstantPool) {
        throw new Exception("TODO")
    }

    def _compareTo(other: ConstInfo): Int = {
        // TODO
        return 0
    }
}
