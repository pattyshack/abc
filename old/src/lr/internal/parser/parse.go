package parser

import (
	"io"
)

func Parse(filename string, reader io.Reader) (*Grammar, error) {

	return LRParse(NewLexer(filename, reader), Reducer{})
}
