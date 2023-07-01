import java.io.DataInputStream
import java.io.DataOutput


//
// nop
// stack: ... -> ...
//
class Nop(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.NOP, "nop") {
}

//
// aconst_null
// stack: ... -> ..., null
//
class AconstNull(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.ACONST_NULL, "aconst_null") {
}

//
// "PushI" <value>
// stack: ... -> ..., value
//
class PushI(owner: AttributeOwner, v: Int) extends Operation(owner) {
    def this(owner: AttributeOwner) = this(owner, 0)

    var value = v
    var _constInt: ConstIntegerInfo = null
    if (value < -32768 || value > 32767) {
        _constInt = owner.constants().getInteger(value)
    }

    def serialize(output: DataOutput) {
        value match {
            case -1 => output.write(OpCode.ICONST_M1)
            case 0 => output.write(OpCode.ICONST_0)
            case 1 => output.write(OpCode.ICONST_1)
            case 2 => output.write(OpCode.ICONST_2)
            case 3 => output.write(OpCode.ICONST_3)
            case 4 => output.write(OpCode.ICONST_4)
            case 5 => output.write(OpCode.ICONST_5)
            case _ => {
                if (Const.INT8_MIN <= value && value <= Const.INT8_MAX) {
                    output.write(OpCode.BIPUSH)
                    output.writeByte(value)
                } else if (Const.INT16_MIN <= value &&
                           value <= Const.INT16_MAX) {
                    output.write(OpCode.SIPUSH)
                    output.writeShort(value)
                } else {
                    if (_constInt.index <= Const.UINT8_MAX) {
                        output.write(OpCode.LDC)
                        output.writeByte(_constInt.index)
                    } else {
                        output.write(OpCode.LDC_W)
                        output.writeShort(_constInt.index)
                    }
                }
            }
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        value = opCode match {
            case OpCode.ICONST_M1 => -1
            case OpCode.ICONST_0 => 0
            case OpCode.ICONST_1 => 1
            case OpCode.ICONST_2 => 2
            case OpCode.ICONST_3 => 3
            case OpCode.ICONST_4 => 4
            case OpCode.ICONST_5 => 5
            case OpCode.BIPUSH => input.readByte()
            case OpCode.SIPUSH => input.readShort()
            case _ => throw new Exception(
                    "cannot directly deserialize \"iconst\"")
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": \"iconst\" " + value + "\n"
    }
}

//
// "PushL" <value>
// stack: ... -> ..., value
//
class PushL(owner: AttributeOwner, v: Long) extends Operation(owner) {
    def this(owner: AttributeOwner) = this(owner, 0)

    var value = v
    var _constLong: ConstLongInfo = null
    if (value != 0 && value != 1) {
        _constLong = owner.constants().getLong(value)
    }

    def serialize(output: DataOutput) {
        if (value == 0) {
            output.write(OpCode.LCONST_0)
        } else if (value == 1) {
            output.write(OpCode.LCONST_1)
        } else {
            output.write(OpCode.LDC2_W)
            output.writeShort(_constLong.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        value = opCode match {
            case OpCode.LCONST_0 => 0
            case OpCode.LCONST_1 => 1
            case _ => throw new Exception(
                    "cannot directly deserialize \"lconst\"")
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": \"lconst\" " + value + "\n"
    }
}

//
// "PushF" <value>
// stack: ... -> ..., value
//
class PushF(owner: AttributeOwner, v: Float) extends Operation(owner) {
    def this(owner: AttributeOwner) = this(owner, 0)

    var value = v
    var _constFloat: ConstFloatInfo = null
    if (value != 0 && value != 1 && value != 2) {
        _constFloat = owner.constants().getFloat(value)
    }

    def serialize(output: DataOutput) {
        if (value == 0) {
            output.write(OpCode.FCONST_0)
        } else if (value == 1) {
            output.write(OpCode.FCONST_1)
        } else if (value == 2) {
            output.write(OpCode.FCONST_2)
        } else {
            if (_constFloat.index <= Const.UINT8_MAX) {
                output.write(OpCode.LDC)
                output.writeByte(_constFloat.index)
            } else {
                output.write(OpCode.LDC_W)
                output.writeShort(_constFloat.index)
            }
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        value = opCode match {
            case OpCode.FCONST_0 => 0
            case OpCode.FCONST_1 => 1
            case OpCode.FCONST_2 => 2
            case _ => throw new Exception(
                    "cannot directly deserialize \"fconst\"")
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": \"fconst\" " + value + "\n"
    }
}

//
// "PushD" <value>
// stack: ... -> ..., value
//
class PushD(owner: AttributeOwner, v: Double) extends Operation(owner) {
    def this(owner: AttributeOwner) = this(owner, 0)

    var value = v
    var _constDouble: ConstDoubleInfo = null
    if (value != 0 && value != 1) {
        _constDouble = owner.constants().getDouble(value)
    }

    def serialize(output: DataOutput) {
        if (value == 0) {
            output.write(OpCode.DCONST_0)
        } else if (value == 1) {
            output.write(OpCode.DCONST_1)
        } else {
            output.write(OpCode.LDC2_W)
            output.writeShort(_constDouble.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        value = opCode match {
            case OpCode.DCONST_0 => 0
            case OpCode.DCONST_1 => 1
            case _ => throw new Exception(
                    "cannot directly deserialize \"dconst\"")
        }
    }

    def debugString(indent: String): String = {
        indent + _pcLine() + ": \"dconst\" " + value + "\n"
    }
}

//
// "PushString" <value>
// stack: ... -> ..., value
//
class PushString(owner: AttributeOwner, v: String) extends Operation(owner) {
    var _constString: ConstStringInfo = null
    if (v != null) {
        _constString = owner.constants().getString(v)
    }

    def value(): String = _constString.value()

    def serialize(output: DataOutput) {
        if (_constString.index <= Const.UINT8_MAX) {
            output.write(OpCode.LDC)
            output.writeByte(_constString.index)
        } else {
            output.write(OpCode.LDC_W)
            output.writeShort(_constString.index)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize \"sconst\"")
    }

    def debugString(indent: String): String = {
        indent + _pcLine() + ": \"sconst\" " + _constString.debugString() + "\n"
    }
}

class Ldc(owner: AttributeOwner, v: ConstInfo) extends ByteOperandOp(
        owner,
        OpCode.LDC,
        "ldc",
        false,  // signed
        0) {
    def this(owner: AttributeOwner) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutput) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstIntegerInfo => new PushI(_owner, c.value())
            case c: ConstFloatInfo => new PushF(_owner, c.value())
            case c: ConstStringInfo => new PushString(_owner, c.value())
            // TODO: support class / method type / method handle
            case _ => this
        }
    }
}

class LdcW(owner: AttributeOwner, v: ConstInfo) extends ShortOperandOp(
        owner,
        OpCode.LDC_W,
        "ldc_w",
        false,  // signed
        0) {
    def this(owner: AttributeOwner) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutput) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstIntegerInfo => new PushI(_owner, c.value())
            case c: ConstFloatInfo => new PushF(_owner, c.value())
            case c: ConstStringInfo => new PushString(_owner, c.value())
            // TODO: support class / method type / method handle
            case _ => this
        }
    }
}

class Ldc2W(owner: AttributeOwner, v: ConstInfo) extends ShortOperandOp(
        owner,
        OpCode.LDC2_W,
        "ldc2_w",
        false,  // signed
        0) {
    def this(owner: AttributeOwner) = this(owner, null)

    var _const = v

    override def serialize(output: DataOutput) {
        operand = _const.index
        super.serialize(output)
    }

    override def deserialize(
            startAddress: Int,
            opCode: Int,
            input: DataInputStream) {
        super.deserialize(startAddress, opCode, input)
        _const = _owner.constants().getByIndex(operand)
    }

    override def canonicalForm(): Operation = {
        return _const match {
            case c: ConstDoubleInfo => new PushD(_owner, c.value())
            case c: ConstLongInfo => new PushL(_owner, c.value())
            case _ => throw new Exception(
                    "Unexpected constant type: " + _const.typeName())
        }
    }
}
