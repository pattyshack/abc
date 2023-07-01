package yacc

import (
	"fmt"
	"io"

	"github.com/pattyshack/abc/src/lr/internal/parser"
)

type ParseContext struct {
	parser.Reducer

	Tokens []parser.LRToken

	*parser.Grammar
	Err error

	currPos parser.LRLocation
}

func newParseContext(filename string, reader io.Reader) (*ParseContext, error) {
	lexer := parser.NewLexer(filename, reader)

	tokens := []parser.LRToken{}
	for {
		token, err := lexer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return &ParseContext{Tokens: tokens}, nil
}

func (pc *ParseContext) SetDefinitions(defs []parser.Definition) {
	pc.Definitions = defs
}

func (pc *ParseContext) Error(errStr string) {
	pc.Err = fmt.Errorf(errStr + "(" + pc.currPos.String() + ")")
}

func (pc *ParseContext) Lex(val *LrSymType) int {
	if len(pc.Tokens) == 0 {
		return 0 // eof
	}

	tok := pc.Tokens[0]
	pc.Tokens = pc.Tokens[1:]
	switch t := tok.(type) {
	case *parser.Token:
		val.Token = t
	case *parser.LRGenericSymbol:
		val.Generic_ = t
	}

	pc.currPos = tok.Loc()

	return LRSymbolTypeToYaccTokenNum(tok.Id())
}

func Parse(
	filename string,
	reader io.Reader) (
	*parser.Grammar,
	error) {

	ctx, err := newParseContext(filename, reader)
	if err != nil {
		return nil, err
	}

	LrParse(ctx)

	if ctx.Err != nil {
		return nil, ctx.Err
	}

	return ctx.Grammar, nil
}
