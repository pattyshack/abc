import java.io.DataInputStream
import java.io.DataOutputStream

import scala.reflect.ClassTag


class ConstantValueAttribute(
        o: AttributeOwner) extends Attribute(o, "ConstantValue") {
    var _constant: ConstInfo = null

    def _checkOwnerType[T <: FieldType: ClassTag]() {
        if (_owner.fieldType() == null) {
            throw new Exception("owner does not have field type")
        }

        val cls = implicitly[ClassTag[T]].runtimeClass
        _owner.fieldType() match {
            case t: T if cls.isInstance(t) => return
            case _ => throw new Exception("field type mismatch")
        }
    }

    def _checkOwnerStringType() {
        if (_owner.fieldType() == null) {
            throw new Exception("owner does not have field type")
        }

        _owner.fieldType() match {
            case t: ObjectType => {
                if (!t.isJavaString()) {
                    throw new Exception("field type mismatch")
                }
            }
            case _ => throw new Exception("field type mismatch")
        }
    }

    def intValue(): Int = {
        _checkOwnerType[IntType]()
        _constant match {
            case v: ConstIntegerInfo => return v.value()
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setIntValue(v: Int) {
        _checkOwnerType[IntType]()
        _constant = _owner.constants().getInteger(v)
    }

    def shortValue(): Short = {
        _checkOwnerType[ShortType]()
        _constant match {
            case v: ConstIntegerInfo => return v.value().toShort
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setShortValue(v: Short) {
        _checkOwnerType[ShortType]()
        _constant = _owner.constants().getInteger(v)
    }

    def charValue(): Char = {
        _checkOwnerType[CharType]()
        _constant match {
            case v: ConstIntegerInfo => return v.value().toChar
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setCharValue(v: Char) {
        _checkOwnerType[CharType]()
        _constant = _owner.constants().getInteger(v)
    }

    def byteValue(): Byte = {
        _checkOwnerType[ByteType]()
        _constant match {
            case v: ConstIntegerInfo => return v.value().toByte
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setByteValue(v: Byte) {
        _checkOwnerType[ByteType]()
        _constant = _owner.constants().getInteger(v)
    }

    def boolValue(): Boolean = {
        _checkOwnerType[BoolType]()
        _constant match {
            case v: ConstIntegerInfo => return v.value() != 0
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setBoolValue(v: Boolean) {
        _checkOwnerType[BoolType]()
        _constant = _owner.constants().getInteger(if (v) 1 else 0)
    }

    def longValue(): Long = {
        _checkOwnerType[LongType]()
        _constant match {
            case v: ConstLongInfo => return v.value()
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setLongValue(v: Long) {
        _checkOwnerType[LongType]()
        _constant = _owner.constants().getLong(v)
    }

    def floatValue(): Float = {
        _checkOwnerType[FloatType]()
        _constant match {
            case v: ConstFloatInfo => return v.value()
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setFloatValue(v: Float) {
        _checkOwnerType[FloatType]()
        _constant = _owner.constants().getFloat(v)
    }

    def doubleValue(): Double = {
        _checkOwnerType[DoubleType]()
        _constant match {
            case v: ConstDoubleInfo => return v.value()
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setDoubleValue(v: Double) {
        _checkOwnerType[DoubleType]()
        _constant = _owner.constants().getDouble(v)
    }

    def stringValue(): String = {
        _checkOwnerStringType()
        _constant match {
            case v: ConstStringInfo => return v.value()
            case _ => throw new Exception("const type mismatch")
        }
    }
    def setStringValue(v: String) {
        _checkOwnerStringType()
        _constant = _owner.constants().getString(v)
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_name.index)
        output.writeInt(2)
        output.writeShort(_constant.index)
    }

    def deserialize(
            n: ConstUtf8Info,
            attrLength: Int,
            input: DataInputStream) {
        if (n.compareTo(_name) != 0) {
            throw new Exception("Unexpected attribute name: " + n.value())
        }
        if (attrLength != 2) {
            throw new Exception("Unexpected attribute length")
        }
        _constant = _owner.constants().getByIndex(input.readUnsignedShort())

        // sanity check
        _owner.fieldType() match {
            case _: ByteType => byteValue()
            case _: CharType => charValue()
            case _: DoubleType => doubleValue()
            case _: FloatType => floatValue()
            case _: IntType => intValue()
            case _: LongType => longValue()
            case _: ShortType => shortValue()
            case _: BoolType => boolValue()
            case t: ObjectType => {
                if (!t.isJavaString()) {
                    throw new Exception(
                            "cannot use constant value with non-String object")
                }
                stringValue()
            }
        }
    }

    def debugString(indent: String): String = {
        return indent + name() + ": " + _constant.debugValue()
    }
}
