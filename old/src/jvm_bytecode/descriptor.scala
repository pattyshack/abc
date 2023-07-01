import java.util.Vector

import scala.collection.JavaConversions._


trait DescriptorType extends Comparable[DescriptorType] {
    def descriptorString(): String

    def compareTo(other: DescriptorType): Int = {
        return descriptorString().compareTo(other.descriptorString())
    }
}

trait FieldType extends DescriptorType {
    // see table 2.3 / section 2.11.1
    def isCat1(): Boolean
    def isCat2(): Boolean

    // see table 2.3 / page 29
    def categorySize(): Int = {
        if (isCat1()) {
            return 1
        }
        if (isCat2()) {
            return 2
        }
        return 0
    }
}

trait BaseType extends FieldType {
    def arrayType(): Int
}

trait RefType extends FieldType {
    def isCat1(): Boolean = true
    def isCat2(): Boolean = false
}

trait BaseIntType extends BaseType {
    def isCat1(): Boolean = true
    def isCat2(): Boolean = false
}

// NOTE: BoolType is not used for stack frame computation.
class BoolType extends BaseIntType {
    def descriptorString(): String = "Z"

    def arrayType(): Int = 4
}

// NOTE: ByteType is not used for stack frame computation.
class ByteType extends BaseIntType {
    def descriptorString(): String = "B"

    def arrayType(): Int = 8
}

// NOTE: CharType is not used for stack frame computation.
class CharType extends BaseIntType {
    def descriptorString(): String = "C"

    def arrayType(): Int = 5
}

// NOTE: ShortType is not used for stack frame computation.
class ShortType extends BaseIntType {
    def descriptorString(): String = "S"

    def arrayType(): Int = 9
}

class IntType extends BaseIntType {
    def descriptorString(): String = "I"

    def arrayType(): Int = 10
}

class DoubleType extends BaseType {
    def isCat1(): Boolean = false
    def isCat2(): Boolean = true

    def descriptorString(): String = "D"

    def arrayType(): Int = 7
}

class FloatType extends BaseType {
    def isCat1(): Boolean = true
    def isCat2(): Boolean = false

    def descriptorString(): String = "F"

    def arrayType(): Int = 6
}

class LongType extends BaseType {
    def isCat1(): Boolean = false
    def isCat2(): Boolean = true

    def descriptorString(): String = "J"

    def arrayType(): Int = 11
}

class ArrayType(t: FieldType) extends RefType {
    val itemType = t

    def descriptorString(): String = "[" + itemType.descriptorString()
}

class ObjectType(i: Boolean, s: String) extends RefType {
    def this(s: String) = this(false, s)

    var isUninitialized = i

    val name = s

    def descriptorString(): String = "L" + name + ";"

    def isJavaString(): Boolean = name == Const.JAVA_STRING

    def isJavaObject(): Boolean = name == Const.JAVA_OBJECT
}

class ParameterTypes extends Comparable[ParameterTypes] {
    var _parameters = new Vector[FieldType]()

    def add(field: FieldType) {
        _parameters.add(field)
    }

    def size(): Int = _parameters.size()

    def descriptorString(): String = {
        var result = "("
        for (p <- _parameters) {
            result += p.descriptorString()
        }
        result += ")"

        return result
    }

    def compareTo(other: ParameterTypes): Int = {
        return descriptorString().compareTo(other.descriptorString())
    }
}

class MethodType extends DescriptorType {
    var parameters = new ParameterTypes()
    var returnType: FieldType = null  // null for void

    def descriptorString(): String = {
        var result = parameters.descriptorString()

        if (returnType != null) {
            result += returnType.descriptorString()
        } else {
            result += "V"
        }
        return result
    }
}

//
// Only usable for stack frame verification
//

// NOTE: will write unusable type as top type in stack map table attribute
class UnusableType extends FieldType {
    def isCat1(): Boolean = false
    def isCat2(): Boolean = false

    def descriptorString(): String = "__unusable__"
}

class TopType extends FieldType {
    def isCat1(): Boolean = false
    def isCat2(): Boolean = false

    def descriptorString(): String = "__top__"
}

class NullType extends RefType {
    def descriptorString(): String = "__null__"  // fake type
}

// only used for expected type
class CheckRefType extends RefType {
    def descriptorString(): String = "__ref__"  // fake type
}

// only used for expected type
class CheckArrayType extends ArrayType(null) {
    override def descriptorString(): String = "__array__"  // fake type
}

//
// maybe use a real lexer/parser ...
//

object DescriptorToken {
    val BYTE = "B"
    val CHAR = "C"
    val DOUBLE = "D"
    val FLOAT = "F"
    val INT = "I"
    val LONG = "J"
    val OBJECT = "L"
    val SHORT = "S"
    val BOOL = "Z"
    val ARRAY = "["
    val VOID = "V"
    val LPARAM = "("
    val RPARAM = ")"
    val EOF = ""
}

class DescriptorTokenizer(d: String) {
    var descriptorString = d

    var nextPos = 0
    var value = ""

    def lookAhead(): Char = descriptorString.charAt(nextPos)

    def nextToken(): String = {
        value = ""
        if (nextPos >= descriptorString.length) {
            return DescriptorToken.EOF
        }

        value += descriptorString.charAt(nextPos)
        nextPos += 1
        value match {
            case DescriptorToken.BYTE => return DescriptorToken.BYTE
            case DescriptorToken.CHAR => return DescriptorToken.CHAR
            case DescriptorToken.DOUBLE => return DescriptorToken.DOUBLE
            case DescriptorToken.FLOAT => return DescriptorToken.FLOAT
            case DescriptorToken.INT => return DescriptorToken.INT
            case DescriptorToken.LONG => return DescriptorToken.LONG
            case DescriptorToken.SHORT => return DescriptorToken.SHORT
            case DescriptorToken.BOOL => return DescriptorToken.BOOL
            case DescriptorToken.ARRAY => return DescriptorToken.ARRAY
            case DescriptorToken.OBJECT => {
                val end = descriptorString.indexOf(';', nextPos - 1)
                if (end == -1) {
                    throw new Exception(
                            "malformed descriptor: " + descriptorString)
                }
                value = descriptorString.substring(nextPos, end)
                nextPos = end + 1
                return DescriptorToken.OBJECT
            }
            case DescriptorToken.VOID => return DescriptorToken.VOID
            case DescriptorToken.LPARAM => return DescriptorToken.LPARAM
            case DescriptorToken.RPARAM => return DescriptorToken.RPARAM
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }
}

class DescriptorParser(s: String) {
    val descriptorString = s
    var tokenizer = new DescriptorTokenizer(s)

    def _parseField(): FieldType = {
        tokenizer.nextToken() match {
            case DescriptorToken.BYTE => return new ByteType()
            case DescriptorToken.CHAR => return new CharType()
            case DescriptorToken.DOUBLE => return new DoubleType()
            case DescriptorToken.FLOAT => return new FloatType()
            case DescriptorToken.INT => return new IntType()
            case DescriptorToken.LONG => return new LongType()
            case DescriptorToken.SHORT => return new ShortType()
            case DescriptorToken.BOOL => return new BoolType()
            case DescriptorToken.ARRAY => return new ArrayType(_parseField())
            case DescriptorToken.OBJECT =>
                    return new ObjectType(tokenizer.value)
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }

    def parseFieldDescriptor(): FieldType = {
        val f = _parseField()
        tokenizer.nextToken() match {
            case DescriptorToken.EOF => return f
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }

    def parseMethodDescriptor(): MethodType = {
        if (tokenizer.nextToken() != DescriptorToken.LPARAM) {
            throw new Exception("malformed descriptor: " + descriptorString)
        }

        var method = new MethodType()

        var shouldContinue = true
        while (shouldContinue) {
            try {
                if (tokenizer.lookAhead() == ')') {
                    tokenizer.nextToken()
                    shouldContinue = false
                } else {
                    method.parameters.add(_parseField())
                }
            } catch {
                case ex: IndexOutOfBoundsException => throw new Exception(
                        "malformed descriptor: " + descriptorString)
            }
        }

        try {
            if (tokenizer.lookAhead() == 'V') {
                tokenizer.nextToken()
            } else {
                method.returnType = _parseField()
            }
        } catch {
            case ex: IndexOutOfBoundsException => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }

        tokenizer.nextToken() match {
            case DescriptorToken.EOF => return method
            case _ => throw new Exception(
                    "malformed descriptor: " + descriptorString)
        }
    }
}

