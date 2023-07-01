import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


// see page 114- 116 for details
class InnerClassEntry(
        o: AttributeOwner,
        ic: String,
        oc: String,
        in: String) {
    def this(o: AttributeOwner) = this(o, null, null, null)

    var _owner = o

    var _innerClass: ConstClassInfo = null
    if (ic != null) {
        _innerClass = _owner.constants().getClass(ic)
    }

    var _outerClass: ConstClassInfo = null
    if (oc != null) {
        _outerClass = _owner.constants().getClass(oc)
    }

    var _innerName: ConstUtf8Info = null
    if (in != null) {
        _innerName = _owner.constants().getUtf8(in)
    }

    var _access = new InnerClassAccessFlags(this)

    def serialize(output: DataOutputStream) {
        output.writeShort(_innerClass.index)

        if (_outerClass == null) {
            output.writeShort(0)
        } else {
            output.writeShort(_outerClass.index)
        }

        if (_innerName == null) {
            output.writeShort(0)
        } else {
            output.writeShort(_innerName.index)
        }

        _access.serialize(output)
    }

    def innerClass(): String = _innerClass.className()
    def outerClass(): String = {
        if (_outerClass == null) {
            return null
        }
        return _outerClass.className()
    }

    // original name of innerClass
    def originalInnerName(): String = {
        if (_innerName == null) {
            return null
        }
        return _innerName.value()
    }



    def deserialize(input: DataInputStream) {
        _innerClass = _owner.constants().getClassByIndex(
                input.readUnsignedShort())

        val outerClassIndex = input.readUnsignedShort()
        if (outerClassIndex == 0) {
            _outerClass = null
        } else {
            _outerClass = _owner.constants().getClassByIndex(outerClassIndex)
        }

        val innerNameIndex = input.readUnsignedShort()
        if (innerNameIndex == 0) {
            _innerName = null
        } else {
            _innerName = _owner.constants().getUtf8ByIndex(innerNameIndex)
        }

        _access.deserialize(input)
    }

    def debugString(indent: String): String = {
        var oc = outerClass()
        if (oc == null) {
            oc = "???"
        }

        var in = originalInnerName()
        if (in == null) {
            in = "???"
        }

        return indent + innerClass() + " of " + oc + " (name: " + in +
                " flags: " + _access.debugString() + ")\n"

    }
}

class InnerClassesAttribute(
        o: AttributeOwner) extends Attribute(o, "InnerClasses") {
    var _innerClasses = new Vector[InnerClassEntry]()

    def _attrSize(numClasses: Int): Int = {
        return 2 + 8 * numClasses
    }

    def classes(): Vector[InnerClassEntry] = _innerClasses

    def add(c: InnerClassEntry) {
        _innerClasses.add(c)
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(_attrSize(_innerClasses.size()))
        output.writeShort(_innerClasses.size())

        for (inner <- _innerClasses) {
            inner.serialize(output)
        }
    }

    def deserialize(name: ConstUtf8Info,
                    attrLength: Int,
                    input: DataInputStream) {
        if (name.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + name.value())
        }
        val numClasses = input.readUnsignedShort()
        val expectedLength = _attrSize(numClasses)
        if (expectedLength != attrLength) {
            throw new Exception("Unexpected length mismatch")
        }

        for (_ <- 1 to numClasses) {
            var inner = new InnerClassEntry(_owner)
            inner.deserialize(input)
            add(inner)
        }
    }

    def debugString(indent: String): String = {
        if (_innerClasses.isEmpty()) {
            return indent + name() + ": (no inner classes)\n"
        }
        var result = indent + name() + ":\n"

        for (c <- _innerClasses) {
            result += c.debugString(indent + "  ")
        }

        return result
    }
}

