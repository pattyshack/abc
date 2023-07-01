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


object CodeScopeReconstructor {
    def reconstruct(
            owner: AttributeOwner,
            exceptions: Vector[ExceptionEntry],
            ops: Vector[Operation]): CodeScope = {

        _checkExceptionOverlaps(exceptions)

        val jumpTargets = _collectJumpTargets(exceptions, ops)

        var result = new CodeScope(owner, null, -1)
        result.pc = 0
        result._endPc = ops.lastElement().pc + 5 // fake it

        _createExceptionSubsections(exceptions, result)

        var pcBlockMap = new TreeMap[Int, CodeBlock]()

        var ifBranches = new HashMap[Int, CodeScope]()

        var prevBlock: CodeBlock = null
        var currBlock: CodeBlock = null
        for (op <- ops) {
            if (jumpTargets.contains(op.pc)) {
                if (prevBlock != null) {
                    prevBlock.implicitGoto = currBlock
                    prevBlock._endPc = currBlock.pc
                }
                prevBlock = currBlock

                var section = result._getMostSpecificSection(op.pc)
                if (prevBlock != null) {
                    val prevScope = prevBlock._parentScope
                    var candidate = ifBranches.get(op.pc)
                    if (candidate != null) {  // is if branch
                        if (candidate._contains(prevScope) &&
                                section._contains(candidate)) {
                            section = candidate
                        }
                    } else {
                        prevBlock._ops.lastElement() match {
                            case i: IfBaseOp => {
                                // make the else branch sit next to the
                                // condition
                                if (section._contains(prevScope)) {
                                    section = prevScope
                                    ifBranches.put(i._tmpOffset + i.pc,
                                                   prevScope)
                                }
                            }
                            case _ => {}
                        }
                    }
                }

                currBlock = section.newBlock()
                currBlock.pc = op.pc
                pcBlockMap.put(op.pc, currBlock)
            }
            currBlock._add(op)
        }
        if (currBlock != null) {
            if (prevBlock != null) {
                prevBlock.implicitGoto = currBlock
                prevBlock._endPc = currBlock.pc
            }
            currBlock._endPc = currBlock.pc + 5  // fake it
        }

        (new ScopePcUpdater(result)).apply()

        for (op <- ops) {
            op.bindBlockRefs(pcBlockMap)
        }

        (new EntryPointSetter(result)).apply()

        return result
    }

    // 1. ensure there are no partial overlaps
    // 2. ensure more specific scope is before less specific scope
    def _checkExceptionOverlaps(exceptions: Vector[ExceptionEntry]) {
        // TODO check exception scope bounds
    }

    def _createExceptionSubsections(
            exceptions: Vector[ExceptionEntry],
            global: CodeScope) {
        // create try sections
        for (i <- (exceptions.size() - 1).to(0, -1)) {
            var entry = exceptions.elementAt(i)

            var section = global._getMostSpecificSection(entry.startPc)
            if (section.pc < entry.startPc || entry.endPc < section._endPc) {
                section = section.newSubSection()
                section.pc = entry.startPc
                section._endPc = entry.endPc
            }
            entry._tmpSection = section
        }

        // create catch sections.
        for (entry <- exceptions) {
            var section = global._getMostSpecificSection(entry.handlerPc)
            if (section.pc < entry.handlerPc) {
                section = section.newSubSection()
                section.pc = entry.handlerPc
                section._endPc = entry.handlerPc + 1  // fake it
            }
            entry._tmpSection.shareExceptionHandle(entry.className(), section)
        }
    }

    def _collectJumpTargets(
            exceptions: Vector[ExceptionEntry],
            ops: Vector[Operation]): TreeSet[Int] = {
        var jumpTargets = new TreeSet[Int]()

        for (ex <- exceptions) {
            jumpTargets.add(ex.startPc)
            jumpTargets.add(ex.endPc)
            jumpTargets.add(ex.handlerPc)
        }

        for (op <- ops) {
            op match {
                case _: CodeSegment => throw new Exception(
                        "cannot group non-basic ops")
                case _: Return => jumpTargets.add(op.pc + 1)
                case _: ReturnValue => jumpTargets.add(op.pc + 1)
                case _: Athrow => jumpTargets.add(op.pc + 1)
                case g: Goto => {
                    jumpTargets.add(g.pc + g._tmpOffset)
                    jumpTargets.add(g.pc + 3)  // next op
                }
                case i: IfBaseOp => {
                    jumpTargets.add(i.pc + i._tmpOffset)  // if branch
                    jumpTargets.add(i.pc + 3)  // else branch
                }
                case s: Switch => {
                    jumpTargets.add(s.pc + s._tmpDefaultOffset)
                    for (offset <- s._tmpOffset.values()) {
                        jumpTargets.add(s.pc + offset)
                    }
                }
                case _ => {}
            }
        }

        jumpTargets.add(0)
        return jumpTargets
    }
}
