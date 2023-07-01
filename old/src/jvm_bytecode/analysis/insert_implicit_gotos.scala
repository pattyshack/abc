import java.util.Vector

import scala.collection.JavaConversions._


class InsertImplicitGotos extends CodeAnalysisPass {
    def analyze(rootScope: CodeScope) {
        new ImplicitGotoInserter(rootScope).apply()
    }
}

class ImplicitGotoInserter(root: CodeScope) extends CodeVisitor(root) {
    var indirections = new Vector[CodeBlock]()

    override def visitBlock(block: CodeBlock) {
        if (!block._hasControlOp) {
            val target = block.getImplicitGoto()
            if (target == null) {
                throw new Exception("no implicit goto - pc: " + block.pc)
            }
            block.goto(target.getEntryBlock())
            return
        }

        var lastOp = block._ops.lastElement()
        lastOp match {
            // this simplifies control flow flattening since we no
            // longer need to pair the else branch immediately after the
            // condition operation.
            case i: IfBaseOp => {
                var indirection = new CodeBlock(
                        block._owner,
                        block._parentScope)
                indirection.lineContext = i.line
                indirection.goto(i._elseBranch)
                i._elseBranch = indirection
                indirections.add(indirection)
            }
            case _ => {}
        }
    }

    override def apply() {
        super.apply()

        for (block <- indirections) {
            block._parentScope._addBlock(block)
        }
    }
}
