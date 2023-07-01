import java.io.ByteArrayInputStream
import java.io.DataInputStream
import java.io.DataOutput
import java.util.TreeMap
import java.util.Vector


abstract class Operation(o: AttributeOwner) {
    var _owner: AttributeOwner = o
    var line = -1
    var pc = -1

    def canonicalForm(): Operation = this

    def serialize(output: DataOutput)

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream)

    // used for control flow graph reconstruction.
    def bindBlockRefs(table: TreeMap[Int, CodeBlock]) {}

    def debugString(indent: String): String

    def _writeShortOffset(target: CodeBlock, output: DataOutput) {
        val offset = target.pc - pc
        if (offset < Const.INT16_MIN || offset > Const.INT16_MAX) {
            throw new Exception("code too large (goto_w not supported)")
        }
        output.writeShort(offset)
    }

    def _pcLine(): String = "" + pc + "(" + line + ")"
}

// operations of the form: <op code>
abstract class NoOperandOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": " + _mnemonic + "\n"
    }
}

// operations of the form: <op code> <byte operand>
abstract class ByteOperandOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        signed: Boolean,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var isSigned = signed
    var operand = v

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
        output.writeByte(operand)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        if (isSigned) {
            operand = input.readByte()
        } else {
            operand = input.readUnsignedByte()
        }
    }

    def debugString(indent: String): String = {
        indent + _pcLine() + ": " + _mnemonic + " " + operand + "\n"
    }
}

// operations of the form: <op code> <byte operand1> <byte operand2>
abstract class TwoByteOperandsOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        signed1: Boolean,
        v1: Int,
        signed2: Boolean,
        v2: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var isSigned1 = signed1
    var operand1 = v1
    var isSigned2 = signed2
    var operand2 = v2

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
        output.writeByte(operand1)
        output.writeByte(operand2)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        if (isSigned1) {
            operand1 = input.readByte()
        } else {
            operand1 = input.readUnsignedByte()
        }

        if (isSigned2) {
            operand2 = input.readByte()
        } else {
            operand2 = input.readUnsignedByte()
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": " + _mnemonic + " " +
                operand1 + " " + operand2 + "\n"
    }
}

// operations of the form: <op code> <short operand>
class ShortOperandOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        signed: Boolean,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var isSigned = signed
    var operand = v

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
        output.writeShort(operand)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }

        if (isSigned) {
            operand = input.readShort()
        } else {
            operand = input.readUnsignedShort()
        }
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": " + _mnemonic + " " + operand + "\n"
    }
}

// operations of the form: <op code> <int operand>
class IntOperandOp(
        owner: AttributeOwner,
        opCode: Int,
        mnemonic: String,
        v: Int) extends Operation(owner) {
    val _opCode = opCode
    val _mnemonic = mnemonic
    var operand = v

    def serialize(output: DataOutput) {
        output.writeByte(_opCode)
        output.writeInt(operand)
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        if (opCode != _opCode) {
            throw new Exception("Unexpected op-code: " + opCode)
        }
        operand = input.readInt()
    }

    def debugString(indent: String): String = {
        return indent + _pcLine() + ": " + _mnemonic + " " + operand + "\n"
    }
}

object Operation {
    def parseWide(owner: AttributeOwner, input: DataInputStream): Operation = {
        val opCode = input.readUnsignedByte()
        val index = input.readUnsignedShort()

        return opCode match {
            case OpCode.ILOAD => new LoadI(owner, index)
            case OpCode.FLOAD => new LoadF(owner, index)
            case OpCode.ALOAD => new LoadF(owner, index)
            case OpCode.LLOAD => new LoadL(owner, index)
            case OpCode.DLOAD => new LoadD(owner, index)
            case OpCode.ISTORE => new StoreI(owner, index)
            case OpCode.FSTORE => new StoreF(owner, index)
            case OpCode.ASTORE => new StoreA(owner, index)
            case OpCode.LSTORE => new StoreL(owner, index)
            case OpCode.DSTORE => new StoreD(owner, index)
            case OpCode.RET => new Ret(owner, index)
            case OpCode.IINC => new Iinc(owner, index, input.readShort())
        }
    }

    def deserialize(owner: AttributeOwner,
                    codeBytes: Array[Byte]): Vector[Operation] = {
        var byteStream = new ByteArrayInputStream(codeBytes)
        var dataStream = new DataInputStream(byteStream)

        var result = new Vector[Operation]()

        while (byteStream.available() > 0) {
            val pc = codeBytes.length - byteStream.available()
            result.add(parse(owner, pc, dataStream))
        }

        return result
    }

    def parse(owner: AttributeOwner,
              startAddress: Int,
              input: DataInputStream): Operation = {

        val opCode = input.readUnsignedByte()
        var operation: Operation = opCode match {
            case 0 => new Nop(owner)
            case 1 => new AconstNull(owner)
            case 2 => new PushI(owner)
            case 3 => new PushI(owner)
            case 4 => new PushI(owner)
            case 5 => new PushI(owner)
            case 6 => new PushI(owner)
            case 7 => new PushI(owner)
            case 8 => new PushI(owner)
            case 9 => new PushL(owner)
            case 10 => new PushL(owner)
            case 11 => new PushF(owner)
            case 12 => new PushF(owner)
            case 13 => new PushF(owner)
            case 14 => new PushD(owner)
            case 15 => new PushD(owner)
            case 16 => new PushI(owner)
            case 17 => new PushI(owner)
            case 18 => new Ldc(owner)
            case 19 => new LdcW(owner)
            case 20 => new Ldc2W(owner)
            case 21 => new LoadI(owner)
            case 22 => new LoadL(owner)
            case 23 => new LoadF(owner)
            case 24 => new LoadD(owner)
            case 25 => new LoadA(owner)
            case 26 => new LoadI(owner)
            case 27 => new LoadI(owner)
            case 28 => new LoadI(owner)
            case 29 => new LoadI(owner)
            case 30 => new LoadL(owner)
            case 31 => new LoadL(owner)
            case 32 => new LoadL(owner)
            case 33 => new LoadL(owner)
            case 34 => new LoadF(owner)
            case 35 => new LoadF(owner)
            case 36 => new LoadF(owner)
            case 37 => new LoadF(owner)
            case 38 => new LoadD(owner)
            case 39 => new LoadD(owner)
            case 40 => new LoadD(owner)
            case 41 => new LoadD(owner)
            case 42 => new LoadA(owner)
            case 43 => new LoadA(owner)
            case 44 => new LoadA(owner)
            case 45 => new LoadA(owner)
            case 46 => new LoadFromIArray(owner)
            case 47 => new LoadFromLArray(owner)
            case 48 => new LoadFromFArray(owner)
            case 49 => new LoadFromDArray(owner)
            case 50 => new LoadFromAArray(owner)
            case 51 => new LoadFromBArray(owner)
            case 52 => new LoadFromCArray(owner)
            case 53 => new LoadFromSArray(owner)
            case 54 => new StoreI(owner)
            case 55 => new StoreL(owner)
            case 56 => new StoreF(owner)
            case 57 => new StoreD(owner)
            case 58 => new StoreA(owner)
            case 59 => new StoreI(owner)
            case 60 => new StoreI(owner)
            case 61 => new StoreI(owner)
            case 62 => new StoreI(owner)
            case 63 => new StoreL(owner)
            case 64 => new StoreL(owner)
            case 65 => new StoreL(owner)
            case 66 => new StoreL(owner)
            case 67 => new StoreF(owner)
            case 68 => new StoreF(owner)
            case 69 => new StoreF(owner)
            case 70 => new StoreF(owner)
            case 71 => new StoreD(owner)
            case 72 => new StoreD(owner)
            case 73 => new StoreD(owner)
            case 74 => new StoreD(owner)
            case 75 => new StoreA(owner)
            case 76 => new StoreA(owner)
            case 77 => new StoreA(owner)
            case 78 => new StoreA(owner)
            case 79 => new StoreIntoIArray(owner)
            case 80 => new StoreIntoLArray(owner)
            case 81 => new StoreIntoFArray(owner)
            case 82 => new StoreIntoDArray(owner)
            case 83 => new StoreIntoAArray(owner)
            case 84 => new StoreIntoBArray(owner)
            case 85 => new StoreIntoCArray(owner)
            case 86 => new StoreIntoSArray(owner)
            case 87 => new Pop(owner)
            case 88 => new Pop2(owner)
            case 89 => new Dup(owner)
            case 90 => new DupX1(owner)
            case 91 => new DupX2(owner)
            case 92 => new Dup2(owner)
            case 93 => new Dup2X1(owner)
            case 94 => new Dup2X2(owner)
            case 95 => new Swap(owner)
            case 96 => new Iadd(owner)
            case 97 => new Ladd(owner)
            case 98 => new Fadd(owner)
            case 99 => new Dadd(owner)
            case 100 => new Isub(owner)
            case 101 => new Lsub(owner)
            case 102 => new Fsub(owner)
            case 103 => new Dsub(owner)
            case 104 => new Imul(owner)
            case 105 => new Lmul(owner)
            case 106 => new Fmul(owner)
            case 107 => new Dmul(owner)
            case 108 => new Idiv(owner)
            case 109 => new Ldiv(owner)
            case 110 => new Fdiv(owner)
            case 111 => new Ddiv(owner)
            case 112 => new Irem(owner)
            case 113 => new Lrem(owner)
            case 114 => new Frem(owner)
            case 115 => new Drem(owner)
            case 116 => new Ineg(owner)
            case 117 => new Lneg(owner)
            case 118 => new Fneg(owner)
            case 119 => new Dneg(owner)
            case 120 => new Ishl(owner)
            case 121 => new Lshl(owner)
            case 122 => new Ishr(owner)
            case 123 => new Lshr(owner)
            case 124 => new Iushr(owner)
            case 125 => new Lushr(owner)
            case 126 => new Iand(owner)
            case 127 => new Land(owner)
            case 128 => new Ior(owner)
            case 129 => new Lor(owner)
            case 130 => new Ixor(owner)
            case 131 => new Lxor(owner)
            case 132 => new Iinc(owner)
            case 133 => new I2l(owner)
            case 134 => new I2f(owner)
            case 135 => new I2d(owner)
            case 136 => new L2i(owner)
            case 137 => new L2f(owner)
            case 138 => new L2d(owner)
            case 139 => new F2i(owner)
            case 140 => new F2l(owner)
            case 141 => new F2d(owner)
            case 142 => new D2i(owner)
            case 143 => new D2l(owner)
            case 144 => new D2f(owner)
            case 145 => new I2b(owner)
            case 146 => new I2c(owner)
            case 147 => new I2s(owner)
            case 148 => new Lcmp(owner)
            case 149 => new Fcmpl(owner)
            case 150 => new Fcmpg(owner)
            case 151 => new Dcmpl(owner)
            case 152 => new Dcmpg(owner)
            case 153 => new Ifeq(owner)
            case 154 => new Ifne(owner)
            case 155 => new Iflt(owner)
            case 156 => new Ifge(owner)
            case 157 => new Ifgt(owner)
            case 158 => new Ifle(owner)
            case 159 => new IfIcmpeq(owner)
            case 160 => new IfIcmpne(owner)
            case 161 => new IfIcmplt(owner)
            case 162 => new IfIcmpge(owner)
            case 163 => new IfIcmpgt(owner)
            case 164 => new IfIcmple(owner)
            case 165 => new IfAcmpeq(owner)
            case 166 => new IfAcmpne(owner)
            case 167 => new Goto(owner)
            case 168 => new Jsr(owner)
            case 169 => new Ret(owner)
            case 170 => new Switch(owner)
            case 171 => new Switch(owner)
            case 172 => new Ireturn(owner)
            case 173 => new Lreturn(owner)
            case 174 => new Freturn(owner)
            case 175 => new Dreturn(owner)
            case 176 => new Areturn(owner)
            case 177 => new Return(owner)
            case 178 => new Getstatic(owner)
            case 179 => new Putstatic(owner)
            case 180 => new Getfield(owner)
            case 181 => new Putfield(owner)
            case 182 => new Invokevirtual(owner)
            case 183 => new Invokespecial(owner)
            case 184 => new Invokestatic(owner)
            case 185 => new Invokeinterface(owner)
            // XXX: maybe implement this, but I feel dumb after reading the
            // description ...
            case 186 => throw new Exception("invokedynamic not implemented")
            case 187 => new New(owner)
            case 188 => new Newarray(owner)
            case 189 => new Anewarray(owner)
            case 190 => new Arraylength(owner)
            case 191 => new Athrow(owner)
            case 192 => new Checkcast(owner)
            case 193 => new Instanceof(owner)
            case 194 => new Monitorenter(owner)
            case 195 => new Monitorexit(owner)
            case 196 => return parseWide(owner, input)
            case 197 => new Multianewarray(owner)
            case 198 => new Ifnull(owner)
            case 199 => new Ifnonnull(owner)
            case 200 => new GotoW(owner)
            case 201 => new JsrW(owner)
            case _ => throw new Exception("Unknown op code" + opCode)
        }

        operation.deserialize(startAddress, opCode, input)
        var result = operation.canonicalForm()
        result.pc = startAddress
        return result
    }
}

