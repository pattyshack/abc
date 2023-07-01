import java.io.DataInputStream
import java.io.File
import java.io.FileInputStream

import scala.collection.JavaConversions._


object Javap {
    def main(args: Array[String]) {
        for (filename <- args) {
            println("Parsing: " + filename)

            var file = new File(filename)
            var input = new DataInputStream(new FileInputStream(file))

            var classInfo = new ClassInfo()
            try {
                classInfo.deserialize(input)
            } finally {
                input.close()
            }

            println("Classfile " + file.getAbsolutePath())
            javap(classInfo)
        }
    }

    def javap(classInfo: ClassInfo) {
        println("  Minor version: " + classInfo.minorVersion())
        println("  Major version: " + classInfo.majorVersion())
        println("  Flags: " + classInfo.access.debugString())

        println("  Class name: " + classInfo._thisClass.debugString())
        if (classInfo._superClass == null) {
            println("  Super class: (NULL) ")
        } else {
            println("  Super name: " + classInfo._superClass.debugString())
        }
        if (classInfo._interfaces.isEmpty()) {
            println("  Interfaces: (NONE)")
        } else {
            println("  Interfaces:")
            for (iface <- classInfo._interfaces) {
                println(iface.debugString())
            }
        }

        println("\nAttributes:")
        println(classInfo.attributes().debugString("  "))

        println("Constants:")
        for (info <- classInfo.constants()._tmpConstInfosByIndex.values()) {
            println(info.debugString())
        }

        println("\nFields:")
        for (field <- classInfo.fields()) {
            println("  " + field.name())
            println("    Type: " + field.fieldTypeString())
            println("    Flags: " + field.access().debugString())
            println("    Attributes:")
            println(field.attributes().debugString("      "))
        }

        println("Methods:")
        for (method <- classInfo.methods()) {
            println("  " + method.name())
            println("    Type: " + method.methodTypeString())
            println("    Flags: " + method.access().debugString())
            println("    Attributes:")
            println(method.attributes().debugString("      "))
        }
    }
}
