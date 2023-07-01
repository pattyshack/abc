import scala.collection.JavaConversions._


trait AnalysisPass {
    def apply(c: ClassInfo)
}

trait CodeAnalysisPass extends AnalysisPass {
    var _class: ClassInfo = null
    var _method: MethodInfo = null
    var _codeAttr: CodeAttribute = null

    def apply(c: ClassInfo) {
        _class = c
        for (m <- c.methods()) {
            _method = m
            _codeAttr = m._attributes._code
            if (_codeAttr != null) {
                analyze(_codeAttr.code)
            }
        }
    }

    def analyze(rootScope: CodeScope)
}

