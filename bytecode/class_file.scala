import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


object ClassFile {
    val MAGIC = 0xcafebabe
}

class ClassFile {
    var _minorVersion = 0
    var _majorVersion = 51  // jvm7

    var _constants = new ConstantPool()

    var _access = new ClassAccessFlags()

    var _thisClass: ConstClassInfo = null  // the current class
    var _superClass: ConstClassInfo = null  // null if Object; non-null otherwise
    var _interfaces = new Vector[ConstClassInfo]()

    var _fields = new FieldPool(_constants)
    var _methods = new MethodPool()
    var _attributes = new ClassAttributes()

    def minorVersion(): Int = _minorVersion
    def majorVersion(): Int = _majorVersion
    def constants(): ConstantPool = _constants
    def access(): ClassAccessFlags = _access
    def thisClass(): ConstClassInfo = _thisClass
    def superClass(): ConstClassInfo = _superClass
    def interfaces(): Vector[ConstClassInfo] = _interfaces
    def fields(): FieldPool = _fields
    def methods(): MethodPool = _methods
    def attribute(): ClassAttributes = _attributes

    def serialize(output: DataOutputStream) {
        output.writeInt(ClassFile.MAGIC)

        output.writeShort(_minorVersion)
        output.writeShort(_majorVersion)

        _constants.serialize(output)

        _access.serialize(output)

        output.writeShort(_thisClass.index)
        var _superClassIndex = 0
        if (_superClass != null) {  // i.e., not Object
            _superClassIndex = _superClass.index
        }
        output.writeShort(_superClassIndex)
        output.writeShort(_interfaces.length)
        for (iface <- _interfaces) {
            output.writeShort(iface.index)
        }

        _fields.serialize(output)
        _methods.serialize(output)
        _attributes.serialize(output)
    }

    def deserialize(input: DataInputStream) {
        if (input.readInt() != ClassFile.MAGIC) {
            throw new Exception("Invalid magic")
        }

        _minorVersion = input.readUnsignedShort()
        _majorVersion = input.readUnsignedShort()

        _constants.deserialize(input)

        _access.deserialize(input)

        val _thisClassIndex = input.readUnsignedShort()
        _thisClass = _constants.getClassByIndex(_thisClassIndex)

        val _superClassIndex = input.readUnsignedShort()
        if (_superClassIndex == 0) {
            _superClass = null
        } else {
            _superClass = _constants.getClassByIndex(_superClassIndex)
        }

        val _interfacesCount = input.readUnsignedShort()
        for (_ <- 1 to _interfacesCount) {
            _interfaces.add(
                    _constants.getClassByIndex(input.readUnsignedShort()))
        }

        _fields.deserialize(input)
        /* TODO
        _methods.deserialize(input, _constants)
        _attributes.deserialize(input, _constants)
        */
    }
}
