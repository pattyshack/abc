import java.io.DataInputStream
import java.io.DataOutput

abstract class Store(
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
// istore <local var index>
// stack: ..., value -> ...
//
class StoreI(owner: AttributeOwner, index: Int)
        extends Store(owner, OpCode.ISTORE, OpCode.ISTORE_0, "istore", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// lstore <local var index>
// stack: ..., value -> ...
//
class StoreL(owner: AttributeOwner, index: Int)
        extends Store(owner, OpCode.LSTORE, OpCode.LSTORE_0, "lstore", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// fstore <local var index>
// stack: ..., value -> ...
//
class StoreF(owner: AttributeOwner, index: Int)
        extends Store(owner, OpCode.FSTORE, OpCode.FSTORE_0, "fstore", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// dstore <local var index>
// stack: ..., value -> ...
//
class StoreD(owner: AttributeOwner, index: Int)
        extends Store(owner, OpCode.DSTORE, OpCode.DSTORE_0, "dstore", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

//
// astore <local var index>
// stack: ..., value -> ...
//
class StoreA(owner: AttributeOwner, index: Int)
        extends Store(owner, OpCode.ASTORE, OpCode.ASTORE_0, "astore", index) {
    def this(owner: AttributeOwner) = this(owner, 0)
}

abstract class StoreIntoArray(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

abstract class StoreIntoBaseIntArray(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends StoreIntoArray(owner, opCode, mnemonic) {
}

//
// iastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoIArray(owner: AttributeOwner)
        extends StoreIntoBaseIntArray(owner, OpCode.IASTORE, "iastore") {
}

//
// lastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoLArray(owner: AttributeOwner)
        extends StoreIntoArray(owner, OpCode.LASTORE, "lastore") {
}

//
// fastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoFArray(owner: AttributeOwner)
        extends StoreIntoArray(owner, OpCode.FASTORE, "fastore") {
}

//
// dastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoDArray(owner: AttributeOwner)
        extends StoreIntoArray(owner, OpCode.DASTORE, "dastore") {
}

//
// aastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoAArray(owner: AttributeOwner)
        extends StoreIntoArray(owner, OpCode.AASTORE, "aastore") {
}

//
// bastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoBArray(owner: AttributeOwner)
        extends StoreIntoBaseIntArray(owner, OpCode.BASTORE, "bastore") {
}

//
// castore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoCArray(owner: AttributeOwner)
        extends StoreIntoBaseIntArray(owner, OpCode.CASTORE, "castore") {
}

//
// sastore
// stack: ..., arrayref, index, value -> ...
//
class StoreIntoSArray(owner: AttributeOwner)
        extends StoreIntoBaseIntArray(owner, OpCode.SASTORE, "sastore") {
}
