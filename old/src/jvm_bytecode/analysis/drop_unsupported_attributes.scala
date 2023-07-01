import scala.collection.JavaConversions._


class DropUnsupportedAttributes extends AnalysisPass {
    def apply(c: ClassInfo) {
        _clearUnsupported(c._attributes)
        for (field <- c.fields()) {
            _clearUnsupported(field._attributes)
        }
        for (method <- c.methods()) {
            var attrs = method._attributes
            _clearUnsupported(attrs)
            if (attrs._code != null) {
                _clearUnsupported(attrs._code.attributes)
            }
        }
    }

    def _clearUnsupported(g: AttributeGroup) {
        for (attr <- g._unsupported) {
            println("Dropping unsupported attribute: " + attr.name())
        }
        g._unsupported.clear()
    }
}
