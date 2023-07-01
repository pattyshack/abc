import scala.collection.JavaConversions._


class CheckJumpTargets extends CodeAnalysisPass {
    def analyze(rootScope: CodeScope) {
        new JumpTargetChecker(rootScope).apply()
    }
}

class JumpTargetChecker(root: CodeScope) extends CodeVisitor(root) {
    override def visitBlock(block: CodeBlock) {
        block._ops.lastElement() match {
            case g: Goto => {
                checkScope(block, g._targetBlock)
            }
            case i: IfBaseOp => {
                checkScope(block, i._ifBranch)
                checkScope(block, i._elseBranch)
            }
            case s: Switch => {
                checkScope(block, s._defaultBranch)
                for (c <- s._table.values()) {
                    checkScope(block, c)
                }
            }
            case _ => {}
        }
    }

    // check lexical scoping:
    // 1. can jump to any block in current / parent scopes
    // 2. can only jump to entry point for all other scopes
    def checkScope(current: CodeBlock, target: CodeBlock) {
        val currentScope = current._parentScope
        val targetScope = target._parentScope

        if (targetScope._contains(currentScope)) {
            return
        }

        if (target != targetScope.getEntryBlock()) {
            throw new Exception("cannot jump to arbitrary block")
        }
    }
}
