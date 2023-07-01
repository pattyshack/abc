package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	template "github.com/pattyshack/abc/src/template/internal"
	"github.com/pattyshack/abc/src/template/internal/manual_codegen"
	"github.com/pattyshack/abc/src/template/internal/templated_codegen"
)

func main() {
	useManualCodeGen := flag.Bool(
		"use-manual-code-gen",
		false,
		"Use hand written code generator instead template based code " +
        "generator, primarily used for bootstrapping.")

	shouldPrintGenerated := flag.Bool(
		"print-generated",
		false,
		"For testing only")

	output := flag.String("o", "", "template code gen output")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Printf("Usage of %s:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		return
	}

	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lexer, err := template.NewLexer(filename, file)
	if err != nil {
		panic(err)
	}

	templateFile, err := template.Parse(lexer, template.ReducerImpl{})
	if err != nil {
		panic(err)
	}

	var codeGenTemplate io.WriterTo
	if *useManualCodeGen {
		fmt.Println("Using manual code generator")
		codeGenTemplate = manual_codegen.NewTemplate(filename, templateFile)
	} else {
		codeGenTemplate = templated_codegen.NewTemplate(filename, templateFile)
	}

	if *output != "" {
		outputFile, err := os.Create(*output)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		_, err = codeGenTemplate.WriteTo(outputFile)
		if err != nil {
			panic(err)
		}
	}

	if *shouldPrintGenerated {
		fmt.Println("File (Generated):", filename)
		fmt.Println("====================================")

		_, err = codeGenTemplate.WriteTo(os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}
