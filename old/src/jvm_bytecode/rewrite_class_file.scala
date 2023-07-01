import java.io.DataInputStream
import java.io.DataOutputStream
import java.io.File
import java.io.FileInputStream
import java.io.FileOutputStream

import scala.collection.JavaConversions._


object RewriteClassFile {
    def main(args: Array[String]) {
        for (filename <- args) {
            println("rewriting class file: " + filename)
            rewriteClass(filename)
        }
    }

    def rewriteClass(filename: String) {
        var input = new DataInputStream(new FileInputStream(filename))

        var classInfo = new ClassInfo()
        try {
            classInfo.deserialize(input)
        } finally {
            input.close()
        }

        classInfo.analyze()

        var output = new DataOutputStream(new FileOutputStream(filename))
        try {
            classInfo.serialize(output)
        } finally {
            output.close()
        }

        //Javap.javap(classInfo)
    }
}
