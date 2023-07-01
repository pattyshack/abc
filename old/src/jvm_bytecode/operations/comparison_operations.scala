import java.io.DataInputStream
import java.io.DataOutput
import java.util.TreeMap


// stack: ..., long value1, long value2 -> int result
class Lcmp(owner: AttributeOwner)
        extends NoOperandOp(owner, OpCode.LCMP, "lcmp") {
}

// stack: ..., float value1, float value2 -> int result
abstract class Fcmp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

// stack: ..., double value1, double value2 -> int result
abstract class Dcmp(owner: AttributeOwner, opCode: Int, mnemonic: String)
        extends NoOperandOp(owner, opCode, mnemonic) {
}

class Fcmpg(owner: AttributeOwner) extends Fcmp(owner, OpCode.FCMPG, "fcmpg") {
}

class Fcmpl(owner: AttributeOwner) extends Fcmp(owner, OpCode.FCMPL, "fcmpl") {
}

class Dcmpg(owner: AttributeOwner) extends Dcmp(owner, OpCode.DCMPG, "dcmpg") {
}

class Dcmpl(owner: AttributeOwner) extends Dcmp(owner, OpCode.DCMPL, "dcmpl") {
}

abstract class IfBaseOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var _ifBranch: CodeBlock = null
    if (ifBranch != null) {
        _ifBranch = ifBranch.getEntryBlock()
    }
    var _elseBranch: CodeBlock = null
    if (elseBranch != null) {
        _elseBranch = elseBranch.getEntryBlock()
    }

    // only used during deserialization
    var _tmpOffset = 0

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
        _writeShortOffset(_ifBranch, output)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        _tmpOffset = input.readShort()
    }

    override def bindBlockRefs(table: TreeMap[Int, CodeBlock]) {
        _ifBranch = table.get(pc + _tmpOffset)
        if (_ifBranch == null) {
            throw new Exception("can't find if block")
        }

        val elseEntry = table.higherEntry(pc)
        if (elseEntry == null) {
            throw new Exception("can't find else block")
        }
        _elseBranch = elseEntry.getValue()
    }

    def debugString(indent: String): String = {
        var ifPc = "???"
        if (_ifBranch != null) {
            ifPc = "" + _ifBranch.pc
        }
        var elsePc = "???"
        if (_elseBranch != null) {
            elsePc = "" + _elseBranch.pc
        }
        return indent + _pcLine() + ": " + _mnemonic + " " + ifPc +
                " " + elsePc + "\n"
    }
}

// stack: ..., int value -> ... (compare against 0)
abstract class IfIOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                ifBranch,
                elseBranch) {
}

// stack: ..., int value1, int value2 -> ...
abstract class IfCmpIOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                ifBranch,
                elseBranch) {
}

// stack: ..., ref value -> ...
abstract class IfAOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                ifBranch,
                elseBranch) {
}

// stack: ..., ref value1, ref value2 -> ...
abstract class IfCmpAOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfBaseOp(
                owner,
                opCode,
                mnemonic,
                ifBranch,
                elseBranch) {
}

//
// int conditional branching
//

class Ifeq(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFEQ,
                "ifeq",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifne(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFNE,
                "ifne",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Iflt(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFLT,
                "iflt",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifle(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFLE,
                "ifle",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifgt(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFGT,
                "ifgt",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifge(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfIOp(
                owner,
                OpCode.IFGE,
                "ifge",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmpeq(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPEQ,
                "if_icmpeq",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmpne(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPNE,
                "if_icmpne",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmplt(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPLT,
                "if_icmplt",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmpge(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPGE,
                "if_icmpge",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmpgt(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPGT,
                "if_icmpgt",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfIcmple(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpIOp(
                owner,
                OpCode.IF_ICMPLE,
                "if_icmple",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

//
// ref conditional branching
//

class IfAcmpeq(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpAOp(
                owner,
                OpCode.IF_ACMPEQ,
                "if_acmpeq",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class IfAcmpne(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfCmpAOp(
                owner,
                OpCode.IF_ACMPNE,
                "if_acmpne",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifnonnull(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment) extends IfAOp(
                owner,
                OpCode.IFNONNULL,
                "ifnonnull",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}

class Ifnull(
        owner: AttributeOwner,
        ifBranch: CodeSegment,
        elseBranch: CodeSegment)
        extends IfAOp(
                owner,
                OpCode.IFNULL,
                "ifnull",
                ifBranch,
                elseBranch) {
    def this(owner: AttributeOwner) = this(owner, null, null)
}
