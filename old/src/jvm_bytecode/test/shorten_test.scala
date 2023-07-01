object ShortenTest {
    def main(arg: Array[String]) {
        var c = new ClassInfo("ClassName")
        var m = c._methods.add("func", new MethodType())
        var a = m._attributes
        a._code = new CodeAttribute(m)

        var code = new CodeScope(m, null, 0)
        a._code.code = code

        var b1 = code.newBlock()
        var s2 = code.newSubSection()
        var s3 = code.newSubSection()

        var b21 = s2.newBlock()
        var s22 = s2.newSubSection()
        var b23 = s2.newBlock()

        var b221 = s22.newBlock()

        var b31 = s3.newBlock()
        var b32 = s3.newBlock()
        var s33 = s3.newSubSection()

        var b331 = s33.newBlock()
        var s332 = s33.newSubSection()

        var b3321 = s332.newBlock()
        b3321.pushI(123)
        var b3322 = s332.newBlock()
        b3322.returnVoid()

        var e21 = s2.newExceptionHandle(null)
        e21.newBlock().returnI()

        var e3a = s3.newExceptionHandle(null)
        e3a.newBlock().returnL()

        var e3b = s3.newExceptionHandle("Throwable")
        e3b.newBlock().returnL()

        s3.implicitGoto = null

        var e22 = s22.newExceptionHandle(null)
        e22.newBlock().returnI()

        b23.implicitGoto = null

        var e332 = s332.newExceptionHandle(null)
        var e332s = e332.newSubSection()
        e332s.newBlock().pushI(321)
        e332s.newBlock().returnVoid()

        c.analyze()

        PcAssigner.assignSegmentIdsAndPcs(code)
        println(code.debugString(""))
    }
}
