import java.io.ByteArrayOutputStream
import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.HashMap
import java.util.Vector


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
    var tagTopoOrder = new Vector[Int]()

    // no dependencies
    tagTopoOrder.add(INTEGER)
    tagTopoOrder.add(LONG)
    tagTopoOrder.add(FLOAT)
    tagTopoOrder.add(DOUBLE)
    tagTopoOrder.add(UTF8)
    // depends on utf8
    tagTopoOrder.add(STRING)
    tagTopoOrder.add(CLASS)
    tagTopoOrder.add(METHOD_TYPE)
    tagTopoOrder.add(NAME_AND_TYPE)
    // depends on class / name and type
    tagTopoOrder.add(FIELD_REF)
    tagTopoOrder.add(METHOD_REF)
    tagTopoOrder.add(INTERFACE_METHOD_REF)
    // depends on ref infos
    tagTopoOrder.add(METHOD_HANDLE)
    // depends on (external attribute) bootstrap method index
    tagTopoOrder.add(INVOKE_DYNAMIC)

    var tagTopoOrderMap = new HashMap[Int, Int]()

    for (i <- 0 to (tagTopoOrder.size() - 1)) {
        tagTopoOrderMap.put(tagTopoOrder.elementAt(i), i)
    }

    def tagOrder(tag: Int): Int = {
        if (!tagTopoOrderMap.containsKey(tag)) {
            throw new Exception("Unknown tag type: " + tag)
        }
        return tagTopoOrderMap.get(tag)
    }
}

abstract class ConstInfo(c: ClassInfo) extends Comparable[ConstInfo] {
    var _owner = c

    var _used = false

    var index = 0

    def indexSize(): Int = 1

    def tag(): Int

    def typeName(): String

    def debugString(): String = {
        val indexValue = _debugIndexValue()
        if (indexValue != "") {
            return String.format(
                    "%7s = %-18s %-15s  // %s",
                    "#" + index,
                    typeName(),
                    indexValue,
                    debugValue())
        }

        return String.format(
                "%7s = %-18s %s",
                "#" + index,
                typeName(),
                debugValue())
    }

    def _debugIndexValue(): String = ""

    def debugValue(): String

    def markUsed() {
        _used = true
    }

    def serialize(output: DataOutputStream)

    def deserialize(parsedTag: Int, input: DataInputStream)

    def bindConstReferences()

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

// See section 4.4.7 page 85-86
//
// TODO: Fix encoding / decoding.  jvm uses a non-standard "modified-utf8"
// encoding.  Why use the standard when you can reinvent your own ...
class ConstUtf8Info(o: ClassInfo, v: String) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, "")

    var _value: String = v

    def tag(): Int = ConstInfo.UTF8

    def typeName(): String = "Utf8"

    def value(): String = _value

    def debugValue(): String = {
        return _value.replaceAll("\\p{C}", "?")
    }

    def serialize(output: DataOutputStream) {
        var buffer = new ByteArrayOutputStream()
        (new DataOutputStream(buffer)).writeUTF(_value)
        val utf8Value = buffer.toByteArray()

        if (utf8Value.length > Const.UINT16_MAX + 2) {
            throw new Exception("utf8 string too long")
        }

        output.writeByte(tag())
        output.write(utf8Value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }

        val length = input.readUnsignedShort()

        if (length == 0) {
            _value = ""
        } else {
            var utf8Bytes = new Array[Byte](length)
            input.readFully(utf8Bytes)

            _value = new String(utf8Bytes, "UTF-8")
        }
    }

    def bindConstReferences() {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstUtf8Info => {
                return _value.compareTo(other._value)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.4 page 82-83
class ConstIntegerInfo(o: ClassInfo, v: Int) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0)

    var _value: Int = v

    def tag(): Int = ConstInfo.INTEGER

    def typeName(): String = "Integer"

    def value(): Int = _value

    def debugValue(): String = {
        return "" + _value
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeInt(_value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _value = input.readInt()
    }

    def bindConstReferences() {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstIntegerInfo => {
                if (_value < other._value) {
                    return -1
                }
                if (_value > other._value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.5 page 83-85
class ConstLongInfo(o: ClassInfo, v: Long) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0)

    var _value: Long = v

    def tag(): Int = ConstInfo.LONG

    override def indexSize(): Int = 2

    def typeName(): String = "Long"

    def value(): Long = _value

    def debugValue(): String = {
        return "" + _value + "l"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeLong(_value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _value = input.readLong()
    }

    def bindConstReferences() {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstLongInfo => {
                if (_value < other._value) {
                    return -1
                }
                if (_value > other._value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.4 page 82-83
class ConstFloatInfo(o: ClassInfo, v: Float) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0)

    var _value: Float = v

    def tag(): Int = ConstInfo.FLOAT

    def typeName(): String = "Float"

    def value(): Float = _value

    def debugValue(): String = {
        return "" + _value + "f"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeFloat(_value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _value = input.readFloat()
    }

    def bindConstReferences() {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstFloatInfo => {
                if (_value < other._value) {
                    return -1
                }
                if (_value > other._value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// See section 4.4.5 page 83-85
class ConstDoubleInfo(o: ClassInfo, v: Double) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0)

    var _value: Double = v

    def tag(): Int = ConstInfo.DOUBLE

    override def indexSize(): Int = 2

    def typeName(): String = "Double"

    def value(): Double = _value

    def debugValue(): String = {
        return "" + _value + "d"
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeDouble(_value)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _value = input.readDouble()
    }

    def bindConstReferences() {
        // nothing to bind
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstDoubleInfo => {
                if (_value < other._value) {
                    return -1
                }
                if (_value > other._value) {
                    return 1
                }
                return 0
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.3 page 81-82
class ConstStringInfo(o: ClassInfo, v: ConstUtf8Info) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, null)

    var _utf8String: ConstUtf8Info = v

    // only used during deserialization
    var _tmpUtf8StringIndex = 0

    def tag(): Int = ConstInfo.STRING

    def typeName(): String = "String"

    def value(): String = _utf8String.value()

    override def _debugIndexValue(): String = {
        return "#" + _tmpUtf8StringIndex
    }

    def debugValue(): String = {
        return _utf8String.debugValue()
    }

    override def markUsed() {
        _used = true
        _utf8String.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_utf8String.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpUtf8StringIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _utf8String = _owner.constants().getUtf8ByIndex(_tmpUtf8StringIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstStringInfo => {
                return _utf8String.compareTo(other._utf8String)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.1 page 79-80
class ConstClassInfo(o: ClassInfo, n: ConstUtf8Info) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, null)

    var _className: ConstUtf8Info = n

    // only used during deserialization
    var _tmpClassNameIndex = 0

    def tag(): Int = ConstInfo.CLASS

    def typeName(): String = "Class"

    def className(): String = _className.value()

    override def _debugIndexValue(): String = {
        return "#" + _tmpClassNameIndex
    }

    def debugValue(): String = {
        return _className.debugValue()
    }

    override def markUsed() {
        _used = true
        _className.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_className.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpClassNameIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _className = _owner.constants().getUtf8ByIndex(_tmpClassNameIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstClassInfo => {
                return _className.compareTo(other._className)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.9 page 89
class ConstMethodTypeInfo(o: ClassInfo, d: MethodType) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, null)

    var _methodDescriptor: MethodType = d

    var _descriptorString: ConstUtf8Info = null
    if (d != null) {
        _descriptorString = _owner.constants().getUtf8(d.descriptorString())
    }

    // only used during deserialization
    var _tmpDescriptorIndex = 0

    def tag(): Int = ConstInfo.METHOD_TYPE

    def typeName(): String = "MethodType"

    def methodDescriptor(): MethodType = _methodDescriptor
    def descriptor(): String = _descriptorString.value()

    override def _debugIndexValue(): String = {
        return "#" + _tmpDescriptorIndex
    }

    def debugValue(): String = {
        return _descriptorString.debugValue()
    }

    override def markUsed() {
        _used = true
        _descriptorString.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_descriptorString.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpDescriptorIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _descriptorString = _owner.constants().getUtf8ByIndex(
                _tmpDescriptorIndex)
        var parser = new DescriptorParser(_descriptorString.value)
        _methodDescriptor = parser.parseMethodDescriptor()
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstMethodTypeInfo => {
                return _descriptorString.compareTo(other._descriptorString)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.6 page 85
class ConstNameAndTypeInfo(
        o: ClassInfo,
        n: ConstUtf8Info,
        d: ConstUtf8Info) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, null, null)

    var _name: ConstUtf8Info = n
    // NOTE: we can't parse the descriptor yet since we don't know
    // if it's either a method or field descriptor.
    var _descriptorString: ConstUtf8Info = d

    // only used during deserialization
    var _tmpNameIndex = 0
    var _tmpDescriptorIndex = 0

    def tag(): Int = ConstInfo.NAME_AND_TYPE

    def typeName(): String = "NameAndType"

    def name(): String = _name.value()
    def descriptorString(): String = _descriptorString.value()

    override def _debugIndexValue(): String = {
        return "#" + _tmpNameIndex + ":#" + _tmpDescriptorIndex
    }

    def debugValue(): String = {
        return _name.debugValue() + ":" + _descriptorString.debugValue()
    }

    override def markUsed() {
        _used = true
        _name.markUsed()
        _descriptorString.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_name.index)
        output.writeShort(_descriptorString.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpNameIndex = input.readUnsignedShort()
        _tmpDescriptorIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _name = _owner.constants().getUtf8ByIndex(_tmpNameIndex)
        _descriptorString = _owner.constants().getUtf8ByIndex(
                _tmpDescriptorIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstNameAndTypeInfo => {
                val c = _name.compareTo(other._name)
                if (c != 0) {
                    return c
                }
                return _descriptorString.compareTo(other._descriptorString)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.2 page 80
abstract class ConstRefInfo(
        o: ClassInfo,
        c: ConstClassInfo,
        n: ConstNameAndTypeInfo,
        d: DescriptorType) extends ConstInfo(o) {

    var _classInfo: ConstClassInfo = c
    var _nameAndType: ConstNameAndTypeInfo = n
    var _descriptor: DescriptorType = d

    // only used during deserialization
    var _tmpClassIndex = 0
    var _tmpNameAndTypeIndex = 0

    def _isFieldRef(): Boolean = false

    def className(): String = _classInfo.className()

    def referenceName(): String = _nameAndType.name()

    def descriptorString(): String = _nameAndType.descriptorString()

    def fieldDescriptor(): FieldType = {
        _descriptor match {
            case f: FieldType => return f
            case _ => throw new Exception("unexpected descriptor type")
        }
    }

    def methodDescriptor(): MethodType = {
        _descriptor match {
            case m: MethodType => return m
            case _ => throw new Exception("unexpected descriptor type")
        }
    }

    override def _debugIndexValue(): String = {
        return "#" + _tmpClassIndex + ".#" + _tmpNameAndTypeIndex
    }

    def debugValue(): String = {
        return _classInfo.debugValue() + "." + _nameAndType.debugValue()
    }

    override def markUsed() {
        _used = true
        _classInfo.markUsed()
        _nameAndType.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_classInfo.index)
        output.writeShort(_nameAndType.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _tmpClassIndex = input.readUnsignedShort()
        _tmpNameAndTypeIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _classInfo = _owner.constants().getClassByIndex(_tmpClassIndex)
        _nameAndType = _owner.constants().getNameAndTypeByIndex(
                _tmpNameAndTypeIndex)

        var parser = new DescriptorParser(_nameAndType.descriptorString())
        if (_isFieldRef()) {
            _descriptor = parser.parseFieldDescriptor()
        } else {
            _descriptor = parser.parseMethodDescriptor()
        }
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstRefInfo => {
                val c = _classInfo.compareTo(other._classInfo)
                if (c != 0) {
                    return c
                }
                return _nameAndType.compareTo(other._nameAndType)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

// see section 4.4.2 page 80
class ConstFieldRefInfo(
        o: ClassInfo,
        c: ConstClassInfo,
        n: ConstNameAndTypeInfo,
        d: FieldType) extends ConstRefInfo(o, c, n, d) {
    def this(o: ClassInfo) = this(o, null, null, null)

    def tag(): Int = ConstInfo.FIELD_REF

    def typeName(): String = "FieldRef"

    override def _isFieldRef(): Boolean = true
}

abstract class ConstBaseMethodRefInfo(
        o: ClassInfo,
        c: ConstClassInfo,
        n: ConstNameAndTypeInfo,
        d: MethodType) extends ConstRefInfo(o, c, n, d) {
}

// see section 4.4.2 page 80
class ConstMethodRefInfo(
        o: ClassInfo,
        c: ConstClassInfo,
        n: ConstNameAndTypeInfo,
        d: MethodType) extends ConstBaseMethodRefInfo(o, c, n, d) {
    def this(o: ClassInfo) = this(o, null, null, null)

    def tag(): Int = ConstInfo.METHOD_REF

    def typeName(): String = "MethodRef"
}

// see section 4.4.2 page 80
class ConstInterfaceMethodRefInfo(
        o: ClassInfo,
        c: ConstClassInfo,
        n: ConstNameAndTypeInfo,
        d: MethodType) extends ConstBaseMethodRefInfo(o, c, n, d) {
    def this(o: ClassInfo) = this(o, null, null, null)

    def tag(): Int = ConstInfo.INTERFACE_METHOD_REF

    def typeName(): String = "InterfaceMethodRef"
}

// see section 4.4.8 page 87-89
//
// NOTE: use ConstMethodHandleInfo.New<Kind>MethodHandle(ref) to initialize
// the constructor!!!
class ConstMethodHandleInfo(
        o: ClassInfo,
        kind: Byte,
        ref: ConstRefInfo) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0, null)

    var _referenceKind: Byte = kind
    var _reference: ConstRefInfo = ref

    // only used during deserialization
    var _tmpReferenceIndex = 0

    def tag(): Int = ConstInfo.METHOD_HANDLE

    def typeName(): String = "MethodHandle"

    def referenceKind(): Byte = _referenceKind

    def className(): String = _reference.className()

    def referenceName(): String = _reference.referenceName()

    def descriptorString(): String = _reference.descriptorString()

    def fieldDescriptor(): FieldType = _reference.fieldDescriptor()

    def methodDescriptor(): MethodType = _reference.methodDescriptor()

    override def _debugIndexValue(): String = {
        return "" + _referenceKind + " #" + _tmpReferenceIndex
    }

    def debugValue(): String = {
        return "" + _referenceKind + " " + _reference.debugValue()
    }

    def _checkIsFieldRef() {
        _reference match {
            case _: ConstFieldRefInfo => return
            case _ => throw new Exception("invalid ref type")
        }
    }

    def _checkIsBaseMethodRef() {
        _reference match {
            case _: ConstBaseMethodRefInfo => return
            case _ => throw new Exception("invalid ref type")
        }
    }

    def _checkIsMethodRef() {
        _reference match {
            case _: ConstMethodRefInfo => return
            case _ => throw new Exception("invalid ref type")
        }
    }

    def _checkIsInterfaceMethodRef() {
        _reference match {
            case _: ConstInterfaceMethodRefInfo => return
            case _ => throw new Exception("invalid ref type")
        }
    }

    override def markUsed() {
        _used = true
        _reference.markUsed()
    }

    def serialize(output: DataOutputStream) {
        _referenceKind match {
            case 1 => _checkIsFieldRef()
            case 2 => _checkIsFieldRef()
            case 3 => _checkIsFieldRef()
            case 4 => _checkIsFieldRef()
            case 5 => _checkIsMethodRef()
            case 6 => _checkIsBaseMethodRef()
            case 7 => _checkIsBaseMethodRef()
            case 8 => _checkIsMethodRef()
            case 9 => _checkIsInterfaceMethodRef()
            case _ => throw new Exception("Unknown reference kind")
        }

        output.writeByte(tag())
        output.writeByte(_referenceKind)
        output.writeShort(_reference.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _referenceKind = input.readByte()
        _tmpReferenceIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _reference = _owner.constants().getRefByIndex(_tmpReferenceIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstMethodHandleInfo => {
                if (_referenceKind < other._referenceKind) {
                    return -1
                }
                if (_referenceKind > other._referenceKind) {
                    return 1
                }
                return _reference.compareTo(other._reference)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}

object ConstMethodHandleInfo {
    //
    // kinds with field ref
    //
    def NewGetFieldMethodHandle(
            owner: ClassInfo,
            ref: ConstFieldRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 1, ref)
    }

    def NewGetStaticMethodHandle(
            owner: ClassInfo,
            ref: ConstFieldRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 2, ref)
    }

    def NewPutFieldMethodHandle(
            owner: ClassInfo,
            ref: ConstFieldRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 3, ref)
    }

    def NewPutStaticMethodHandle(
            owner: ClassInfo,
            ref: ConstFieldRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 4, ref)
    }

    //
    // kinds with method ref
    //
    def NewInvokeVirtualMethodHandle(
            owner: ClassInfo,
            ref: ConstMethodRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 5, ref)
    }

    def NewNewInvokeVirtualMethodHandle(
            owner: ClassInfo,
            ref: ConstMethodRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 8, ref)
    }

    //
    // kinds with either method or interface method ref
    //
    def NewInvokeStaticMethodHandle(
            owner: ClassInfo,
            ref: ConstBaseMethodRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 6, ref)
    }

    def NewInvokeSpecialMethodHandle(
            owner: ClassInfo,
            ref: ConstBaseMethodRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 7, ref)
    }

    //
    // kinds with interface method ref
    //
    def NewInvokeInterfaceMethodHandle(
            owner: ClassInfo,
            ref: ConstInterfaceMethodRefInfo): ConstMethodHandleInfo = {
        return new ConstMethodHandleInfo(owner, 9, ref)
    }
}


// TODO: bind bootstrap method attr index to real entry
class ConstInvokeDynamicInfo(
        o: ClassInfo,
        i: Int,
        n: ConstNameAndTypeInfo,
        m: MethodType) extends ConstInfo(o) {
    def this(o: ClassInfo) = this(o, 0, null, null)

    var _bootstrapMethodAttrIndex = i
    var _nameAndType: ConstNameAndTypeInfo = n
    var _methodType: MethodType = m

    // only used during deserialization
    var _tmpNameAndTypeIndex = 0

    def tag(): Int = ConstInfo.INVOKE_DYNAMIC

    def typeName(): String = "InvokeDynamic"

    def bootstrapMethodAttrIndex(): Int = _bootstrapMethodAttrIndex
    def refName(): String = _nameAndType.name()
    def descriptorString(): String = _nameAndType.descriptorString()
    def methodType(): MethodType = _methodType

    override def _debugIndexValue(): String = {
        return "" + _bootstrapMethodAttrIndex + " #" + _tmpNameAndTypeIndex
    }

    def debugValue(): String = {
        return "" + _bootstrapMethodAttrIndex + " " + _nameAndType.debugValue()
    }

    override def markUsed() {
        _used = true
        _nameAndType.markUsed()
    }

    def serialize(output: DataOutputStream) {
        output.writeByte(tag())
        output.writeShort(_bootstrapMethodAttrIndex)
        output.writeShort(_nameAndType.index)
    }

    def deserialize(parsedTag: Int, input: DataInputStream) {
        if (parsedTag != tag()) {
            throw new Exception("unexpected tag")
        }
        _bootstrapMethodAttrIndex = input.readUnsignedShort()
        _tmpNameAndTypeIndex = input.readUnsignedShort()
    }

    def bindConstReferences() {
        _nameAndType = _owner.constants().getNameAndTypeByIndex(
                _tmpNameAndTypeIndex)
    }

    def _compareTo(o: ConstInfo): Int = {
        o match {
            case other: ConstInvokeDynamicInfo => {
                if (_bootstrapMethodAttrIndex <
                        other._bootstrapMethodAttrIndex) {
                    return -1
                }
                if (_bootstrapMethodAttrIndex >
                        other._bootstrapMethodAttrIndex) {
                    return 1
                }
                return _nameAndType.compareTo(other._nameAndType)
            }
            case _ => throw new Exception("unexpected other type")
        }
    }
}
