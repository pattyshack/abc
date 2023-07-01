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


class ExceptionTarget(e: ConstClassInfo, t: CodeScope) {
    var exception: ConstClassInfo = e
    var target: CodeScope = t
}

class NamedLocalEntry(n: String, t: FieldType, i: Int) {
    val name = n
    val fieldType = t
    val index = i
}

// a group of code segments which will be written out as a single continuous
// unit.
//
// Also acts as lexical / exception try scope
//
// NOTE: scope's implicit goto is only used by the last segment. All other
// segments will implicitly goto the next segment in the scope (useless
// implicitGoto is explicitly set)
class CodeScope(
        owner: AttributeOwner,
        parent: CodeScope,
        nextNamedLocal: Int) extends CodeSegment(owner, parent) {
    def this(parent: CodeScope) = this(
            parent._owner,
            parent,
            parent._nextNamedLocal)

    var _segments = new Vector[CodeSegment]()
    var _blocks = new Vector[CodeBlock]()
    var _subsections = new Vector[CodeScope]()

    var _exceptionTargets = new Vector[ExceptionTarget]()

    var _entryPoint: CodeSegment = null

    var _nextNamedLocal = nextNamedLocal  // disable if -1
    var _namedLocals = new HashMap[String, NamedLocalEntry]()

    def _disableNamedLocals() {
        if (_nextNamedLocal >= 0) {
            for (s <- _subsections) {
                s._disableNamedLocals()
            }
            _nextNamedLocal = -1
        }
    }

    def defineLocal(name: String, fieldType: FieldType) {
        if (_nextNamedLocal < 0) {
            throw new Exception("named local variable disabled")
        }

        if (_namedLocals.containsKey(name)) {
            throw new Exception(name + " declared multiple times in same scope")
        }

        _namedLocals.put(
                name,
                new NamedLocalEntry(name, fieldType, _nextNamedLocal))
        _nextNamedLocal += fieldType.categorySize()
    }

    def getLocal(name: String): NamedLocalEntry = {
        if (_nextNamedLocal < 0) {
            throw new Exception("named local variable disabled")
        }

        val entry = _namedLocals.get(name)
        if (entry != null) {
            return entry
        }

        if (_parentScope == null) {
            throw new Exception(name + " not defined")
        }

        return _parentScope.getLocal(name)
    }

    def getEntryBlock(): CodeBlock = {
        if (_entryPoint == null) {
            return newBlock()
        }

        _entryPoint match {
            case b: CodeBlock => return b
            case s: CodeScope => return s.getEntryBlock()
        }
    }

    def newBlock(): CodeBlock = {
        var block = new CodeBlock(_owner, this)
        return _addBlock(block)
    }

    def _addSegment(seg: CodeSegment) {
        if (_entryPoint == null) {
            _entryPoint = seg
        } else {
            var last = _segments.lastElement()
            if (last.implicitGoto == null) {
                last.implicitGoto = seg
            }
        }

        _segments.add(seg)
    }

    def _addBlock(block: CodeBlock): CodeBlock = {
        _addSegment(block)
        _blocks.add(block)
        return block
    }

    def newSubSection(): CodeScope = {
        val section = new CodeScope(this)
        _addSegment(section)
        _subsections.add(section)
        return section
    }

    // only used for reconstruction
    def _getMostSpecificSection(pc: Int): CodeScope = {
        for (s <- _subsections) {
            if (s.pc <= pc && pc < s._endPc) {
                return s._getMostSpecificSection(pc)
            }
        }

        return this
    }

    // add new exception handle for the current section.
    // NOTE:
    //   - call order matters.
    //   - pass in null to catch all
    //   - cannot use this on top scope
    def newExceptionHandle(exceptionClassName: String): CodeScope = {
        if (_parentScope == null) {
            throw new Exception(
                    "cannot create exception handle for top level scope")
        }

        var exception: ConstClassInfo = null
        if (exceptionClassName != null) {
            exception =_owner.constants().getClass(exceptionClassName)
        }

        val target = _parentScope.newSubSection()

        _exceptionTargets.add(new ExceptionTarget(exception, target))
        return target
    }

    // For sharing exception section (needed for javac)
    def shareExceptionHandle(exceptionClassName: String, target: CodeScope) {
        if (_parentScope == null) {
            throw new Exception(
                    "cannot create exception handle for top level scope")
        }

        var exception: ConstClassInfo = null
        if (exceptionClassName != null) {
            exception =_owner.constants().getClass(exceptionClassName)
        }

        _exceptionTargets.add(new ExceptionTarget(exception, target))
    }

    def generateLineNumberTable(): LineNumberTableAttribute = {
        var table = new TreeMap[Int, Int]()

        var blocks = new Vector[CodeBlock]()
        _collectBlocks(blocks)
        Collections.sort(blocks)

        var currLine = -1
        for (block <- blocks) {
            for (op <- block._ops) {
                if (op.line >= 0 && op.line != currLine) {
                    table.put(op.pc, op.line)
                    currLine = op.line
                }
            }
        }

        if (table.isEmpty()) {
            return null
        }

        var attr = new LineNumberTableAttribute(_owner)
        attr.table = table
        return attr
    }

    // this assumes pc are assigned and segments are sorted
    def _collectExceptionEntries(result: Vector[ExceptionEntry]) {
        for (section <- _subsections) {
            section._collectExceptionEntries(result)
        }

        for (entry <- _exceptionTargets) {
            result.add(new ExceptionEntry(
                    _owner,
                    pc,
                    _endPc,
                    entry.target.getEntryBlock().pc,
                    entry.exception))
        }
    }

    // returns true if this equals, or contains, the other section
    def _contains(other: CodeScope): Boolean = {
        if (this == other) {
            return true
        }

        var tmp = other
        while (tmp._parentScope != null) {
            if (this == tmp._parentScope) {
                return true
            }
            tmp = tmp._parentScope
        }

        return false
    }

    def _collectBlocks(result: Vector[CodeBlock]) {
        for (seg <- _segments) {
            seg match {
                case b: CodeBlock => result.add(b)
                case s: CodeScope => s._collectBlocks(result)
            }
        }
    }

    def serialize(output: DataOutput) {
        var blocks = PcAssigner.assignSegmentIdsAndPcs(this)

        if (_endPc > Const.UINT16_MAX) {
            // code attributes can't handle code length > 64K ...
            throw new Exception("code too large")
        }

        output.writeInt(_endPc)
        for (block <- blocks) {
            block.serialize(output)
        }

        var exceptions = new Vector[ExceptionEntry]()
        _collectExceptionEntries(exceptions)

        output.writeShort(exceptions.size())
        for (entry <- exceptions) {
            entry.serialize(output)
        }
    }

    def deserialize(startAddress: Int, opCode: Int, input: DataInputStream) {
        throw new Exception("cannot directly deserialize code section")
    }

    def sort() {
        Collections.sort(_segments)
        Collections.sort(_blocks)
        Collections.sort(_subsections)

        for (section <- _subsections) {
            section.sort()
        }
    }

    def debugString(indent: String): String = {
        sort()

        var result = indent + "Section (pc: [" + pc + ", " + _endPc +
                ") segment: " + segmentId + " entry pt: " +
                _entryPoint.segmentId + " reachable: " + _reachable + ")\n"
        for (segment <- _segments) {
            result += segment.debugString(indent + "  ")
        }

        return result
    }
}

