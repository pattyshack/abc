import java.io.DataInputStream
import java.io.DataOutputStream


//
// pop
// stack: ..., value -> ...
//
class Pop(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.POP, "pop") {
}

//
// pop2
// stack: ..., value1, value2 -> ... (category 1 values)
// or
// stack: ..., value -> ... (category 2 value)
// (see table 2.3  / pg 29)
//
class Pop2(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.POP2, "pop2") {
}

//
// dup
// stack: ..., value -> ..., value, value
//
class Dup(owner: AttributeOwner) extends NoOperandOp(owner, OpCode.DUP, "dup") {
}

//
// dup_x1
// stack: ..., value2, value1 -> ..., value1, value2, value1
// (value1 and value2 must be category 1; see table 2.3 / pg 29)
//
class DupX1(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.DUP_X1, "dup_x1") {
}

//
// dup_x2
// stack: ..., value2, value1 -> ..., value1, value2, value1
// (value1 must be category 1, and value2 must be category 2;
// see table 2.3 / pg 29)
//
class DupX2(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.DUP_X2, "dup_x2") {
}

//
// dup2
// stack: ..., value1, value2 ->
//          ..., value1, value2, value1, value2 (category 1 values)
// or
// stack: ..., value -> ..., value, value (category 2 value)
// (see table 2.3  / pg 29)
//
class Dup2(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.DUP2, "dup2") {
}

//
// dup2_x1
// stack: ..., value3, value2, value1 ->
//          ..., value2, value1, value3, value2, value1 (category 1 values)
// or
// stack: ..., value2, value1 -> ..., value1, value2, value1
//          (value1 is category 2, value2 is category 1)
// (see table 2.3  / pg 29)
//
class Dup2X1(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.DUP2_X1, "dup2_x1") {
}

//
// dup2_x2
// stack: ..., value4, value3, value2, value1 ->
//          ..., value2, value1, value4, value3, value2, value1
//          (category 1 values)
// or
// stack: ..., value3, value2, value1 -> ..., value1, value3, value2, value1
//          (value1 is category 2, value2/value3 are category 1)
// (see table 2.3  / pg 29)
//
class Dup2X2(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.DUP2_X2, "dup2_x2") {
}

//
// swap
// stack: ..., value1, value2 -> ..., value2, value1
//
class Swap(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.SWAP, "swap") {
}
