package main

import (
	"io"
	"strings"
)

type Expr interface {
	String() string
	str(i int) string
}

func indent(i int) string {
	result := ""
	for i > 0 {
		i -= 1
		result += "    "
	}
	return result
}

type Id struct {
	GenericSymbol

	Value string
}

func (id *Id) String() string {
	return id.str(0)
}

func (id *Id) str(i int) string {
	return indent(i) + "ID=" + id.Value
}

type Err struct {
	Err error
}

func (Err) Id() SymbolId { return ErrorToken }

func (Err) Loc() Location { return Location{} }

func (e *Err) String() string {
	return e.str(0)
}

func (e *Err) str(i int) string {
	return indent(i) + "ERROR=" + e.Err.Error()
}

type Block struct {
	List []Expr
}

func (b *Block) String() string {
	return b.str(0)
}

func (b *Block) str(i int) string {
	result := indent(i) + "BLOCK={\n"
	for _, item := range b.List {
		result += item.str(i+1) + "\n"
	}
	result += indent(i) + "}"
	return result
}

type Binary struct {
	Left  Expr
	Op    *GenericSymbol
	Right Expr
}

func (b *Binary) String() string {
	return b.str(0)
}

func (b *Binary) str(i int) string {
	result := indent(i) + "BINARY=\n"
	result += b.Left.str(i+1) + "\n"
	result += indent(i+1) + "OP=" + string(b.Op.Id()) + "\n"
	result += b.Right.str(i + 1)
	return result
}

type BasicLexer struct {
	Tokens []Token
}

func NewBasicLexer(input string) *BasicLexer {
	tokens := []Token{}
	for _, str := range strings.Split(input, " ") {
		switch str {
		case "":
		case "{":
			tokens = append(tokens, &GenericSymbol{'{', Location{}})
		case "}":
			tokens = append(tokens, &GenericSymbol{'}', Location{}})
		case "+":
			tokens = append(tokens, &GenericSymbol{'+', Location{}})
		case "-":
			tokens = append(tokens, &GenericSymbol{'-', Location{}})
		default:
			tokens = append(
				tokens,
				&Id{GenericSymbol{IdToken, Location{}}, str})
		}
	}

	return &BasicLexer{tokens}
}

func (bl *BasicLexer) Next() (Token, error) {
	if len(bl.Tokens) == 0 {
		return nil, io.EOF
	}

	head := bl.Tokens[0]
	bl.Tokens = bl.Tokens[1:]
	return head, nil
}

func (BasicLexer) CurrentLocation() Location { return Location{} }

type ReducerImpl struct{}

func (ReducerImpl) ParseAllToRoot(
	marker *GenericSymbol,
	list []Expr) (
	interface{},
	error) {

	return list, nil
}

func (ReducerImpl) ParseBlockToRoot(
	marker *GenericSymbol,
	block *Block) (
	interface{},
	error) {

	return block, nil
}

func (ReducerImpl) AddToExprList(list []Expr, expr Expr) ([]Expr, error) {
	return append(list, expr), nil
}

func (ReducerImpl) NilToExprList() ([]Expr, error) {
	return []Expr{}, nil
}

func (ReducerImpl) IdToAtom(id *Id) (Expr, error) {
	return id, nil
}

func (ReducerImpl) ErrorToAtom(err *Err) (Expr, error) {
	return err, nil
}

func (ReducerImpl) BlockToAtom(block *Block) (Expr, error) {
	return block, nil
}

func (ReducerImpl) AtomToExpr(atom Expr) (Expr, error) {
	return atom, nil
}

func (ReducerImpl) BinaryToExpr(
	left Expr,
	op *GenericSymbol,
	right Expr) (
	Expr,
	error) {

	return &Binary{left, op, right}, nil
}

func (ReducerImpl) PlusToOp(plus *GenericSymbol) (*GenericSymbol, error) {
	return plus, nil
}

func (ReducerImpl) MinusToOp(minus *GenericSymbol) (*GenericSymbol, error) {
	return minus, nil
}

func (ReducerImpl) ToBlock(
	lbrace *GenericSymbol,
	list []Expr,
	rbrace *GenericSymbol) (
	*Block,
	error) {

	return &Block{list}, nil
}
