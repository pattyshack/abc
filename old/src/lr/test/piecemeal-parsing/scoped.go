package main

import (
	"fmt"
	"io"
)

type LookAheadLexer struct {
	base Lexer

	next Token
	err  error
}

func (l *LookAheadLexer) Next() (Token, error) {
	if l.next != nil || l.err != nil {
		next := l.next
		err := l.err
		l.next = nil
		l.err = nil
		return next, err
	}

	return l.base.Next()
}

func (l *LookAheadLexer) Peek() (Token, error) {
	if l.next == nil && l.err == nil {
		l.next, l.err = l.base.Next()
	}
	return l.next, l.err
}

type PrependedLexer struct {
	tokens []Token
	lexer  Lexer
}

func NewPrependedLexer(lexer Lexer, tokens ...Token) *PrependedLexer {
	return &PrependedLexer{tokens, lexer}
}

func (l *PrependedLexer) Next() (Token, error) {
	if len(l.tokens) == 0 {
		return l.lexer.Next()
	}

	head := l.tokens[0]
	l.tokens = l.tokens[1:]
	return head, nil
}

func (PrependedLexer) CurrentLocation() Location { return Location{} }

type ScopedLexer struct {
	base *LookAheadLexer

	depth int

	scopeEnded bool
}

func NewScopedLexer(input string) *ScopedLexer {
	return &ScopedLexer{
		&LookAheadLexer{NewBasicLexer(input), nil, nil},
		0,
		false}
}

func (l *ScopedLexer) Next() (Token, error) {
	if l.scopeEnded {
		l.scopeEnded = false
		return nil, io.EOF
	}

	token, err := l.base.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() != '{' && token.Id() != '}' {
		return token, nil
	}

	if token.Id() == '}' {
		if l.depth > 0 {
			l.depth -= 1
			l.scopeEnded = true
			return token, nil
		} else {
			return &Err{
				fmt.Errorf("Unexpected RBRACE.  No matching LBRACE"),
			}, nil
		}
	}

	// lbrace
	blockDepth := l.depth
	l.depth += 1

	block, parseErr := ParseBlock(NewPrependedLexer(l, token), ReducerImpl{})
	if parseErr == nil {
		return &Symbol{SymbolId_: BlockType, Block: block}, nil
	}

	// Drop the fake end marker in case the syntax error was caused by an
	// unexpected rbrace
	l.scopeEnded = false

	for l.depth > blockDepth {
		next, err := l.base.Peek()
		if err != nil {
			return &Err{parseErr}, nil
		}

		if next.Id() == '{' {
			l.depth += 1
		} else if next.Id() == '}' {
			l.depth -= 1
		}

		_, _ = l.base.Next()
	}

	return &Err{parseErr}, nil
}

func (ScopedLexer) CurrentLocation() Location { return Location{} }
