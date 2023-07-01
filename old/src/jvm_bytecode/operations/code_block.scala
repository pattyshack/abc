import java.io.DataInputStream
import java.io.DataOutput
import java.util.Collections
import java.util.HashMap
import java.util.HashSet
import java.util.Stack
import java.util.TreeMap
import java.util.TreeSet
import java.util.Vector

import scala.collection.JavaConversions._


abstract class CodeSegment(
        owner: AttributeOwner,
        parent: CodeScope)
        extends Operation(owner) with Comparable[CodeSegment] {
    var _parentScope: CodeScope = parent

    // not inclusive.
    var _endPc = -1

    var _unorderedId = -1  // used for map lookup when segmentId is unavailable

    var segmentId = -1

    var _reachable = false

    // the stack frame state before executing operations in the current segment
    var startState: StackFrame = null

    // the max stack, max locals, and the local variable types that remain
    // "stable" throughout the execution of the current segment.  Unstable
    // variables are either chopped off from the tail, or are set to
    // UnusableType.  Stack should remain empty.
    var stableState: StackFrame = null

    var implicitGoto: CodeSegment = null

    def getImplicitGoto(): CodeSegment = {
        if (implicitGoto != null) {
            return implicitGoto
        }

        if (_parentScope == null) {
            return null
        }

        return _parentScope.getImplicitGoto()
    }

    def getRootScope(): CodeScope = {
        if (_parentScope != null) {
            return _parentScope.getRootScope()
        }

        this match {
            case s: CodeScope => return s
        }
    }

    def getEntryBlock(): CodeBlock

    def compareTo(other: CodeSegment): Int = {
        if (segmentId < other.segmentId) {
            return -1
        } else if (segmentId > other.segmentId) {
            return 1
        }

        if (pc < other.pc) {
            return -1
        } else if (pc > other.pc) {
            return 1
        }

        return 0
    }
}

// a linear section in the control flow graph.
//
// Only the last op maybe a jump or a return.  No other ops in the block
// can be a jump (or a return).
//
// if implicitGoto is set and the last ops is not a jump/return, then
// and goto is inserted to the end of the block during verification.
class CodeBlock(owner: AttributeOwner, parent: CodeScope)
        extends CodeSegment(owner, parent) {
    var lineContext = -1

    var _ops = new Vector[Operation]()

    var _hasControlOp = false

    def _add(op: Operation) {
        if (_hasControlOp) {
            throw new Exception(
                    "Cannot add more ops after control op: " +
                            _ops.lastElement().pc)
        }

        if (lineContext >= 0) {
            op.line = lineContext
        }
        _ops.add(op)

        _hasControlOp = op match {
            case _: Return => true
            case _: ReturnValue => true
            case _: Athrow => true
            case g: Goto => true
            case i: IfBaseOp => true
            case s: Switch => true
            case _ => false
        }
    }

    //
    // Non-control operations
    //

    def pushNull() { _add(new AconstNull(_owner)) }
    def pushI(v: Int) { _add(new PushI(_owner, v)) }
    def pushL(v: Long) { _add(new PushL(_owner, v)) }
    def pushF(v: Float) { _add(new PushF(_owner, v)) }
    def pushD(v: Double) { _add(new PushD(_owner, v)) }
    def pushString(v: String) { _add(new PushString(_owner, v)) }
    // TODO: support other ldc constant types

    def load(name: String) {
        val entry = _parentScope.getLocal(name)
        val op = entry.fieldType match {
            case _: BaseIntType => new LoadI(_owner, entry.index)
            case _: FloatType => new LoadF(_owner, entry.index)
            case _: LongType => new LoadL(_owner, entry.index)
            case _: DoubleType => new LoadD(_owner, entry.index)
            case _: RefType => new LoadA(_owner, entry.index)
        }
        _add(op)
    }

    def loadI(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new LoadI(_owner, index))
    }
    def loadL(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new LoadL(_owner, index))
    }
    def loadF(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new LoadF(_owner, index))
    }
    def loadD(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new LoadD(_owner, index))
    }
    def loadA(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new LoadA(_owner, index))
    }

    def store(name: String) {
        val entry = _parentScope.getLocal(name)
        val op = entry.fieldType match {
            case _: BaseIntType => new StoreI(_owner, entry.index)
            case _: FloatType => new StoreF(_owner, entry.index)
            case _: LongType => new StoreL(_owner, entry.index)
            case _: DoubleType => new StoreD(_owner, entry.index)
            case _: RefType => new StoreA(_owner, entry.index)
        }
        _add(op)
    }

    def storeI(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new StoreI(_owner, index))
    }
    def storeL(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new StoreL(_owner, index))
    }
    def storeF(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new StoreF(_owner, index))
    }
    def storeD(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new StoreD(_owner, index))
    }
    def storeA(index: Int) {
        getRootScope()._disableNamedLocals()
        _add(new StoreA(_owner, index))
    }

    def loadFromIArray() { _add(new LoadFromIArray(_owner)) }
    def loadFromLArray() { _add(new LoadFromLArray(_owner)) }
    def loadFromFArray() { _add(new LoadFromFArray(_owner)) }
    def loadFromDArray() { _add(new LoadFromDArray(_owner)) }
    def loadFromAArray() { _add(new LoadFromAArray(_owner)) }
    def loadFromBArray() { _add(new LoadFromBArray(_owner)) }
    def loadFromCArray() { _add(new LoadFromCArray(_owner)) }
    def loadFromSArray() { _add(new LoadFromSArray(_owner)) }

    def storeIntoIArray() { _add(new StoreIntoIArray(_owner)) }
    def storeIntoLArray() { _add(new StoreIntoLArray(_owner)) }
    def storeIntoFArray() { _add(new StoreIntoFArray(_owner)) }
    def storeIntoDArray() { _add(new StoreIntoDArray(_owner)) }
    def storeIntoAArray() { _add(new StoreIntoAArray(_owner)) }
    def storeIntoBArray() { _add(new StoreIntoBArray(_owner)) }
    def storeIntoCArray() { _add(new StoreIntoCArray(_owner)) }
    def storeIntoSArray() { _add(new StoreIntoSArray(_owner)) }

    def pop() { _add(new Pop(_owner)) }
    def pop2() { _add(new Pop2(_owner)) }
    def dup() { _add(new Dup(_owner)) }
    def dupX1() { _add(new DupX1(_owner)) }
    def dupX2() { _add(new DupX2(_owner)) }
    def dup2() { _add(new Dup2(_owner)) }
    def dup2X1() { _add(new Dup2X1(_owner)) }
    def dup2X2() { _add(new Dup2X2(_owner)) }
    def swap() { _add(new Swap(_owner)) }

    def addI() { _add(new Iadd(_owner)) }
    def subI() { _add(new Isub(_owner)) }
    def mulI() { _add(new Imul(_owner)) }
    def divI() { _add(new Idiv(_owner)) }
    def remI() { _add(new Irem(_owner)) }
    def negI() { _add(new Ineg(_owner)) }
    def shlI() { _add(new Ishl(_owner)) }
    def shrI() { _add(new Ishr(_owner)) }
    def ushrI() { _add(new Iushr(_owner)) }
    def andI() { _add(new Iand(_owner)) }
    def orI() { _add(new Ior(_owner)) }
    def xorI() { _add(new Ixor(_owner)) }
    def incI(index: Int, v: Int) { _add(new Iinc(_owner, index, v)) }

    def addL() { _add(new Ladd(_owner)) }
    def subL() { _add(new Lsub(_owner)) }
    def mulL() { _add(new Lmul(_owner)) }
    def divL() { _add(new Ldiv(_owner)) }
    def remL() { _add(new Lrem(_owner)) }
    def negL() { _add(new Lneg(_owner)) }
    def shlL() { _add(new Lshl(_owner)) }
    def shrL() { _add(new Lshr(_owner)) }
    def ushrL() { _add(new Lushr(_owner)) }
    def andL() { _add(new Land(_owner)) }
    def orL() { _add(new Lor(_owner)) }
    def xorL() { _add(new Lxor(_owner)) }

    def addF() { _add(new Fadd(_owner)) }
    def subF() { _add(new Fsub(_owner)) }
    def mulF() { _add(new Fmul(_owner)) }
    def divF() { _add(new Fdiv(_owner)) }
    def remF() { _add(new Frem(_owner)) }
    def negF() { _add(new Fneg(_owner)) }

    def addD() { _add(new Dadd(_owner)) }
    def subD() { _add(new Dsub(_owner)) }
    def mulD() { _add(new Dmul(_owner)) }
    def divD() { _add(new Ddiv(_owner)) }
    def remD() { _add(new Drem(_owner)) }
    def negD() { _add(new Dneg(_owner)) }

    def i2l() { _add(new I2l(_owner)) }
    def i2f() { _add(new I2f(_owner)) }
    def i2d() { _add(new I2d(_owner)) }
    def l2i() { _add(new L2i(_owner)) }
    def l2f() { _add(new L2f(_owner)) }
    def l2d() { _add(new L2d(_owner)) }
    def f2i() { _add(new F2i(_owner)) }
    def f2l() { _add(new F2l(_owner)) }
    def f2d() { _add(new F2d(_owner)) }
    def d2i() { _add(new D2i(_owner)) }
    def d2l() { _add(new D2l(_owner)) }
    def d2f() { _add(new D2f(_owner)) }
    def i2b() { _add(new I2b(_owner)) }
    def i2c() { _add(new I2c(_owner)) }
    def i2s() { _add(new I2s(_owner)) }

    def cmpL() { _add(new Lcmp(_owner)) }
    def cmplF() { _add(new Fcmpl(_owner)) }
    def cmpgF() { _add(new Fcmpg(_owner)) }
    def cmplD() { _add(new Dcmpl(_owner)) }
    def cmpgD() { _add(new Dcmpg(_owner)) }

    def getStatic(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Getstatic(_owner, className, fieldName, fieldType))
    }
    def putStatic(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Putstatic(_owner, className, fieldName, fieldType))
    }
    def getField(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Getfield(_owner, className, fieldName, fieldType))
    }
    def putField(
            className: String,
            fieldName: String,
            fieldType: FieldType) {
        _add(new Putfield(_owner, className, fieldName, fieldType))
    }

    def invokeVirtual(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokevirtual(_owner, className, methodName, methodType))
    }
    def invokeSpecial(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokespecial(_owner, className, methodName, methodType))
    }
    def invokeStatic(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokestatic(_owner, className, methodName, methodType))
    }
    def invokeInterface(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _add(new Invokeinterface(_owner, className, methodName, methodType))
    }

    def newA(t: ObjectType) { _add(new New(_owner, t)) }
    def newArray(t: BaseType) { _add(new Newarray(_owner, t)) }
    def aNewArray(t: RefType) { _add(new Anewarray(_owner, t)) }
    def multiANewArray(t: ArrayType, d: Int) {
        _add(new Multianewarray(_owner, t, d))
    }

    def arrayLength() { _add(new Arraylength(_owner)) }

    def checkCast(t: ObjectType) { _add(new Checkcast(_owner, t)) }
    def instanceOf(t: ObjectType) { _add(new Instanceof(_owner, t)) }

    def enterMonitor() { _add(new Monitorenter(_owner)) }
    def exitMonitor() { _add(new Monitorexit(_owner)) }

    //
    // Control operations
    //

    def ifEq(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifeq(_owner, ifScope, elseScope))
    }
    def ifNe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifne(_owner, ifScope, elseScope))
    }
    def ifLt(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Iflt(_owner, ifScope, elseScope))
    }
    def ifGe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifge(_owner, ifScope, elseScope))
    }
    def ifGt(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifgt(_owner, ifScope, elseScope))
    }
    def ifLe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifle(_owner, ifScope, elseScope))
    }
    def ifICmpEq(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmpeq(_owner, ifScope, elseScope))
    }
    def ifICmpNe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmpne(_owner, ifScope, elseScope))
    }
    def ifICmpLt(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmplt(_owner, ifScope, elseScope))
    }
    def ifICmpGe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmpge(_owner, ifScope, elseScope))
    }
    def ifICmpGt(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmpgt(_owner, ifScope, elseScope))
    }
    def ifICmpLe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfIcmple(_owner, ifScope, elseScope))
    }
    def ifACmpEq(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfAcmpeq(_owner, ifScope, elseScope))
    }
    def ifACmpNe(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new IfAcmpne(_owner, ifScope, elseScope))
    }
    def ifNull(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifnull(_owner, ifScope, elseScope))
    }
    def ifNonNull(ifScope: CodeSegment, elseScope: CodeSegment) {
        _add(new Ifnonnull(_owner, ifScope, elseScope))
    }

    def goto(target: CodeSegment) { _add(new Goto(_owner, this, target)) }

    def switch(defaultBranch: CodeSegment): Switch = {
        val switch = new Switch(_owner, defaultBranch)
        _add(switch)
        return switch
    }

    // XXX: maybe infer return type from method signature?
    def returnI() { _add(new Ireturn(_owner)) }
    def returnL() { _add(new Lreturn(_owner)) }
    def returnF() { _add(new Freturn(_owner)) }
    def returnD() { _add(new Dreturn(_owner)) }
    def returnA() { _add(new Areturn(_owner)) }
    def returnVoid() { _add(new Return(_owner)) }

    def throwA() { _add(new Athrow(_owner)) }

    def getEntryBlock(): CodeBlock = this

    def serialize(output: DataOutput) {
        for (op <- _ops) {
            op.serialize(output)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("Cannot directly deserialize code block")
    }

    def debugString(indent: String): String = {
        var result = indent + "Block (pc: [" + pc + ", " + _endPc +
                ") segment: " + segmentId + " reachable: " + _reachable + ")\n"
        for (op <- _ops) {
            result += op.debugString(indent + "  ")
        }
        return result
    }
}

