import java.io.DataInputStream
import java.io.DataOutput


// stack: ..., value1, value2 -> ..., result
abstract class BinaryIOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryIOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryLOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryLOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., (long) value1, (int) value2 -> (long) result
abstract class ShiftLOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryFOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryFOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1, value2 -> ..., result
abstract class BinaryDOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., value1 -> ..., result
abstract class UnaryDOp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

//
// int operations
//

class Iadd(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IADD, "iadd") {
}

class Isub(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.ISUB, "isub") {
}

class Imul(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IMUL, "imul") {
}

class Idiv(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IDIV, "idiv") {
}

class Irem(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IREM, "irem") {
}

class Ineg(owner: AttributeOwner)
        extends UnaryIOp(owner, OpCode.INEG, "ineg") {
}

class Ishl(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.ISHL, "ishl") {
}

class Ishr(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.ISHR, "ishr") {
}

class Iushr(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IUSHR, "iushr") {
}

class Iand(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IAND, "iand") {
}

class Ior(owner: AttributeOwner) extends BinaryIOp(owner, OpCode.IOR, "ior") {
}

class Ixor(owner: AttributeOwner)
        extends BinaryIOp(owner, OpCode.IXOR, "ixor") {
}

// iinc <local var index> <const int>
// increment local variable by const without modifying stack
class Iinc(owner: AttributeOwner, index: Int, v: Int)
        extends TwoByteOperandsOp(
                owner,
                OpCode.IINC,
                "iinc",
                false,  // signed
                index,
                true,  // signed
                v) {
    def this(owner: AttributeOwner) = this(owner, -1, 0)

    override def serialize(output: DataOutput) {
        if (operand1 <= Const.UINT8_MAX &&
            (Const.INT8_MIN <= operand2 && operand2 <= Const.INT8_MAX)) {
            super.serialize(output)
        } else {
            output.writeByte(OpCode.WIDE)
            output.writeByte(OpCode.IINC)
            output.writeShort(operand1)
            output.writeShort(operand2)
        }
    }
}

//
// long operations
//

class Ladd(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LADD, "ladd") {
}

class Lsub(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LSUB, "lsub") {
}

class Lmul(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LMUL, "lmul") {
}

class Ldiv(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LDIV, "ldiv") {
}

class Lrem(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LREM, "lrem") {
}

class Lneg(owner: AttributeOwner)
        extends UnaryLOp(owner, OpCode.LNEG, "lneg") {
}

class Lshl(owner: AttributeOwner) extends ShiftLOp(owner, OpCode.LSHL, "lshl") {
}

class Lshr(owner: AttributeOwner) extends ShiftLOp(owner, OpCode.LSHR, "lshr") {
}

class Lushr(owner: AttributeOwner)
        extends ShiftLOp(owner, OpCode.LUSHR, "lushr") {
}

class Land(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LAND, "land") {
}

class Lor(owner: AttributeOwner) extends BinaryLOp(owner, OpCode.LOR, "lor") {
}

class Lxor(owner: AttributeOwner)
        extends BinaryLOp(owner, OpCode.LXOR, "lxor") {
}

//
// float operations
//

class Fadd(owner: AttributeOwner)
        extends BinaryFOp(owner, OpCode.FADD, "fadd") {
}

class Fsub(owner: AttributeOwner)
        extends BinaryFOp(owner, OpCode.FSUB, "fsub") {
}

class Fmul(owner: AttributeOwner)
        extends BinaryFOp(owner, OpCode.FMUL, "fmul") {
}

class Fdiv(owner: AttributeOwner)
        extends BinaryFOp(owner, OpCode.FDIV, "fdiv") {
}

class Frem(owner: AttributeOwner)
        extends BinaryFOp(owner, OpCode.FREM, "frem") {
}

class Fneg(owner: AttributeOwner) extends UnaryFOp(owner, OpCode.FNEG, "fneg") {
}

//
// double operations
//

class Dadd(owner: AttributeOwner)
        extends BinaryDOp(owner, OpCode.DADD, "dadd") {
}

class Dsub(owner: AttributeOwner)
        extends BinaryDOp(owner, OpCode.DSUB, "dsub") {
}

class Dmul(owner: AttributeOwner)
        extends BinaryDOp(owner, OpCode.DMUL, "dmul") {
}

class Ddiv(owner: AttributeOwner)
        extends BinaryDOp(owner, OpCode.DDIV, "ddiv") {
}

class Drem(owner: AttributeOwner)
        extends BinaryDOp(owner, OpCode.DREM, "drem") {
}

class Dneg(owner: AttributeOwner) extends UnaryDOp(owner, OpCode.DNEG, "dneg") {
}

