import java.io.DataInputStream
import java.io.DataOutputStream


class FieldInfo(
        c: ClassInfo,
        n: String,
        f: FieldType) extends AttributeOwner {
    def this(c: ClassInfo) = this(c, null, null)

    var _owner = c

    var _access = new FieldAccessFlags(this)
    var _name: ConstUtf8Info = null
    if (n != null) {
        _name = _owner.constants.getUtf8(n)
    }

    var _fieldType: FieldType = f
    var _fieldTypeString: ConstUtf8Info = null
    if (f != null) {
        _fieldTypeString = constants().getUtf8(_fieldType.descriptorString())
    }

    var _attributes = new FieldAttributes(this)

    def constants(): ConstantPool = _owner.constants()

    def access(): FieldAccessFlags = _access
    def name(): String = _name.value()
    def fieldTypeString(): String = _fieldTypeString.value()
    override def fieldType(): FieldType = _fieldType
    def attributes(): FieldAttributes = _attributes

    def serialize(output: DataOutputStream) {
        _access.serialize(output)
        output.writeShort(_name.index)
        output.writeShort(_fieldTypeString.index)
        _attributes.serialize(output)
    }

    def deserialize(input: DataInputStream, constants: ConstantPool) {
        _access.deserialize(input)
        _name = constants.getUtf8ByIndex(input.readUnsignedShort())

        _fieldTypeString = constants.getUtf8ByIndex(input.readUnsignedShort())
        var parser = new DescriptorParser(_fieldTypeString.value())
        _fieldType = parser.parseFieldDescriptor()

        _attributes.deserialize(input)
    }
}
