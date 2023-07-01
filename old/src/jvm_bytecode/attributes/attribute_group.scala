import java.io.DataInputStream
import java.io.DataOutputStream
import java.util.Vector

import scala.collection.JavaConversions._


abstract class AttributeGroup(o: AttributeOwner) {
    var _owner: AttributeOwner = o

    var _unsupported = new Vector[UnsupportedAttribute]()

    def unsupported(): Vector[UnsupportedAttribute] = _unsupported

    def allAttributes(): Vector[Attribute]

    def debugString(indent: String): String = {
        val attributes = allAttributes()

        if (attributes.isEmpty()) {
            return indent + "(no attributes)\n"
        }

        var result = ""
        for (attr <- attributes) {
            result += attr.debugString(indent) + "\n"
        }

        return result
    }

    def serialize(output: DataOutputStream) {
        val attributes = allAttributes()

        output.writeShort(attributes.size())
        for (attr <- attributes) {
            attr.serialize(output)
        }
    }

    def _readAttributes(inputStream: DataInputStream): Vector[Attribute] = {
        var attributes = new Vector[Attribute]()

        val attrCount = inputStream.readUnsignedShort()
        for (_ <- 1 to attrCount) {
            val name = _owner.constants().getUtf8ByIndex(
                    inputStream.readUnsignedShort())
            var attr = name.value() match {
/* TODO implement more attributes
AnnotationDefault
MethodParameters
StackMapTable

RuntimeVisibleParameterAnnotations
RuntimeInvisibleParameterAnnotations
RuntimeVisibleAnnotations
RuntimeInvisibleAnnotations
RuntimeVisibleTypeAnnotations
RuntimeInvisibleTypeAnnotations
*/
                case "BootstrapMethods" => new BootstrapMethodsAttribute(_owner)
                case "Code" => new CodeAttribute(_owner)
                case "ConstantValue" => new ConstantValueAttribute(_owner)
                case "Deprecated" => new DeprecatedAttribute(_owner)
                case "EnclosingMethod" => new EnclosingMethodAttribute(_owner)
                case "Exceptions" => new ExceptionsAttribute(_owner)
                case "InnerClasses" => new InnerClassesAttribute(_owner)
                case "LineNumberTable" => new LineNumberTableAttribute(_owner)
                case "LocalVariableTable" =>
                        new LocalVariableTableAttribute(_owner)
                case "LocalVariableTypeTable" =>
                        new LocalVariableTypeTableAttribute(_owner)
                case "Signature" => new SignatureAttribute(_owner)
                case "SourceDebugExtension" =>
                        new SourceDebugExtensionAttribute(_owner)
                case "SourceFile" => new SourceFileAttribute(_owner)
                case "Synthetic" => new SyntheticAttribute(_owner)
                case _ => new UnsupportedAttribute(_owner)
            }

            attr.deserialize(name, inputStream.readInt(), inputStream)
            attributes.add(attr)
        }

        return attributes
    }

    def deserialize(input: DataInputStream)
}

class ClassAttributes(c: ClassInfo) extends AttributeGroup(c) {
    var _sourceFile: SourceFileAttribute = null
    var _signature: SignatureAttribute = null
    var _innerClasses: InnerClassesAttribute = null
    var _enclosingMethod: EnclosingMethodAttribute = null
    var _bootstrapMethods: BootstrapMethodsAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null
    var _sourceDebugExtension: SourceDebugExtensionAttribute = null

    def sourceFile(): String = {
        if (_sourceFile == null) {
            return null
        }
        return _sourceFile.value()
    }
    def setSourceFile(s: String) {
        if (s == null) {
            _sourceFile = null
        } else {
            _sourceFile = new SourceFileAttribute(_owner, s)
        }
    }

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def innerClasses(): InnerClassesAttribute = innerClasses
    def addInnerClass(
            innerClass: String,
            outerClass: String,
            innerName: String): InnerClassEntry = {
        val inner = new InnerClassEntry(
                _owner,
                innerClass,
                outerClass,
                innerName)

        if (_innerClasses == null) {
            _innerClasses = new InnerClassesAttribute(_owner)
        }

        _innerClasses.add(inner)
        return inner
    }

    def enclosingMethod(): EnclosingMethodAttribute = _enclosingMethod
    def setEnclosingMethod(
            className: String,
            methodName: String,
            methodType: MethodType) {
        _enclosingMethod = new EnclosingMethodAttribute(
                _owner,
                className,
                methodName,
                methodType)
    }
    def clearEnclosingMethod() {
        _enclosingMethod = null
    }

    def bootstrapMethods(): BootstrapMethodsAttribute = _bootstrapMethods
    def addBootstrapMethods(
            mh: ConstMethodHandleInfo): BootstrapMethodEntry = {
        if (_bootstrapMethods == null) {
            _bootstrapMethods = new BootstrapMethodsAttribute(_owner)
        }
        return _bootstrapMethods.add(mh)
    }
    def clearBootstrapMethods() {
        _bootstrapMethods = null
    }

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def sourceDebugExtension(): String = {
        return new String(_sourceDebugExtension.bytes(), "UTF-8")
    }
    def setSourceDebugExtension(s: String) {
        if (s == null) {
            _sourceDebugExtension = null
        } else {
            _sourceDebugExtension = new SourceDebugExtensionAttribute(_owner, s)
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO add more attributes

        if (_sourceFile != null) {
            allAttributes.add(_sourceFile)
        }
        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_innerClasses != null) {
            allAttributes.add(_innerClasses)
        }
        if (_enclosingMethod != null) {
            allAttributes.add(_enclosingMethod)
        }
        if (_bootstrapMethods != null) {
            allAttributes.add(_bootstrapMethods)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
        }
        if (_sourceDebugExtension != null) {
            allAttributes.add(_sourceDebugExtension)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO add more attributes
                case attr: BootstrapMethodsAttribute => {
                    if (_bootstrapMethods != null) {
                        throw new Exception(
                                "multiple bootstrap methods attribute")
                    }
                    _bootstrapMethods = attr
                }
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: EnclosingMethodAttribute => {
                    if (_enclosingMethod != null) {
                        throw new Exception(
                                "multiple enclosing method attribute")
                    }
                    _enclosingMethod = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: InnerClassesAttribute => {
                    if (_innerClasses != null) {
                        throw new Exception("multiple inner classes attribute")
                    }
                    _innerClasses = attr
                }
                case attr: SourceDebugExtensionAttribute => {
                    if (_sourceDebugExtension != null) {
                        throw new Exception(
                                "multiple source debug extension attribute")
                    }
                    _sourceDebugExtension = attr
                }
                case attr: SourceFileAttribute => {
                    if (_sourceFile != null) {
                        throw new Exception("multiple sourceFile attribute")
                    }
                    _sourceFile = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected class attribute: " + a.name())
            }
        }
    }
}

class FieldAttributes(f: FieldInfo) extends AttributeGroup(f) {
    var _signature: SignatureAttribute = null
    var _constantValue: ConstantValueAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def constantValue(): ConstantValueAttribute = _constantValue

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO add more attributes

        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO add more attributes
                case attr: ConstantValueAttribute => {
                    if (_constantValue != null) {
                        throw new Exception("multiple constant value attribute")
                    }
                    _constantValue = attr
                }
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected field attribute: " + a.name())
            }
        }
    }
}

class MethodAttributes(m: MethodInfo) extends AttributeGroup(m) {
    var _signature: SignatureAttribute = null
    var _exceptions: ExceptionsAttribute = null
    var _code: CodeAttribute = null
    var _deprecated: DeprecatedAttribute = null
    var _synthetic: SyntheticAttribute = null

    def signature(): String = {
        if (_signature == null) {
            return null
        }
        return _signature.value()
    }
    def setSignature(s: String) {
        if (s == null) {
            _signature = null
        } else {
            _signature = new SignatureAttribute(_owner, s)
        }
    }

    def exceptions(): Vector[String] = {
        if (_exceptions == null) {
            return null
        }
        return _exceptions.exceptions()
    }
    def addException(exceptionName: String) {
        if (_exceptions == null) {
            _exceptions = new ExceptionsAttribute(_owner)
        }
        _exceptions.add(exceptionName)
    }

    def code(): CodeAttribute = _code

    def isDeprecated(): Boolean = _deprecated != null
    def setIsDeprecated(b: Boolean) {
        if (b) {
            _deprecated = new DeprecatedAttribute(_owner)
        } else {
            _deprecated = null
        }
    }

    def isSynthetic(): Boolean = _synthetic != null
    def setIsSynthetic(b: Boolean) {
        if (b) {
            _synthetic = new SyntheticAttribute(_owner)
        } else {
            _synthetic = null
        }
    }

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO add more attributes
        if (_signature != null) {
            allAttributes.add(_signature)
        }
        if (_exceptions != null) {
            allAttributes.add(_exceptions)
        }
        if (_code != null) {
            allAttributes.add(_code)
        }
        if (_deprecated != null) {
            allAttributes.add(_deprecated)
        }
        if (_synthetic != null) {
            allAttributes.add(_synthetic)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO add more attributes
                case attr: CodeAttribute => {
                    if (_code != null) {
                        throw new Exception("multiple code attribute")
                    }
                    _code = attr
                }
                case attr: DeprecatedAttribute => {
                    if (_deprecated != null) {
                        throw new Exception("multiple deprecated attribute")
                    }
                    _deprecated = attr
                }
                case attr: ExceptionsAttribute => {
                    if (_exceptions != null) {
                        throw new Exception("multiple exceptions attribute")
                    }
                    _exceptions = attr
                }
                case attr: SignatureAttribute => {
                    if (_signature != null) {
                        throw new Exception("multiple signature attribute")
                    }
                    _signature = attr
                }
                case attr: SyntheticAttribute => {
                    if (_synthetic != null) {
                        throw new Exception("multiple synthetic attribute")
                    }
                    _synthetic = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected method attribute: " + a.name())
            }
        }
    }
}

class CodeAttributes(c: CodeAttribute) extends AttributeGroup(c) {
    var lineNumberTable: LineNumberTableAttribute = null

    var localVariableTable: LocalVariableTableBaseAttribute = null

    def allAttributes(): Vector[Attribute] = {
        var allAttributes = new Vector[Attribute]()

        // TODO add more attributes

        if (lineNumberTable != null) {
            allAttributes.add(lineNumberTable)
        }

        if (localVariableTable != null) {
            allAttributes.add(localVariableTable)
        }

        for (attr <- _unsupported) {
            allAttributes.add(attr)
        }

        return allAttributes
    }

    def deserialize(input: DataInputStream) {
        for (a <- _readAttributes(input)) {
            a match {
                // TODO add more attributes
                case attr: LineNumberTableAttribute => {
                    if (lineNumberTable != null) {
                        lineNumberTable.mergeFrom(attr)
                    } else {
                        lineNumberTable = attr
                    }
                }
                case attr: LocalVariableTableBaseAttribute => {
                    if (localVariableTable != null) {
                        throw new Exception(
                                "multiple local variable table attribute")
                    }
                    localVariableTable = attr
                }
                case attr: UnsupportedAttribute => _unsupported.add(attr)
                case _ => throw new Exception(
                        "Unexpected method attribute: " + a.name())
            }
        }
    }
}
