import java.io.DataInputStream
import java.io.DataOutput

abstract class Load(
        owner: AttributeOwner,
        opCode: Int,
        shortOpCodeStart: Int,
        mnemonic: String,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    val _shortOpCodeStart = shortOpCodeStart
    var index = v

    def serialize(output: DataOutput) {
        index match {
            case 0 => output.writeByte(_shortOpCodeStart)
            case 1 => output.writeByte(_shortOpCodeStart + 1)
            case 2 => output.writeByte(_shortOpCodeStart + 2)
            case 3 => output.writeByte(_shortOpCodeStart + 3)
            case _ => {
                if (index <= Const.UINT8_MAX) {
                    output.writeByte(_opCode)
                    output.writeByte(index)
                } else {
                    output.writeByte(OpCode.WIDE)
                    output.writeByte(_opCode)
                    output.writeShort(index)
                }
            }
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (_shortOpCodeStart <= opCode && opCode <= _shortOpCodeStart + 3) {
            index = opCode - _shortOpCodeStart
        } else if (opCode == _opCode) {
            index = input.readUnsignedByte()
        } else {
            throw new Exception("Unexpected op code " + opCode)
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": \"" + _mnemonic + "\" " + index + "\n"
    }
}

//
// iload <local var index>
// stack: ... -> ..., value
//
class LoadI(owner: AttributeOwner, index: Int)
        extends Load(owner, OpCode.ILOAD, OpCode.ILOAD_0, "iload", index) {

    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// lload <local var index>
// stack: ... -> ..., value
//
class LoadL(owner: AttributeOwner, index: Int)
        extends Load(owner, OpCode.LLOAD, OpCode.LLOAD_0, "lload", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// fload <local var index>
// stack: ... -> ..., value
//
class LoadF(owner: AttributeOwner, index: Int)
        extends Load(owner, OpCode.FLOAD, OpCode.FLOAD_0, "fload", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// dload <local var index>
// stack: ... -> ..., value
//
class LoadD(owner: AttributeOwner, index: Int)
        extends Load(owner, OpCode.DLOAD, OpCode.DLOAD_0, "dload", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// aload <local var index>
// stack: ... -> ..., value
//
class LoadA(owner: AttributeOwner, index: Int)
        extends Load(owner, OpCode.ALOAD, OpCode.ALOAD_0, "aload", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

abstract class LoadFromArray(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

abstract class LoadFromBaseIntArray(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends LoadFromArray(owner, opCode, mnemonic) {
}

//
// iaload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromIArray(owner: AttributeOwner)
        extends LoadFromBaseIntArray(owner, OpCode.IALOAD, "iaload") {
}

//
// laload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromLArray(owner: AttributeOwner)
        extends LoadFromArray(owner, OpCode.LALOAD, "laload") {
}

//
// faload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromFArray(owner: AttributeOwner)
        extends LoadFromArray(owner, OpCode.FALOAD, "faload") {
}

//
// daload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromDArray(owner: AttributeOwner)
        extends LoadFromArray(owner, OpCode.DALOAD, "daload") {
}

//
// aaload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromAArray(owner: AttributeOwner)
        extends LoadFromArray(owner, OpCode.AALOAD, "aaload") {
}

//
// baload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromBArray(owner: AttributeOwner)
        extends LoadFromBaseIntArray(owner, OpCode.BALOAD, "baload") {
}

//
// caload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromCArray(owner: AttributeOwner)
        extends LoadFromBaseIntArray(owner, OpCode.CALOAD, "caload") {
}

//
// saload
// stack: ..., arrayref, index -> ..., value
//
class LoadFromSArray(owner: AttributeOwner)
        extends LoadFromBaseIntArray(owner, OpCode.SALOAD, "saload") {
}
