// HORRIBLE HACK!!! go yacc does not support %parser-param.  We'll pass in the
// context via yylex instead ...

package py_cst

import (
	"io"
)

func Parse(ctx *Context) int {
	return yyParse(ctx)
}

type Context struct {
	lexer *Lexer

	Statements []Statement

	ParseError error
}

// NOTE: lex errors are returned immediately.
func NewContext(fileName string, reader io.Reader) (*Context, error) {
	lexer, err := NewLexer(fileName, reader)
	if err != nil {
		return nil, err
	}

	return &Context{
		lexer: lexer,
	}, nil
}

func (c *Context) Error(msg string) {
	c.ParseError = c.lexer.ToError(msg)
}

// hand rolling a lexer =...(
func (c *Context) Lex(lval *yySymType) int {
	return c.lexer.Lex(lval)
}

func (c *Context) PrintTokens() {
	c.lexer.PrintTokens()
}

func init() {
	yyErrorVerbose = true
}
