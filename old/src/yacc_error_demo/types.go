package main

type Node interface {
	Children() []Node
	String() string
}

type Token interface {
	Node

	Type() int
}

type parseContext struct {
	tokens []Token

	result Node
	errStr string
}

func (ctx *parseContext) Lex(val *demoSymType) int {
	if len(ctx.tokens) == 0 {
		return 0 // eof
	}

	tok := ctx.tokens[0]
	ctx.tokens = ctx.tokens[1:]
	val.Node = tok
	return tok.Type()
}

func (ctx *parseContext) Error(s string) {
	ctx.errStr = s
}

type token struct {
	val   string
	ttype int
	name  string
}

func id(s string) *token {
	return &token{s, ID, "ID=" + s}
}

func plus() *token {
	return &token{"+", PLUS, "+"}
}

func minus() *token {
	return &token{"-", MINUS, "-"}
}

func lbrace() *token {
	return &token{"{", LBRACE, "{"}
}

func rbrace() *token {
	return &token{"}", RBRACE, "}"}
}

func (token) Children() []Node {
	return nil
}

func (t *token) String() string {
	return t.name
}

func (t *token) Type() int {
	return t.ttype
}

type Binary struct {
	left  Node
	op    Node
	right Node
}

func (bin *Binary) Children() []Node {
	return []Node{bin.left, bin.op, bin.right}
}

func (Binary) String() string {
	return "Binary"
}

type Block struct {
	lbrace Node
	stmts  []Node
	rbrace Node
}

func (block *Block) Children() []Node {
	children := append([]Node{block.lbrace}, block.stmts...)
	return append(children, block.rbrace)
}

func (Block) String() string {
	return "Block"
}

func (Block) Type() int {
	return BLOCK
}

func formatNode(node Node, indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "  "
	}

	if node == nil {
		return indent + "<nil>\n"
	}

	formatted := indent + node.String() + "\n"
	for _, child := range node.Children() {
		formatted += formatNode(child, indentLevel+1)
	}

	return formatted
}

func parse(tokens []Token) (Node, []Token) {
	if len(tokens) == 0 {
		return nil, nil
	}

	currentScope := []Token{tokens[0]}
	remaining := tokens[1:]
	for len(remaining) > 0 {
		tok := remaining[0]
		var node Node
		if tok.Type() == LBRACE {
			node, remaining = parse(remaining)

			currentScope = append(currentScope, node.(Token))
		} else {
			currentScope = append(currentScope, tok)
			remaining = remaining[1:]

			if tok.Type() == RBRACE {
				ctx := &parseContext{tokens: currentScope}
				demoParse(ctx)
				if ctx.errStr != "" {
					return &token{
						ctx.errStr,
						ERROR,
						"ERROR: " + ctx.errStr,
					}, remaining
				} else {
					return ctx.result, remaining
				}
			}
		}
	}

	// Imbalanced scoped.  parse anyways
	ctx := &parseContext{tokens: currentScope}
	demoParse(ctx)
	if ctx.errStr != "" {
		return &token{
			ctx.errStr,
			ERROR,
			"ERROR: " + ctx.errStr,
		}, nil
	} else {
		return ctx.result, nil
	}
}
