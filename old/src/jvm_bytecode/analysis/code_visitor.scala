import java.util.HashMap
import java.util.Vector

import scala.collection.JavaConversions._


class CodeVisitor(root: CodeScope) {
    var rootScope = root

    def visitBlock(b: CodeBlock) {}

    def visitExceptionTarget(target: ExceptionTarget) {
    }

    def visitScope(scope: CodeScope) {
        for (b <- scope._blocks) {
            visitBlock(b)
        }

        for (e <- scope._exceptionTargets) {
            visitExceptionTarget(e)
        }

        for (s <- scope._subsections) {
            visitScope(s)
        }
    }

    def apply() {
        visitScope(rootScope)
    }
}

class PcIdResetter(root: CodeScope, m: HashMap[Int, CodeScope])
        extends CodeVisitor(root) {
    var nextId = 1
    var mapping = m

    override def visitBlock(block: CodeBlock) {
        block.pc = -1
        block._endPc = -1
        block.segmentId = -1

        block._unorderedId = nextId
        nextId += 1
    }

    override def visitScope(scope: CodeScope) {
        super.visitScope(scope)

        scope.pc = -1
        scope._endPc = -1
        scope.segmentId = -1

        mapping.put(nextId, scope)
        scope._unorderedId = nextId
        nextId += 1
    }
}

class ScopePcUpdater(root: CodeScope) extends CodeVisitor(root) {
    override def visitScope(scope: CodeScope) {
        super.visitScope(scope)

        scope.pc = -1
        scope._endPc = -1
        for (seg <- scope._segments) {
            if (scope.pc == -1) {
                scope.pc = seg.pc
                scope._endPc = seg._endPc
            } else {
                if (seg.pc < scope.pc) {
                    scope.pc = seg.pc
                }

                if (seg._endPc > scope._endPc) {
                    scope._endPc = seg._endPc
                }
            }
        }
    }
}

// only for deserialization
class EntryPointSetter(root: CodeScope) extends CodeVisitor(root) {
    override def visitScope(scope: CodeScope) {
        super.visitScope(scope)
        scope._entryPoint = scope._segments.elementAt(0)
    }

    override def apply() {
        rootScope.sort()
        super.apply()
    }
}
