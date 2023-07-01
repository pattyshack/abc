import java.util.Collection
import java.util.HashMap
import java.util.Stack
import java.util.Vector

import scala.collection.JavaConversions._


class RemoveUnreachableSegments extends CodeAnalysisPass {
    def analyze(root: CodeScope) {
        new ReachableResetter(root).apply()

        var scopeMap = new HashMap[Int, CodeScope]()
        new PcIdResetter(root, scopeMap).apply()

        var candidates = new Stack[CodeBlock]()
        candidates.push(root.getEntryBlock())

        while (!candidates.isEmpty()) {
            var block = candidates.pop()
            markAsReachable(block)

            followBlock(block, candidates)

            if (candidates.isEmpty()) {
                generateExceptionCandidates(scopeMap, candidates)
            }
        }

        new UnreachableSegmentsRemover(root).apply()
    }

    def markAsReachable(block: CodeBlock) {
        block._reachable = true
        var scope = block._parentScope
        while (scope != null) {
            scope._reachable = true
            scope = scope._parentScope
        }
    }

    def followBlock(block: CodeBlock, candidates: Stack[CodeBlock]) {
        block._ops.lastElement() match {
            case g: Goto => {
                if (!g._targetBlock._reachable) {
                    candidates.add(g._targetBlock)
                }
            }
            case i: IfBaseOp => {
                if (!i._ifBranch._reachable) {
                    candidates.add(i._ifBranch)
                }
                if (!i._elseBranch._reachable) {
                    candidates.add(i._elseBranch)
                }
            }
            case s: Switch => {
                if (!s._defaultBranch._reachable) {
                    candidates.add(s._defaultBranch)
                }
                for (branch <- s._table.values()) {
                    if (!branch._reachable) {
                        candidates.add(branch)
                    }
                }
            }
            case _ => {}
        }
    }

    def generateExceptionCandidates(
            scopeMap: HashMap[Int, CodeScope],
            candidates: Stack[CodeBlock]) {
        for (scope <- scopeMap.values()) {
            if (scope._reachable) {
                var entry = scope.getEntryBlock()
                if (!entry._reachable) {
                    throw new Exception("unexpected")
                }

                for (entry <- scope._exceptionTargets) {
                    val targetEntry = entry.target.getEntryBlock()
                    if (!targetEntry._reachable) {
                        candidates.add(targetEntry)
                    }
                }
            }
        }
    }
}

class UnreachableSegmentsRemover(root: CodeScope) extends CodeVisitor(root) {
    override def visitScope(scope: CodeScope) {
        var prunedSegments = new Vector[CodeSegment]()
        for (s <- scope._segments) {
            if (s._reachable) {
                prunedSegments.add(s)
            }
        }

        var prunedBlocks = new Vector[CodeBlock]()
        for (s <- scope._blocks) {
            if (s._reachable) {
                prunedBlocks.add(s)
            }
        }

        var prunedSubsections = new Vector[CodeScope]()
        for (s <- scope._subsections) {
            if (s._reachable) {
                prunedSubsections.add(s)
            }
        }

        scope._segments = prunedSegments
        scope._blocks = prunedBlocks
        scope._subsections = prunedSubsections

        super.visitScope(scope)
    }
}

class ReachableResetter(root: CodeScope) extends CodeVisitor(root) {
    override def visitBlock(block: CodeBlock) {
        block._reachable = false
    }

    override def visitScope(scope: CodeScope) {
        super.visitScope(scope)
        scope._reachable = false
    }
}
