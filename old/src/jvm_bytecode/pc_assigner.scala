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


// NOTE/WARNING: the algorithm extremely inefficient/naive, but should be ok
// since # segment << # of ops.  Fix/optimize later as needed.
class SegmentIdAssigner(root: CodeScope) {
    if (root._parentScope != null) {
        throw new Exception("not root ...")
    }

    var rootSection = root

    // section's map id -> section
    var sectionMap: HashMap[Int, CodeScope] = null

    // section's map id -> stack
    var stacksMap: HashMap[Int, Stack[CodeBlock]] = null

    var scopeStack: Stack[CodeScope] = null

    var currentScope: CodeScope = null
    var currentStack: Stack[CodeBlock] = null

    var nextSegmentId = 1

    def _init() {
        sectionMap = new HashMap[Int, CodeScope]()
        (new PcIdResetter(rootSection, sectionMap)).apply()

        scopeStack = new Stack[CodeScope]()

        stacksMap = new HashMap[Int, Stack[CodeBlock]]()
        for (i <- sectionMap.keySet()) {
            stacksMap.put(i, new Stack[CodeBlock]())
        }

        val entryBlock = rootSection.getEntryBlock()
        _pushScopeStack(entryBlock._parentScope)
    }

    def assignIds() {
        _init()
        while (!scopeStack.isEmpty()) {
            if (currentStack.isEmpty()) {
                _updateStacks()
            } else {
                _assignId()
            }
        }
    }

    def _pushScopeStack(s: CodeScope) {
        var nestedScopes = new Stack[CodeScope]()
        var tmp = s
        while (tmp != currentScope) {
            nestedScopes.push(tmp)
            tmp = tmp._parentScope
        }

        currentScope = s
        currentStack = stacksMap.get(s._unorderedId)

        while (!nestedScopes.isEmpty()) {
            scopeStack.push(nestedScopes.pop())
        }
    }

    def _pushSegment(seg: CodeSegment) {
        seg match {
            case b: CodeBlock => currentStack.push(b)
            case s: CodeScope => _pushScopeStack(s)
        }
    }

    def _updateStacks() {
        var minId = nextSegmentId

        if (currentScope._entryPoint.segmentId < 0) {
            _pushSegment(currentScope._entryPoint)
            return
        }

        for (seg <- currentScope._segments) {
            if (seg.segmentId < 0) {
                _pushSegment(seg)
                return
            }
            if (minId > seg.segmentId) {
                minId = seg.segmentId
            }
        }

        currentScope.segmentId = minId
        scopeStack.pop()

        if (scopeStack.isEmpty()) {
            currentScope = null
            currentStack = null
        } else {
            currentScope = scopeStack.peek()
            currentStack = stacksMap.get(currentScope._unorderedId)
        }
    }

    def _assignId() {
        var block = currentStack.pop()
        if (block.segmentId > 0) {  // already visited
            return
        }

        block.segmentId = nextSegmentId
        nextSegmentId += 1

        var candidates = new Vector[CodeBlock]()
        block._ops.lastElement() match {
            case g: Goto => candidates.add(g._targetBlock)
            case i: IfBaseOp => {
                candidates.add(i._ifBranch)

                // else branch should be an implicit goto and must be next
                // to the current block
                if (i._elseBranch._parentScope != currentScope) {
                    throw new Exception("unexpected")
                }
                if (i._elseBranch._ops.size() != 1) {
                    throw new Exception("unexpected")
                }
                i._elseBranch._ops.lastElement() match {
                    case g: Goto => {
                        i._elseBranch.segmentId = nextSegmentId
                        nextSegmentId += 1
                        candidates.add(g._targetBlock)
                    }
                    case _ => throw new Exception("unexpected")
                }
            }
            case s: Switch => {
                for (block <- s._table.values()) {
                    candidates.add(block)
                }
                candidates.add(s._defaultBranch)
            }
            case _ => {}
        }

        var candidateScope: CodeScope = null
        for (block <- candidates) {
            if (block.segmentId < 0) {
                val blockScope = block._parentScope
                stacksMap.get(blockScope._unorderedId).push(block)
                if (currentScope._contains(blockScope)) {
                    candidateScope = blockScope
                }
            }
        }

        if (candidateScope != null && currentScope != candidateScope) {
            _pushScopeStack(candidateScope)
        }
    }
}

class AddressCounter extends DataOutput {
    var pos = 0

    def write(b: Array[Byte]) {
        pos += b.length
    }

    def write(b: Array[Byte], off: Int, len: Int) {
        if (len < 0) {
            throw new Exception("bad len")
        }
        if (len < 0 || (off + len) > b.length) {
            throw new Exception("bad offset/len")
        }
        pos += len
    }

    def write(b: Int) {
        pos += 1
    }

    def writeByte(v: Int) {
        pos += 1
    }

    def writeInt(v: Int) {
        pos += 4
    }

    def writeShort(v: Int) {
        pos += 2
    }

    def writeBoolean(v: Boolean) {
        throw new Exception("Not supported")
    }

    def writeBytes(v: String) {
        throw new Exception("Not supported")
    }

    def writeChar(v: Int) {
        throw new Exception("Not supported")
    }

    def writeChars(v: String) {
        throw new Exception("Not supported")
    }

    def writeDouble(v: Double) {
        throw new Exception("Not supported")
    }

    def writeFloat(v: Float) {
        throw new Exception("Not supported")
    }

    def writeLong(v: Long) {
        throw new Exception("Not supported")
    }

    def writeUTF(v: String) {
        throw new Exception("Not supported")
    }
}

object PcAssigner {
    def assignSegmentIdsAndPcs(root: CodeScope): Vector[CodeBlock] = {
        var segmentIdAssigner = new SegmentIdAssigner(root)
        segmentIdAssigner.assignIds()

        var blocks = new Vector[CodeBlock]()
        root._collectBlocks(blocks)
        Collections.sort(blocks)

        // NOTE: this single pass approach does not work if we need to support
        // goto_w
        var counter = new AddressCounter()
        for (block <- blocks) {
            block.pc = counter.pos
            for (op <- block._ops) {
                op.pc = counter.pos
                op.serialize(counter)
            }
            block._endPc = counter.pos
        }

        (new ScopePcUpdater(root)).apply()

        return blocks
    }
}
