import java.util.Collection
import java.util.TreeMap
import java.util.Vector

import scala.collection.JavaConversions._


class MarkAndSweepConstants extends AnalysisPass {
    var _classInfo: ClassInfo = null
    var _constInfos: TreeMap[ConstInfo, ConstInfo] = null

    def apply(c: ClassInfo) {
        _init(c)
        _setAllToUnused()
        _markUsed()
        _sweepUnused()
    }

    def _init(c: ClassInfo) {
        _classInfo = c
        _constInfos = c._constants._constInfos
    }

    def _setAllToUnused() {
        for (info <- _constInfos.keySet()) {
            info._used = false
        }
    }

    def _sweepUnused() {
        var toRemove = new Vector[ConstInfo]()
        for (info <- _constInfos.keySet()) {
            if (!info._used) {
                toRemove.add(info)
            }
        }

        for (info <- toRemove) {
            _constInfos.remove(info)
            println("UNUSED: " + info.debugString())
        }
    }

    def _markUsed() {
        for (field <- _classInfo.fields()) {
            _markFieldConstants(field)
        }
        for (method <- _classInfo.methods()) {
            _markMethodConstants(method)
        }
        _markAttributeGroupConstants(_classInfo.attributes())

        _classInfo._thisClass.markUsed()
        if (_classInfo._superClass != null) {
            _classInfo._superClass.markUsed()
        }

        for (iface <- _classInfo._interfaces) {
            iface.markUsed()
        }
    }

    def _markFieldConstants(field: FieldInfo) {
        field._name.markUsed()
        field._fieldTypeString.markUsed()
        _markAttributeGroupConstants(field._attributes)
    }

    def _markMethodConstants(method: MethodInfo) {
        method._name.markUsed()
        method._methodTypeString.markUsed()
        _markAttributeGroupConstants(method._attributes)
    }

    def _markAttributeGroupConstants(attributes: AttributeGroup) {
        for (attr <- attributes.allAttributes()) {
            attr._name.markUsed()
            attr match {
                // specific attributes
                case b: BootstrapMethodsAttribute => {
                    for (m <- b._methods) {
                        m._methodHandle.markUsed()
                        for (arg <- m._arguments) {
                            arg.markUsed()
                        }
                    }
                }
                case c: CodeAttribute => {
                    (new CodeConstantsMarker(c.code)).apply()
                    _markAttributeGroupConstants(c.attributes)
                }
                case c: ConstantValueAttribute => { c._constant.markUsed() }
                case e: EnclosingMethodAttribute => {
                    e._enclosingClass.markUsed()
                    if (e._methodNameAndType != null) {
                        e._methodNameAndType.markUsed()
                    }
                }
                case e: ExceptionsAttribute => {
                    for (c <- e._exceptions) {
                        c.markUsed()
                    }
                }
                case i: InnerClassesAttribute => {
                    for (inner <- i._innerClasses) {
                        inner._innerClass.markUsed()
                        if (inner._outerClass != null) {
                            inner._outerClass.markUsed()
                        }
                        if (inner._innerName != null) {
                            inner._innerName.markUsed()
                        }
                    }
                }
                case _: LineNumberTableAttribute => {}
                case l: LocalVariableTableBaseAttribute => {
                    for (entry <- l.table) {
                        entry._fieldName.markUsed()
                        entry._fieldDescriptor.markUsed()
                    }
                }
                // generic attributes
                case _: RawBytesAttribute => {}
                case _: NoValueAttribute => {}
                case s: StringValueAttribute => { s._value.markUsed() }
            }
        }
    }

    class CodeConstantsMarker(root: CodeScope) extends CodeVisitor(root) {
        override def visitExceptionTarget(t: ExceptionTarget) {
            if (t.exception != null) {
                t.exception.markUsed()
            }
        }

        override def visitBlock(b: CodeBlock) {
            for (op <- b._ops) {
                op match {
                    case o: PushI => {
                        if (o._constInt != null) {
                            o._constInt.markUsed()
                        }
                    }
                    case o: PushL => {
                        if (o._constLong != null) {
                            o._constLong.markUsed()
                        }
                    }
                    case o: PushF => {
                        if (o._constFloat != null) {
                            o._constFloat.markUsed()
                        }
                    }
                    case o: PushD => {
                        if (o._constDouble != null) {
                            o._constDouble.markUsed()
                        }
                    }
                    case o: PushString => { o._constString.markUsed() }
                    case o: ClassOp => { o._constClass.markUsed() }
                    case o: FieldOp => { o._constFieldRef.markUsed() }
                    case o: MethodOp => { o._constMethodRef.markUsed() }
                    case _ => {}
                }
            }
        }
    }
}
