import java.io.DataInputStream
import java.io.DataOutputStream


abstract class TruncateI(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends NoOperandOp(owner, opCode, mnemonic) {
}

class I2l(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.I2L, "i2l") {
}

class I2f(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.I2F, "i2f") {
}

class I2d(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.I2D, "i2d") {
}

class I2b(owner: AttributeOwner) extends TruncateI(owner, OpCode.I2B, "i2b") {
}

class I2c(owner: AttributeOwner) extends TruncateI(owner, OpCode.I2C, "i2c") {
}

class I2s(owner: AttributeOwner) extends TruncateI(owner, OpCode.I2S, "i2s") {
}

class L2i(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.L2I, "l2i") {
}

class L2f(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.L2F, "l2f") {
}

class L2d(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.L2D, "l2d") {
}

class F2i(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.F2I, "f2i") {
}

class F2l(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.F2L, "f2l") {
}

class F2d(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.F2D, "f2d") {
}

class D2i(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.D2I, "d2i") {
}

class D2l(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.D2L, "d2l") {
}

class D2f(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.D2F, "d2f") {
}
