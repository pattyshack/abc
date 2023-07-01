import java.io.DataInputStream
import java.io.DataOutputStream
import java.io.EOFException
import java.util.Collection
import java.util.Vector

import scala.collection.JavaConversions._


object ClassInfo {
    val MAGIC = 0xcafebabe
}

class ClassInfo(name: String) extends AttributeOwner {
    def this() = this(null)

    var _minorVersion = 0
    var _majorVersion = 51  // jvm7

    var _constants = new ConstantPool(this)

    var _access = new ClassAccessFlags(this)

    var _thisClass: ConstClassInfo = null  // the current class
    if (name != null) {
        _thisClass = _constants.getClass(name)
    }
    var _superClass: ConstClassInfo = null  // if null then Object
    var _interfaces = new Vector[ConstClassInfo]()

    var _fields = new FieldPool(this)
    var _methods = new MethodPool(this)
    var _attributes = new ClassAttributes(this)

    def minorVersion(): Int = _minorVersion
    def majorVersion(): Int = _majorVersion

    def constants(): ConstantPool = _constants

    def access(): ClassAccessFlags = _access

    def thisClassName(): String = _thisClass.className()
    def superClassName(): String = {
        if (_superClass == null) {
            return null
        }
        return _superClass.className()
    }
    def interfaces(): Vector[ConstClassInfo] = _interfaces

    def fields(): Collection[FieldInfo] = _fields.fields()
    def methods(): Collection[MethodInfo] = _methods.methods()
    def attributes(): ClassAttributes = _attributes

    // performs various analysis / optimizaton.  NOTE: each subsection should
    // mark used constants as the last step.
    def analyze() {
        new InsertImplicitGotos().apply(this)
        new ShortenSimpleGotoChains().apply(this)
        new AdjustEntryPoints().apply(this)
        new RemoveUnreachableSegments().apply(this)
        new CheckJumpTargets().apply(this)

        // TODO: prune unused code blocks

        new DropUnsupportedAttributes().apply(this)
        new MarkAndSweepConstants().apply(this)

        // TODO: fix constant info index assignment (must be two pass)
    }

    def serialize(output: DataOutputStream) {
        output.writeInt(ClassInfo.MAGIC)

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
        if (input.readInt() != ClassInfo.MAGIC) {
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
        _methods.deserialize(input)
        _attributes.deserialize(input)

        var isEof = false
        try {
            input.readByte()
        } catch {
            case ex: EOFException => isEof = true
        }

        if (!isEof) {
            throw new Exception("Unparsed bytes at the end of class file")
        }
    }
}
