class Test {
    var intValue: Int = 1
    var longValue: Long = 2
    var floatValue: Float = 3
    var doubleValue: Double = 4
    var stringValue: String = "Hello world"

    def f(): Int = intValue

    def g(): Long = f().toLong

    def h(): Double = g().toDouble

    def h2(): Double = g().toDouble + h()
}

trait Trait1 {
    def t1(): Int = 1
}

trait Trait2 {
    def t2(): Int = 2
}

abstract class Abstract1 {
    def t3(): Int = 3
}

class Real extends Abstract1 with Trait1 with Trait2 {}
