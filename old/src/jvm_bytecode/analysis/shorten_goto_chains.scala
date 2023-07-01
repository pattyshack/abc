import scala.collection.JavaConversions._


class ShortenSimpleGotoChains extends CodeAnalysisPass {
    def analyze(root: CodeScope) {
        new SimpleGotoChainShortener(root).apply()
    }
}

object SimpleGotoChain {
    def shorten(block: CodeBlock): CodeBlock = {
        var target = block
        while (true) {
            var newTarget = _followSimpleGoto(target)
            if (newTarget == null) {
                return target
            } else {
                target = newTarget
            }
        }
        return null
    }

    def _followSimpleGoto(block: CodeBlock): CodeBlock = {
        if (block._ops.size() != 1) {
            return null
        }

        block._ops.lastElement() match {
            case g: Goto => return g._targetBlock
            case _ => return null
        }
    }
}

class SimpleGotoChainShortener(root: CodeScope) extends CodeVisitor(root) {
    override def visitBlock(block: CodeBlock) {
        block._ops.lastElement() match {
            case g: Goto => {
                g._targetBlock = SimpleGotoChain.shorten(g._targetBlock)
            }
            case i: IfBaseOp => {
                i._ifBranch = SimpleGotoChain.shorten(i._ifBranch)
            }
            case s: Switch => {
                s._defaultBranch = SimpleGotoChain.shorten(s._defaultBranch)
                for (c <- s._table.entrySet()) {
                    c.setValue(SimpleGotoChain.shorten(c.getValue()))
                }
            }
            case _ => {}
        }
    }

}

class AdjustEntryPoints extends CodeAnalysisPass {
    def analyze(root: CodeScope) {
        new EntryPointAdjuster(root).apply()
    }
}

class EntryPointAdjuster(root: CodeScope) extends CodeVisitor(root) {
    override def visitScope(scope: CodeScope) {
        super.visitScope(scope)

        var target = SimpleGotoChain.shorten(scope.getEntryBlock())
        if (!scope._contains(target._parentScope)) {
            // This scope is useless.  do nothing.
            return
        }

        val subScope = findSubScope(scope, target._parentScope)
        if (subScope == null) {  // i.e. block in same scope
            scope._entryPoint = target
        } else {
            scope._entryPoint = subScope
        }
    }

    def findSubScope(current: CodeScope, other: CodeScope): CodeScope = {
        if (current == other) {
            return null
        }

        var tmp = other
        while (tmp._parentScope != current) {
            tmp = tmp._parentScope
        }

        return tmp
    }
}
