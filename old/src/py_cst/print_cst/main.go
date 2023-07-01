package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pattyshack/abc/src/py_cst"
)

func main() {
	printTokens := flag.Bool("print-tokens", false, "print tokens")
	printParsedTree := flag.Bool("print-parsed-tree", true, "print parsed tree")
	flag.Parse()

	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReaderSize(file, 1024*1024)
		bytes, _ := reader.Peek(1024 * 1024)

		idx := strings.Index(string(bytes), "pyxl")
		if idx >= 0 {
			fmt.Println("Skipping pyxl:", filename)
			continue
		}

		idx = strings.Index(string(bytes), "__future__")
		if idx >= 0 {
			idx = strings.Index(string(bytes), "print_function")
			if idx >= 0 {
				fmt.Println("Skipping python3 print:", filename)
				continue
			}
		}

		fmt.Println("Parsing:", filename)
		ctx, err := py_cst.NewContext(filename, reader)
		if err != nil {
			fmt.Println("FAILED:", err)
		} else {
			if *printTokens {
				ctx.PrintTokens()
			}

			py_cst.Parse(ctx)
			if ctx.ParseError != nil {
				fmt.Println("FAILED:", ctx.ParseError)
			} else {
				if *printParsedTree {
					for _, s := range ctx.Statements {
						fmt.Println(s)
					}
				}

				fmt.Println()
				fmt.Println("OK", filename)
			}
		}

		fmt.Println("=============================================")
	}
}
