package main

import "fmt"

func printTree(name string, tokens ...Token) {
	fmt.Println(name, "==========")
	node, remaining := parse(tokens)
	fmt.Println("Tree:")
	fmt.Print(formatNode(node, 1))
	fmt.Println("Remaining:")
	fmt.Println(remaining)
}

func main() {
	printTree(
		"empty",
		lbrace(),
		rbrace(),
	)

	printTree(
		"basic",
		lbrace(),

		id("a"),

		id("b"),
		plus(),
		id("c"),

		id("d"),
		minus(),
		id("e"),

		&Block{
			lbrace(),
			nil,
			rbrace(),
		},

		&token{"bad", ERROR, "ERROR: bad"},

		rbrace(),
	)

	printTree(
		"manual block reduction",
		lbrace(),

		id("a"),

		lbrace(),
		id("f"),

		lbrace(),
		id("g"),
		rbrace(),

		lbrace(),
		id("h"),
		minus(),
		rbrace(),

		rbrace(),

		lbrace(),
		plus(),
		rbrace(),

		rbrace(),
	)
}
