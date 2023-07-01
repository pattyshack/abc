import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Collection
import java.util.HashMap
import java.util.TreeMap
import java.util.Vector

import scala.collection.JavaConversions._
import scala.reflect.ClassTag


class ConstantPool(owner: ClassInfo) {
    var _owner = owner

    var _finalized = false

    // using TreeMap instead of TreeSet to simplify deduplication / lookup.
    var _constInfos = new TreeMap[ConstInfo, ConstInfo]()

    // only used for deserialization.  invalidate when _constInfos is modified.
    var _tmpConstInfosByIndex: TreeMap[Int, ConstInfo] = null

    def _get[T <: ConstInfo: ClassTag](info: T): T = {
        if (_constInfos.containsKey(info)) {
            val cls = implicitly[ClassTag[T]].runtimeClass
            _constInfos.get(info) match {
                case t: T if cls.isInstance(t) => return t
                case _ => throw new Exception("unexpected const info type")
            }
        }

        if (_finalized) {
            throw new Exception("Cannot add more constants to finalized pool")
        }

        // invalidate only when we actually mutated the pool.
        _tmpConstInfosByIndex = null

        _constInfos.put(info, info)
        return info
    }

    def getUtf8(value: String): ConstUtf8Info = {
        return _get[ConstUtf8Info](new ConstUtf8Info(_owner, value))
    }

    def getInteger(value: Int): ConstIntegerInfo = {
        return _get[ConstIntegerInfo](new ConstIntegerInfo(_owner, value))
    }

    def getLong(value: Long): ConstLongInfo = {
        return _get[ConstLongInfo](new ConstLongInfo(_owner, value))
    }

    def getFloat(value: Float): ConstFloatInfo = {
        return _get[ConstFloatInfo](new ConstFloatInfo(_owner, value))
    }

    def getDouble(value: Double): ConstDoubleInfo = {
        return _get[ConstDoubleInfo](new ConstDoubleInfo(_owner, value))
    }

    def getString(value: String): ConstStringInfo = {
        return _get[ConstStringInfo](
                new ConstStringInfo(_owner, getUtf8(value)))
    }

    def getClass(name: String): ConstClassInfo = {
        return _get[ConstClassInfo](new ConstClassInfo(_owner, getUtf8(name)))
    }

    def getMethodType(methodType: MethodType): ConstMethodTypeInfo = {
        return _get[ConstMethodTypeInfo](
                new ConstMethodTypeInfo(
                        _owner,
                        methodType))
    }

    def getNameAndType(
            name: String,
            descriptorString: String): ConstNameAndTypeInfo = {
        return _get[ConstNameAndTypeInfo](
                new ConstNameAndTypeInfo(
                        _owner,
                        getUtf8(name),
                        getUtf8(descriptorString)))
    }

    def getFieldRef(
            className: String,
            refName: String,
            descriptor: FieldType): ConstFieldRefInfo = {
        return _get[ConstFieldRefInfo](
                new ConstFieldRefInfo(
                        _owner,
                        getClass(className),
                        getNameAndType(refName, descriptor.descriptorString()),
                        descriptor))
    }

    def getMethodRef(
            className: String,
            refName: String,
            descriptor: MethodType): ConstMethodRefInfo = {
        return _get[ConstMethodRefInfo](
                new ConstMethodRefInfo(
                        _owner,
                        getClass(className),
                        getNameAndType(refName, descriptor.descriptorString()),
                        descriptor))
    }

    def getInterfaceMethodRef(
            className: String,
            refName: String,
            descriptor: MethodType): ConstInterfaceMethodRefInfo = {
        return _get[ConstInterfaceMethodRefInfo](
                new ConstInterfaceMethodRefInfo(
                        _owner,
                        getClass(className),
                        getNameAndType(refName, descriptor.descriptorString()),
                        descriptor))
    }

    def getGetFieldMethodHandle(
            className: String,
            refName: String,
            descriptor: FieldType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewGetFieldMethodHandle(
                        _owner,
                        getFieldRef(className, refName, descriptor)))
    }

    def getGetStaticMethodHandle(
            className: String,
            refName: String,
            descriptor: FieldType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewGetStaticMethodHandle(
                        _owner,
                        getFieldRef(className, refName, descriptor)))
    }

    def getPutFieldMethodHandle(
            className: String,
            refName: String,
            descriptor: FieldType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewPutFieldMethodHandle(
                        _owner,
                        getFieldRef(className, refName, descriptor)))
    }

    def getPutStaticMethodHandle(
            className: String,
            refName: String,
            descriptor: FieldType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewPutStaticMethodHandle(
                        _owner,
                        getFieldRef(className, refName, descriptor)))
    }

    def getInvokeVirtualMethodHandle(
            className: String,
            refName: String,
            descriptor: MethodType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewInvokeVirtualMethodHandle(
                        _owner,
                        getMethodRef(className, refName, descriptor)))
    }

    def getNewInvokeVirtualMethodHandle(
            className: String,
            refName: String,
            descriptor: MethodType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewNewInvokeVirtualMethodHandle(
                        _owner,
                        getMethodRef(className, refName, descriptor)))
    }

    def getNewInvokeStaticMethodHandle(
            className: String,
            refName: String,
            descriptor: MethodType,
            isInterfaceMethod: Boolean): ConstMethodHandleInfo = {
        val ref = if (isInterfaceMethod) {
            getInterfaceMethodRef(className, refName, descriptor)
        } else {
            getMethodRef(className, refName, descriptor)
        }
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewInvokeStaticMethodHandle(
                        _owner,
                        ref))
    }

    def getNewInvokeSpecialMethodHandle(
            className: String,
            refName: String,
            descriptor: MethodType,
            isInterfaceMethod: Boolean): ConstMethodHandleInfo = {
        val ref = if (isInterfaceMethod) {
            getInterfaceMethodRef(className, refName, descriptor)
        } else {
            getMethodRef(className, refName, descriptor)
        }
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewInvokeSpecialMethodHandle(
                        _owner,
                        ref))
    }

    def getNewInvokeInterfaceMethodHandle(
            className: String,
            refName: String,
            descriptor: MethodType): ConstMethodHandleInfo = {
        return _get[ConstMethodHandleInfo](
                ConstMethodHandleInfo.NewInvokeInterfaceMethodHandle(
                        _owner,
                        getInterfaceMethodRef(className, refName, descriptor)))
    }

    def getInvokeDynamic(
            bootstrapMethodAttrIndex: Int,
            className: String,
            refName: String,
            descriptor: MethodType): ConstInvokeDynamicInfo = {
        return _get[ConstInvokeDynamicInfo](
                new ConstInvokeDynamicInfo(
                        _owner,
                        bootstrapMethodAttrIndex,
                        getNameAndType(className,
                                       descriptor.descriptorString()),
                        descriptor))
    }

    def _getByIndex[T <: ConstInfo : ClassTag](index: Int): T = {
        if (_tmpConstInfosByIndex == null) {
            throw new Exception(
                    "can only lookup by index during deserialization")
        }

        val info = _tmpConstInfosByIndex.get(index)
        if (info == null) {
            throw new Exception("entry not found")
        }

        val cls = implicitly[ClassTag[T]].runtimeClass
        info match {
            case t: T if cls.isInstance(t) => return t
            case _ => throw new Exception(
                    "unexpected const info type at " + index)
        }
    }

    def getByIndex(index: Int): ConstInfo = _getByIndex[ConstInfo](index)

    def getUtf8ByIndex(index: Int): ConstUtf8Info = {
        return _getByIndex[ConstUtf8Info](index)
    }

    def getIntegerByIndex(index: Int): ConstIntegerInfo = {
        return _getByIndex[ConstIntegerInfo](index)
    }

    def getLongByIndex(index: Int): ConstLongInfo = {
        return _getByIndex[ConstLongInfo](index)
    }

    def getFloatByIndex(index: Int): ConstFloatInfo = {
        return _getByIndex[ConstFloatInfo](index)
    }

    def getStringByIndex(index: Int): ConstStringInfo = {
        return _getByIndex[ConstStringInfo](index)
    }

    def getNameAndTypeByIndex(index: Int): ConstNameAndTypeInfo = {
        return _getByIndex[ConstNameAndTypeInfo](index)
    }

    def getRefByIndex(index: Int): ConstRefInfo = {
        return _getByIndex[ConstRefInfo](index)
    }

    def getFieldRefByIndex(index: Int): ConstFieldRefInfo = {
        return _getByIndex[ConstFieldRefInfo](index)
    }

    def getBaseMethodRefByIndex(index: Int): ConstBaseMethodRefInfo = {
        return _getByIndex[ConstBaseMethodRefInfo](index)
    }

    def getMethodRefByIndex(index: Int): ConstMethodRefInfo = {
        return _getByIndex[ConstMethodRefInfo](index)
    }

    def getInterfaceMethodRefByIndex(
            index: Int): ConstInterfaceMethodRefInfo = {
        return _getByIndex[ConstInterfaceMethodRefInfo](index)
    }

    def getClassByIndex(index: Int): ConstClassInfo = {
        return _getByIndex[ConstClassInfo](index)
    }

    def getMethodHandleByIndex(index: Int): ConstMethodHandleInfo = {
        return _getByIndex[ConstMethodHandleInfo](index)
    }

    def _assignIndex(): Int = {
        var nextIndex = 1
        for (info <- _constInfos.keySet()) {
            info.index = nextIndex
            nextIndex += info.indexSize()
        }

        if (nextIndex > Const.UINT16_MAX) {
            throw new Exception("const pool too large")
        }
        return nextIndex
    }

    def serialize(output: DataOutputStream) {
        _finalized = true
        _assignIndex()
        val last = _constInfos.lastEntry().getValue()
        output.writeShort(last.index + last.indexSize())

        for (info <- _constInfos.keySet()) {
            info.serialize(output)
        }
    }

    // TODO bind invoke dynamic bootstrap method attr index to reference
    def deserialize(input: DataInputStream) {
        if (!_constInfos.isEmpty()) {
            throw new Exception("deserializing into non-empty constant pool")
        }

        var constants = _parse(input)
        _generateIndexMap(constants)
        _bindConstReferences(constants)
        _populateAndDedup(constants)
    }

    def _parse(input: DataInputStream): Vector[ConstInfo] = {
        var result = new Vector[ConstInfo]()

        val constPoolCount = input.readUnsignedShort()

        var nextIndex = 1
        while (nextIndex < constPoolCount) {
            val tag = input.readByte()
            var info = tag match {
                case ConstInfo.UTF8 => new ConstUtf8Info(_owner)
                case ConstInfo.INTEGER => new ConstIntegerInfo(_owner)
                case ConstInfo.LONG => new ConstLongInfo(_owner)
                case ConstInfo.FLOAT => new ConstFloatInfo(_owner)
                case ConstInfo.DOUBLE => new ConstDoubleInfo(_owner)
                case ConstInfo.STRING => new ConstStringInfo(_owner)
                case ConstInfo.CLASS => new ConstClassInfo(_owner)
                case ConstInfo.NAME_AND_TYPE => new ConstNameAndTypeInfo(_owner)
                case ConstInfo.METHOD_TYPE => new ConstMethodTypeInfo(_owner)
                case ConstInfo.FIELD_REF => new ConstFieldRefInfo(_owner)
                case ConstInfo.METHOD_REF => new ConstMethodRefInfo(_owner)
                case ConstInfo.INTERFACE_METHOD_REF =>
                        new ConstInterfaceMethodRefInfo(_owner)
                case ConstInfo.METHOD_HANDLE =>
                        new ConstMethodHandleInfo(_owner)
                case ConstInfo.INVOKE_DYNAMIC =>
                        new ConstInvokeDynamicInfo(_owner)
                case _ => throw new Exception("Unknown const info type: " + tag)
            }
            info.index = nextIndex
            info.deserialize(tag, input)

            nextIndex += info.indexSize()

            result.add(info)
        }

        return result
    }

    def _generateIndexMap(constants: Collection[ConstInfo]) {
        _tmpConstInfosByIndex = new TreeMap[Int, ConstInfo]()

        for (info <- constants) {
            if (info.index < 1) {
                throw new Exception("invalid index")
            }
            if (_tmpConstInfosByIndex.containsKey(info.index)) {
                throw new Exception("duplicate index")
            }
            _tmpConstInfosByIndex.put(info.index, info)
        }
    }

    def _bindConstReferences(constants: Vector[ConstInfo]) {
        var tagInfos = new HashMap[Int, Vector[ConstInfo]]()
        for (t <- ConstInfo.tagTopoOrder) {
            tagInfos.put(t, new Vector[ConstInfo]())
        }

        for (info <- constants) {
            tagInfos.get(info.tag()).add(info)
        }

        for (t <- ConstInfo.tagTopoOrder) {
            for (info <- tagInfos.get(t)) {
                info.bindConstReferences()
            }
        }
    }

    def _populateAndDedup(constants: Vector[ConstInfo]) {
        for (info <- constants) {
            val first = _constInfos.get(info)
            if (first == null) {
                _constInfos.put(info, info)
            } else {
                _tmpConstInfosByIndex.put(info.index, first)
            }
        }
    }
}
