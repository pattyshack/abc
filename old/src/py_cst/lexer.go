package py_cst

import (
	"fmt"
	"io"
)

const eof = 0

// A "look ahead" lexer which handles line continuation, transforms ';' to
// newlines to simply parsing, and groups comments.
type Lexer struct {
	tokens []*Token

	tokenizeError bool

	endLocation Location

	prevLocation *Location
}

// NOTE: lex errors are return immediately during initialization.
func NewLexer(fileName string, reader io.Reader) (*Lexer, error) {
	lexer := &Lexer{}

	tokenizer := newTokenizer(fileName, reader)

	for {
		token, err := tokenizer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		lexer.tokens = append(lexer.tokens, token)
	}

	lexer.endLocation = tokenizer.CurrentLocation()
	lexer.transformLineContinuations()

	// Sadly, we can't transform semi-colon without changing semantic meaning
	// (due to "suite" reliance on "simple_stmt").

	lexer.transformComments()
	lexer.transformNewLines()

	return lexer, nil
}

func (l *Lexer) transformLineContinuations() {
	result := make([]*Token, 0, len(l.tokens))

	for _, token := range l.tokens {
		if token.TokenType != LINE_CONTINUATION {
			result = append(result, token)
			continue
		}

		if len(result) == 0 {
			// no-op line continuation at the the beginning of the file.
			continue
		}

		prev := result[len(result)-1]
		if prev.TokenType == COMMENT_NEWLINE ||
			prev.TokenType == NEWLINE ||
			prev.TokenType == SEMI_COLON {

			// no-op line continuation
			continue
		}

		prev.TrailingLineContinuation = true
	}

	l.tokens = result
}

func pruneEmpty(groups [][]string) [][]string {
	if len(groups) == 0 {
		return nil
	}

	if len(groups[len(groups)-1]) == 0 {
		return groups[:len(groups)-1]
	}
	return groups
}

func (l *Lexer) transformComments() {
	result := make([]*Token, 0, len(l.tokens))

	var prevNonNewLine *Token
	var leadingCommentGroups [][]string

	for _, token := range l.tokens {
		if token.TokenType != NEWLINE &&
			token.TokenType != COMMENT_NEWLINE {

			result = append(result, token)

			if token.TokenType == INDENT || token.TokenType == DEDENT {
				prevNonNewLine = nil
			} else {
				token.LeadingCommentGroups = pruneEmpty(leadingCommentGroups)
				leadingCommentGroups = nil

				prevNonNewLine = token
			}
			continue
		}

		if token.TokenType == COMMENT_NEWLINE {
			if prevNonNewLine != nil {
				if prevNonNewLine.TrailingCommentGroups == nil {
					prevNonNewLine.TrailingCommentGroups = [][]string{
						[]string{},
					}
				}
				prevNonNewLine.TrailingCommentGroups[0] = append(
					prevNonNewLine.TrailingCommentGroups[0],
					token.Value)

				token.TokenType = NEWLINE
				result = append(result, token)
			} else {
				if len(leadingCommentGroups) == 0 {
					leadingCommentGroups = [][]string{[]string{}}
				}

				last := len(leadingCommentGroups) - 1
				leadingCommentGroups[last] = append(
					leadingCommentGroups[last],
					token.Value)
			}
		} else {
			if prevNonNewLine != nil {
				prevNonNewLine = nil
			} else {
				if len(leadingCommentGroups) > 0 {
					last := len(leadingCommentGroups) - 1
					if len(leadingCommentGroups[last]) > 0 {
						leadingCommentGroups = append(
							leadingCommentGroups,
							[]string{})
					}
				}
			}

			result = append(result, token)
		}
	}

	if len(leadingCommentGroups) > 0 {
		// the parser will conver this into an implicit pass statement
		result = append(
			result,
			&Token{
				Node{
					Location: l.endLocation,
					CommentsFormatting: CommentsFormatting{
						LeadingCommentGroups: leadingCommentGroups,
					},
				},
				NEWLINE,
				"",
			})
	}

	l.tokens = result
}

func (l *Lexer) transformNewLines() {
	openness := 0
	result := make([]*Token, 0, len(l.tokens))

	var prev *Token
	for _, token := range l.tokens {
		switch token.TokenType {
		case LEFT_PARENTHESIS, LEFT_BRACE, LEFT_BRACKET:
			openness += 1
		case RIGHT_PARENTHESIS, RIGHT_BRACE, RIGHT_BRACKET:
			openness -= 1
		}

		if token.TokenType != NEWLINE {
			result = append(result, token)
			prev = token
		} else if openness > 0 {
			if prev.TokenType == NEWLINE {
				panic("This should never happen")
			}
			prev.TrailingNewLine = true
			prev.MergeFrom(&token.Node)
		} else {
			if prev == nil {
				continue
			}
			if prev.TokenType == NEWLINE {
				prev.MergeFrom(&token.Node)
				continue
			}

			result = append(result, token)
			prev = token
		}
	}

	// Pad the last statement with newline since the parser grammar expects it.
	if openness == 0 && prev != nil && prev.TokenType != NEWLINE {
		result = append(
			result,
			&Token{Node{Location: l.endLocation}, NEWLINE, ""})
	}

	l.tokens = result
}

func (l *Lexer) ToError(msg string) error {
	loc := l.endLocation
	if l.prevLocation != nil {
		loc = *l.prevLocation
	}

	return fmt.Errorf("parse error: %s (%v)", msg, loc)
}

func (l *Lexer) Lex(lval *yySymType) int {
	*lval = yySymType{} // clear value in case it's accidentally reused.

	if l.tokenizeError {
		return yyErrCode
	}

	if len(l.tokens) == 0 {
		return eof
	}

	lval.token = l.tokens[0]
	l.tokens = l.tokens[1:]

	l.prevLocation = &lval.token.Location

	return lval.token.TokenType
}

func (l *Lexer) PrintTokens() {
	for _, t := range l.tokens {
		fmt.Println(t)
	}
}
