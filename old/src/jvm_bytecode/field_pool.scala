import java.io.DataOutputStream
import java.io.DataInputStream
import java.util.Collection
import java.util.TreeMap

import scala.collection.JavaConversions._


class FieldPool(c: ClassInfo) {
    var _owner = c

    var _fields = new TreeMap[String, FieldInfo]()

    var _constants: ConstantPool = c.constants()

    def fields(): Collection[FieldInfo] = _fields.values()

    def _add(field: FieldInfo) {
        if (_fields.containsKey(field.name())) {
            throw new Exception("adding duplicate field: " + field.name())
        }
        _fields.put(field.name(), field)
    }

    def add(name: String, fieldType: FieldType): FieldInfo = {
        val field = new FieldInfo(_owner, name, fieldType)
        _add(field)
        return field
    }

    def get(name: String): FieldInfo = {
        val f = _fields.get(name)
        if (f == null) {
            throw new Exception("missing field: " + name)
        }
        return f
    }

    def serialize(output: DataOutputStream) {
        output.writeShort(_fields.size())
        for (field <- fields()) {
            field.serialize(output)
        }
    }

    def deserialize(input: DataInputStream) {
        if (!_fields.isEmpty()) {
            throw new Exception("deserializing into non-empty field pool")
        }

        val fieldCount = input.readUnsignedShort()
        for (i <- 1 to fieldCount) {
            var field = new FieldInfo(_owner)
            field.deserialize(input, _constants)
            _add(field)
        }
    }
}
