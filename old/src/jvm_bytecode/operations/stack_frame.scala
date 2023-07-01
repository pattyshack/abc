import java.util.Stack
import java.util.Vector

import scala.collection.JavaConversions._
import scala.reflect.ClassTag


class StackFrame {
    var maxStack = 0
    var maxLocals = 0

    var stack = new Stack[FieldType]()
    var locals = new Vector[FieldType]()

    def mergeFrom(other: StackFrame, mergeStack: Boolean): Boolean = {
        if (other.maxStack > maxStack) {
            maxStack = other.maxStack
        }
        if (other.maxLocals > maxLocals) {
            maxLocals = other.maxLocals
        }

        var modified = false
        if (mergeStack) {
            modified = _mergeStackFrom(other)
        }

        modified |= _mergeLocalsFrom(other)

        return modified
    }

    def _mergeStackFrom(other: StackFrame): Boolean = {
        if (stack.size() != other.stack.size()) {
            throw new Exception("different stack height ...")
        }

        var modified = false
        var updatedStack = new Stack[FieldType]()
        for (i <- 0.to(stack.size() - 1)) {
            val origType = stack.elementAt(i)
            val newType = _mergeTypes(origType, other.stack.elementAt(i), false)
            if (origType.compareTo(newType) != 0) {
                modified = true
            }
            updatedStack.add(newType)
        }

        stack = updatedStack
        return modified
    }

    def _mergeLocalsFrom(other: StackFrame): Boolean = {
        var modified = locals.size() != other.locals.size()
        var updatedLocals = new Vector[FieldType]()

        var minSize = locals.size()
        if (other.locals.size() < minSize) {
            minSize = other.locals.size()
        }

        for (i <- 0.to(minSize - 1)) {
            val origType = stack.elementAt(i)
            val newType = _mergeTypes(origType, other.stack.elementAt(i), true)
            if (origType.compareTo(newType) != 0) {
                modified = true
            }
            updatedLocals.add(newType)
        }

        locals = updatedLocals
        return modified
    }

    def cloneFrame(): StackFrame = {
        var frame = cloneLocals()
        for (field <- stack) {
            frame.stack.add(field)
        }
        return frame
    }

    def cloneLocals(): StackFrame = {
        var frame = new StackFrame()
        frame.maxStack = maxStack
        frame.maxLocals = maxLocals
        for (field <- locals) {
            frame.locals.add(field)
        }
        return frame
    }

    def push(f: FieldType) {
        f match {
            case _: UnusableType => {
                throw new Exception("cannot push unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly push top type")
            }
            case _: LongType => {
                stack.push(f)
                stack.push(new TopType())
            }
            case _: DoubleType => {
                stack.push(f)
                stack.push(new TopType())
            }
            case _ => {
                stack.push(f)
            }
        }

        if (stack.size() > maxStack) {
            maxStack = stack.size()
        }
    }

    def pop(expected: FieldType): FieldType = {
        expected match {
            case _: UnusableType => {
                throw new Exception("cannot push unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly push top type")
            }
            case _: IntType => return _pop[IntType]()
            case _: FloatType => return _pop[FloatType]()
            case _: LongType => {
                _pop[TopType]()
                return _pop[LongType]()
            }
            case _: DoubleType => {
                _pop[TopType]()
                return _pop[DoubleType]()
            }
            // TODO: check matching array type
            case _: ArrayType => return _pop[ArrayType]()
            // TODO: check matching (or subclass of) ref type
            case _: RefType => return _pop[RefType]()
            case _ => throw new Exception(
                    "Pop-ing unexpected type: " + expected.descriptorString())
        }
    }

    def _pop[T <: FieldType: ClassTag](): FieldType = {
        val f = stack.pop()

        val cls = implicitly[ClassTag[T]].runtimeClass
        f match {
            case t: T if cls.isInstance(t) => return t
            case _ => throw new Exception(
                    "Unexpected type: " + f.descriptorString())
        }
    }

    def store(expected: FieldType, i: Int) {
        val f = pop(expected)

        val needed = i + f.categorySize()
        while (needed > locals.size()) {
            locals.add(new UnusableType())
        }

        locals.elementAt(i) match {
            case _: TopType => {
                // invalidate previous long / double
                locals.setElementAt(new UnusableType(), i - 1)
            }
            case _: DoubleType => {
                // invalidate top
                locals.setElementAt(new UnusableType(), i + 1)
            }
            case _: LongType => {
                // invalidate top
                locals.setElementAt(new UnusableType(), i + 1)
            }
        }

        f match {
            case _: UnusableType => {
                throw new Exception("cannot store unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot directly store top type")
            }
            case _: LongType => {
                locals.setElementAt(f, i)
                locals.setElementAt(new TopType(), i + 1)
            }
            case _: DoubleType => {
                locals.setElementAt(f, i)
                locals.setElementAt(new TopType(), i + 1)
            }
            case _ => {
                locals.setElementAt(f, i)
            }
        }

        if (locals.size() > maxLocals) {
            maxLocals = locals.size()
        }
    }

    def load(expected: FieldType, i: Int) {
        val f = locals.elementAt(i)
        f match {
            case _: UnusableType => {
                throw new Exception("cannot load unusable type")
            }
            case _: TopType => {
                throw new Exception("cannot load top type")
            }
            case _ => {}  // continue
        }

        expected match {
            case _: IntType => _check[IntType](f)
            case _: FloatType => _check[FloatType](f)
            case _: LongType => _check[LongType](f)
            case _: DoubleType => _check[DoubleType](f)
            // TODO: check matching array type
            case _: ArrayType => _check[ArrayType](f)
            // TODO: check matching (or subclass of) ref type
            case _: RefType => _check[RefType](f)
            case _ => throw new Exception(
                    "unexpected load type: " + expected.descriptorString())
        }

        push(f)
    }

    def _check[T <: FieldType: ClassTag](f: FieldType) {
        val cls = implicitly[ClassTag[T]].runtimeClass
        f match {
            case t: T if cls.isInstance(t) => {}
            case _ => throw new Exception(
                    "Unexpected type: " + f.descriptorString())
        }
    }


    def _mergeTypes(
            v1: FieldType,
            v2: FieldType,
            mergeIncompatibleTypes: Boolean): FieldType = {
        if (v1.compareTo(v2) == 0) {
            return v1
        }

        var result: FieldType = v1 match {
            case a1: ArrayType => {
                v2 match {
                    case a2: ArrayType => _mergeArrayTypes(a1, a2)
                    case _: RefType => return new ObjectType(Const.JAVA_OBJECT)
                    case _ => null
                }
            }
            case o1: ObjectType => {
                v2 match {
                    case o2: ObjectType => _mergeObjectTypes(o1, o2)
                    case _: RefType => return new ObjectType(Const.JAVA_OBJECT)
                    case _ => null
                }
            }
            case _ => null
        }

        if (result != null) {
            return result
        }

        if (mergeIncompatibleTypes) {
            return new UnusableType()  // not "top"
        }

        throw new Exception(
                "merging incompatible types: " + v1.descriptorString() +
                        "vs. " + v2.descriptorString())
    }

    // This assumes v1.compareTo(v2) != 0
    def _mergeObjectTypes(v1: ObjectType, v2: ObjectType): ObjectType = {
        // TODO: load class catalog and find most specific common supertype.
        return new ObjectType(Const.JAVA_OBJECT)
    }

    // This assumes v1.compareTo(v2) != 0
    def _mergeArrayTypes(v1: ArrayType, v2: ArrayType): RefType = {
        // TODO: fix array merge semantic (see page 350 / section 4.10.2.2)
        val commonItemType: FieldType = v1.itemType match {
            case _: BaseType => new ObjectType(Const.JAVA_OBJECT)
            case a1: ArrayType => {
                v2.itemType match {
                    case a2: ArrayType => _mergeArrayTypes(a1, a2)
                    case _ => new ObjectType(Const.JAVA_OBJECT)
                }
            }
            case o1: ObjectType => {
                v2.itemType match {
                    case o2: ObjectType => _mergeObjectTypes(o1, o2)
                    case _ => new ObjectType(Const.JAVA_OBJECT)
                }
            }
        }

        return new ArrayType(commonItemType)
    }

    def apply(op: Operation) {
        if (_applyComparisonOp(op)) {
            return
        }
        if (_applyConstantOp(op)) {
            return
        }
        if (_applyControlOp(op)) {
            return
        }
        if (_applyConversionOp(op)) {
            return
        }
        if (_applyLoadOp(op)) {
            return
        }
        if (_applyMathOp(op)) {
            return
        }
        if (_applyReferenceOp(op)) {
            return
        }
        if (_applyStackOp(op)) {
            return
        }
        if (_applyStoreOp(op)) {
            return
        }
        throw new Exception(
                "op not implemented in stack frame: " + op.debugString(""))
    }

    def _applyConstantOp(op: Operation): Boolean = {
        op match {
            case _: PushI => push(new IntType())
            case _: PushL => push(new LongType())
            case _: PushF => push(new FloatType())
            case _: PushD => push(new DoubleType())
            case _: PushString => push(new ObjectType(Const.JAVA_STRING))
            case _: Ldc => {
                // TODO: support class / method type / method handle ...
                throw new Exception("raw ldc unsupported")
            }
            case _: LdcW => {
                // TODO: support class / method type / method handle ...
                throw new Exception("raw ldc_w unsupported")
            }
            case _ => return false
        }

        return true
    }

    def _applyLoadOp(op: Operation): Boolean = {
        op match {
            case o: LoadI => load(new IntType(), o.index)
            case o: LoadL => load(new LongType(), o.index)
            case o: LoadF => load(new FloatType(), o.index)
            case o: LoadD => load(new DoubleType(), o.index)
            case o: LoadA => load(new CheckRefType(), o.index)
            case o: LoadFromBaseIntArray => {
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[BaseIntType](arrType.itemType)
                        push(new IntType())
                    }
                }
            }
            case o: LoadFromLArray => {
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[LongType](arrType.itemType)
                        push(new LongType())
                    }
                }
            }
            case o: LoadFromFArray => {
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[FloatType](arrType.itemType)
                        push(new FloatType())
                    }
                }
            }
            case o: LoadFromDArray => {
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[DoubleType](arrType.itemType)
                        push(new DoubleType())
                    }
                }
            }
            case o: LoadFromAArray => {
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[RefType](arrType.itemType)
                        push(arrType.itemType)
                    }
                }
            }
            case _ => return false
        }

        return true
    }

    def _applyStoreOp(op: Operation): Boolean = {
        op match {
            case o: StoreI => store(new IntType(), o.index)
            case o: StoreL => store(new LongType(), o.index)
            case o: StoreF => store(new FloatType(), o.index)
            case o: StoreD => store(new DoubleType(), o.index)
            case o: StoreA => store(new CheckRefType(), o.index)
            case o: StoreIntoBaseIntArray => {
                pop(new IntType())
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => _check[BaseIntType](
                            arrType.itemType)
                }
            }
            case o: StoreIntoLArray => {
                pop(new LongType())
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => _check[LongType](
                            arrType.itemType)
                }
            }
            case o: StoreIntoFArray => {
                pop(new FloatType())
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => _check[FloatType](
                            arrType.itemType)
                }
            }
            case o: StoreIntoDArray => {
                pop(new DoubleType())
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => _check[DoubleType](
                            arrType.itemType)
                }
            }
            case o: StoreIntoAArray => {
                val item = pop(new CheckRefType())
                pop(new IntType())
                pop(new CheckArrayType()) match {
                    case arrType: ArrayType => {
                        _check[RefType](arrType.itemType)
                        // TODO check arrType.itemType is supertype of item
                    }
                }
            }
            case _ => return false
        }

        return true
    }

    def _applyConversionOp(op: Operation): Boolean = {
        op match {
            case _: I2l => {
                pop(new IntType())
                push(new LongType())
            }
            case _: I2f => {
                pop(new IntType())
                push(new FloatType())
            }
            case _: I2d => {
                pop(new IntType())
                push(new DoubleType())
            }
            case _: TruncateI => {
                pop(new IntType())
                push(new IntType())
            }
            case _: L2i => {
                pop(new LongType())
                push(new IntType())
            }
            case _: L2f => {
                pop(new LongType())
                push(new FloatType())
            }
            case _: L2d => {
                pop(new LongType())
                push(new DoubleType())
            }
            case _: F2i => {
                pop(new FloatType())
                push(new IntType())
            }
            case _: F2l => {
                pop(new FloatType())
                push(new LongType())
            }
            case _: F2d => {
                pop(new FloatType())
                push(new DoubleType())
            }
            case _: D2i => {
                pop(new DoubleType())
                push(new IntType())
            }
            case _: D2l => {
                pop(new DoubleType())
                push(new LongType())
            }
            case _: D2f => {
                pop(new DoubleType())
                push(new FloatType())
            }
            case _ => return false
        }

        return true
    }

    def _applyStackOp(op: Operation): Boolean = {
        op match {
            case _: Pop => {
                _assertCat1(stack.pop())
            }
            case _: Pop2 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                _assertTwoCat1sOrOneCat2(value2, value1)
            }
            case _: Dup => {
                val value = stack.pop()
                _assertCat1(value)
                push(value)
                push(value)
            }
            case _: DupX1 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                _assertCat1(value1)
                _assertCat1(value2)
                push(value1)
                push(value2)
                push(value1)
            }
            case _: DupX2 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                val value3 = stack.pop()
                _assertCat1(value1)
                _assertTwoCat1sOrOneCat2(value3, value2)
                push(value1)
                push(value3)
                push(value2)
                push(value1)
            }
            case _: Dup2 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                _assertTwoCat1sOrOneCat2(value2, value1)
                push(value2)
                push(value1)
                push(value2)
                push(value1)
            }
            case _: Dup2X1 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                val value3 = stack.pop()
                _assertTwoCat1sOrOneCat2(value2, value1)
                _assertCat1(value3)
                push(value2)
                push(value1)
                push(value3)
                push(value2)
                push(value1)
            }
            case _: Dup2X2 => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                val value3 = stack.pop()
                val value4 = stack.pop()
                _assertTwoCat1sOrOneCat2(value2, value1)
                _assertTwoCat1sOrOneCat2(value4, value3)
                push(value2)
                push(value1)
                push(value4)
                push(value3)
                push(value2)
                push(value1)
            }
            case _: Swap => {
                val value1 = stack.pop()
                val value2 = stack.pop()
                _assertCat1(value1)
                _assertCat1(value2)
                push(value1)
                push(value2)
            }
            case _ => return false
        }

        return true
    }

    def _assertCat1(f: FieldType) {
        if (!f.isCat1()) {
            throw new Exception(
                    "given type is not category 1: " + f.descriptorString())
        }
    }

    // value2 before value1 to match semantics in jvm spec
    def _assertTwoCat1sOrOneCat2(value2: FieldType, value1: FieldType) {
        if (value1.isCat1() && value2.isCat1()) {
            return
        }
        if (value2.isCat2()) {
            // sanity check
            value1 match {
                case _: TopType => return
                case _ => throw new Exception("Unexpected")
            }
        }
        throw new Exception(
                "Invalid. pop-ing: " + value2.descriptorString() +
                        " & " + value1.descriptorString())
    }

    def _applyMathOp(op: Operation): Boolean = {
        op match {
            case _: BinaryIOp => {
                pop(new IntType())
                pop(new IntType())
                push(new IntType())
            }
            case _: UnaryIOp => {
                pop(new IntType())
                push(new IntType())
            }
            case o: Iinc => {
                load(new IntType(), o.operand1)
                pop(new IntType())
            }
            case _: BinaryLOp => {
                pop(new LongType())
                pop(new LongType())
                push(new LongType())
            }
            case _: UnaryLOp => {
                pop(new LongType())
                push(new LongType())
            }
            case _: ShiftLOp => {
                pop(new IntType())
                pop(new LongType())
                push(new LongType())
            }
            case _: BinaryFOp => {
                pop(new FloatType())
                pop(new FloatType())
                push(new FloatType())
            }
            case _: UnaryFOp => {
                pop(new FloatType())
                push(new FloatType())
            }
            case _: BinaryDOp => {
                pop(new DoubleType())
                pop(new DoubleType())
                push(new DoubleType())
            }
            case _: UnaryDOp => {
                pop(new DoubleType())
                push(new DoubleType())
            }
            case _ => return false
        }

        return true
    }

    def _applyComparisonOp(op: Operation): Boolean = {
        op match {
            case _: Lcmp => {
                pop(new LongType())
                pop(new LongType())
                push(new IntType())
            }
            case _: Fcmp => {
                pop(new FloatType())
                pop(new FloatType())
                push(new IntType())
            }
            case _: Dcmp => {
                pop(new DoubleType())
                pop(new DoubleType())
                push(new IntType())
            }
            case _: IfIOp => pop(new IntType())
            case _: IfCmpIOp => {
                pop(new IntType())
                pop(new IntType())
            }
            case _: IfAOp => pop(new CheckRefType())
            case _: IfCmpAOp => {
                pop(new CheckRefType())
                pop(new CheckRefType())
            }
            case _ => return false
        }

        return true
    }

    def _applyControlOp(op: Operation): Boolean = {
        op match {
            case _: Goto => {}
            case _: GotoW => {}
            case _: Jsr => throw new Exception("jsr deprecated")
            case _: JsrW => throw new Exception("jsr_w deprecated")
            case _: Ret => throw new Exception("ret deprecated")
            case _: Return => {}
            case _: Ireturn => pop(new IntType())
            case _: Lreturn => pop(new LongType())
            case _: Freturn => pop(new FloatType())
            case _: Dreturn => pop(new DoubleType())
            case _: Areturn => pop(new CheckRefType())
            case _: Switch => pop(new IntType())
            case _ => return false
        }

        return true
    }

    def _applyReferenceOp(op: Operation): Boolean = {
        op match {
            case o: Getstatic => push(o._constFieldRef.fieldDescriptor())
            case o: Putstatic => pop(o._constFieldRef.fieldDescriptor())
            case o: Getfield => {
                pop(new CheckRefType())
                push(o._constFieldRef.fieldDescriptor())
            }
            case o: Putfield => {
                pop(o._constFieldRef.fieldDescriptor())
                pop(new CheckRefType())
            }
            case o: Invokeinterface => _applyMethod(
                    o._constMethodRef.methodDescriptor(),
                    true)  // on obj ref
            case o: Invokespecial => _applyMethod(
                    o._constMethodRef.methodDescriptor(),
                    true)  // on obj ref
            case o: Invokestatic => _applyMethod(
                    o._constMethodRef.methodDescriptor(),
                    false)  // on obj ref
            case o: Invokevirtual => _applyMethod(
                    o._constMethodRef.methodDescriptor(),
                    true)  // on obj ref
            case o: New => push(o._classType)
            case o: Checkcast => {
                pop(new CheckRefType())
                push(o._classType)
            }
            case o: Instanceof => {
                pop(new CheckRefType())
                push(new IntType())
            }
            case o: Newarray => {
                pop(new IntType())
                push(new ArrayType(o._arrayType))
            }
            case o: Anewarray => {
                pop(new IntType())
                push(new ArrayType(o._classType))
            }
            case o: Multianewarray => {
                for (_ <- 1 to o._dimensions) {
                    pop(new IntType())
                }
                push(o._classType)
            }
            case _: Arraylength => {
                pop(new CheckArrayType())
                push(new IntType())
            }
            case _: Athrow => {
                // TODO check is subtype of throwable
                pop(new CheckRefType())
            }
            case _: Monitorenter => {
                pop(new CheckRefType())
            }
            case _: Monitorexit => {
                pop(new CheckRefType())
            }
            case _ => return false
        }

        return true
    }

    def _applyMethod(m: MethodType, onObjRef: Boolean) {
        val params = m.parameters._parameters
        for (i <- 1 to params.size()) {
            pop(params.elementAt(params.size() - i))
        }

        if (onObjRef) {
            pop(new CheckRefType())
        }

        if (m.returnType != null) {
            push(m.returnType)
        }
    }
}
