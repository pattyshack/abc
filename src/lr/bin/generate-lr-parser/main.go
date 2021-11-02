package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen"
	"github.com/pattyshack/abc/src/lr/internal/parser"
	"github.com/pattyshack/abc/src/lr/internal/parser/yacc"
)

func main() {
	useYacc := flag.Bool(
		"use-yacc",
		true,
		"Use yacc generated parser instead of bootstrap parser")

	shouldPrintTokens := flag.Bool("print-tokens", false, "For testing only")
	shouldPrintParsed := flag.Bool("print-parsed", false, "For testing only")
	shouldPrintLRStates := flag.Bool("print-lr-states", false, "For testing only")
	shouldPrintGenerated := flag.Bool("print-generated", false, "For testing only")

	language := flag.String("language", "go", "output/target language")
	output := flag.String("o", "", "parser output")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Printf("Usage of %s:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		return
	}

	filename := flag.Args()[0]

	if *shouldPrintTokens {
		printTokens(filename)
	}

	if *shouldPrintParsed {
		printParsed(filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var parsed *parser.Grammar
	if *useYacc {
		fmt.Println("Using yacc generated parser")
		parsed, err = yacc.Parse(filename, file)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Using self generated parser")
		parsed, err = parser.Parse(filename, file)
		if err != nil {
			panic(err)
		}

	}

	grammar, err := lr.NewGrammar(filename, parsed)
	if err != nil {
		panic(err)
	}

	lrStates := lr.NewLRStates(grammar)

	if *shouldPrintLRStates {
		printLRStates(lrStates)
	}

	if *shouldPrintGenerated || *output != "" {
		builder, err := code_gen.GenerateLRCode(grammar, lrStates, *language)
		if err != nil {
			panic(err)
		}

		buffer := bytes.NewBuffer(nil)
		_, err = builder.WriteTo(buffer)

		if err != nil {
			panic(err)
		}

		bytes := buffer.Bytes()

		if *shouldPrintGenerated {
			fmt.Println("File (Generated): ", filename)
			fmt.Println("==================================")

			fmt.Println(string(bytes))
		}

		if *output != "" {
			err = ioutil.WriteFile(*output, bytes, 0664)
			if err != nil {
				panic(err)
			}
		}
	}
}

func printTokens(filename string) {
	fmt.Println("File (Tokens): ", filename)
	fmt.Println("==================================")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lexer := parser.NewLexer(filename, file)

	for {
		token, err := lexer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Println(token)
	}
}

func printParsed(filename string) {
	fmt.Println("File (Parsed Definitions): ", filename)
	fmt.Println("==================================")

	yacc.LrErrorVerbose = true

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parsed, err := yacc.Parse(filename, file)
	if err != nil {
		panic(err)
	}

	for _, def := range parsed.Definitions {
		fmt.Println(def.Location())
		fmt.Println(def)
		fmt.Println()
	}

	for _, section := range parsed.AdditionalSections {
		fmt.Println("--------------------------")
		fmt.Println(section.Name.Location)
		fmt.Println("Additional Section: ", section.Name.Value)
		fmt.Println(section.Content.Value)
	}
}

func printLRStates(states *lr.LRStates) {
	fmt.Println("File (LR States): ", states.Source)
	fmt.Println("==================================")

	symbols := []string{"$", "*"}
	for _, terms := range [][]*lr.Term{states.Terminals, states.NonTerminals} {
		for _, term := range terms {
			symbols = append(symbols, term.Name)
		}
	}

	gotoCount := 0
	reduceCount := 0

	fmt.Println("States:")
	for _, state := range states.OrderedStates {
		reduce := map[string][]string{}
		fmt.Println("    State", state.StateNum, ":")
		fmt.Println("      Kernel Items:")
		firstNonKernel := true
		for _, item := range state.Items {
			if len(item.Parsed) == 0 && firstNonKernel {
				firstNonKernel = false
				fmt.Println("      Non-kernel Items:")
			}

			if len(item.Expected) == 0 {
				reduceCount += 1
				reduce[item.LookAhead] = append(
					reduce[item.LookAhead],
					item.Name)
			}
			fmt.Println("       ", item)
		}
		fmt.Println("      Reduce:")
		if len(reduce) == 0 {
			fmt.Println("        (nil)")
		}
		for _, symbol := range symbols {
			list := reduce[symbol]
			if len(list) > 0 {
				fmt.Printf("        %s -> %v\n", symbol, list)
			}
		}
		fmt.Println("      Goto:")
		gotoCount += len(state.Goto)
		for _, symbol := range symbols {
			child, ok := state.Goto[symbol]
			if ok {
				fmt.Printf("        %s -> State %d\n", symbol, child.StateNum)
			}
		}

		fmt.Println()
	}

	fmt.Println("Number of shift actions:", gotoCount)
	fmt.Println("Number of reduce actions:", reduceCount)
}