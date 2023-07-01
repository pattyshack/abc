package main

import (
	"fmt"
)

func printTree(input string) {
	fmt.Println("======")
	fmt.Println("Input:", input)
	fmt.Println("======")
	fmt.Println("Basic:")
	list, err := ParseExprList(NewBasicLexer(input), ReducerImpl{})
	if err != nil {
		fmt.Println("Failed to parse:", err)
	} else {
		for _, item := range list {
			fmt.Println(item)
		}
	}
	fmt.Println("------")
	fmt.Println("Scoped:")
	list, err = ParseExprList(NewScopedLexer(input), ReducerImpl{})
	if err != nil {
		fmt.Println("Failed to parse:", err)
	} else {
		for _, item := range list {
			fmt.Println(item)
		}
	}
}

func main() {
	// printTree("a + b")
	// printTree("a + b + c - d")
	// printTree("{ }")
	// printTree("a + { }")

	// printTree("a { }")
	printTree("{ } a b + c d { 1 2 3 4 + { 5 + 6 } + 7 }")
	printTree("{ } } a b + c d { 1 2 3 4 + { { a + + b }  + 6 } + 7 } foo { a + } blah {")
}
