package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/pattyshack/abc/src/lr/parseutil"
)

const (
	textPeekSize = 1024
)

var (
	directiveRegexp = regexp.MustCompile(`\[\[|\$`)

	directiveSymbols = parseutil.Symbols{
		{"$", int(CopySectionToken)},
		{"#", int(CommentToken)},
	}

	// NOTE: for, $, and # directives can be both parameterless and parametered

	parameterlessOnlyDirectives = map[string]struct{}{
		"end":      struct{}{},
		"default":  struct{}{},
		"else":     struct{}{},
		"continue": struct{}{},
		"break":    struct{}{},
		"return":   struct{}{},
	}

	parameteredOnlyDirectives = map[string]struct{}{
		"switch":  struct{}{},
		"case":    struct{}{},
		"if":      struct{}{},
		"else if": struct{}{},
		"error":   struct{}{},
		"embed":   struct{}{},
	}
)

type BodyToken interface {
	Token

	// When true, and the previous statement is text, remove the whitespaces
	// in the text that are adjacent to this statement, potentially up to and
	// including the previous line's newline character.
	TrimLeadingWhitespaces() bool

	// When true, and the next statement is text, remove the whitespaces
	// in the text taht are adjacent to this statement, potentially up to and
	// including the current line's newline character.
	TrimTrailingWhitespaces() bool
}

type rawBodyLexer struct {
	reader *parseutil.LocationReader
}

func (lexer *rawBodyLexer) CurrentLocation() Location {
	return Location(lexer.reader.Location)
}

func (lexer *rawBodyLexer) Next() (BodyToken, error) {
	token, err := lexer.maybeTokenizeText()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeDirective()
	if token != nil || err != nil {
		return token, err
	}

	return nil, fmt.Errorf("Unexpected character at %s", lexer.reader.Location)
}

func (lexer *rawBodyLexer) maybeTokenizeText() (BodyToken, error) {
	loc := lexer.reader.Location
	buffer := bytes.NewBuffer(nil)

	foundEnd := false
	for !foundEnd {
		bytes, err := lexer.reader.Peek(textPeekSize)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF && len(bytes) == 0 {
			if buffer.Len() == 0 {
				return nil, io.EOF
			}

			break
		}

		index := directiveRegexp.FindIndex(bytes)

		endIdx := 0
		if index == nil {
			// The whole chunk does not contain any directive

			endIdx = len(bytes)
			if bytes[endIdx-1] == '[' && len(bytes) == textPeekSize {
				// NOTE: There's a chance that the last character in this
				// chunk belongs to a different token.  Don't consume the
				// '[' and read more bytes to verify.
				endIdx -= 1
			}
		} else {
			endIdx = index[0]

			if bytes[endIdx] == '[' { // found directive / end of text
				if bytes[endIdx+1] != '[' { // sanity check
					panic("Programming error")
				}

				foundEnd = true
			} else if bytes[endIdx] == '$' {
				if endIdx+1 < len(bytes)-1 {
					if bytes[endIdx+1] == '$' { // escaped $
						endIdx += 2
					} else {
						foundEnd = true
					}
				} else {
					if len(bytes) == textPeekSize {
						// NOTE: There's a chance that the last character
						// in // this chunk is an escaped '$'.  Don't
						// consume the '$' and read more bytes to verify.
					} else {
						// last
						foundEnd = true
					}
				}

			} else {
				panic("Programming error")
			}
		}

		if endIdx > 0 {
			_, err = buffer.Write(bytes[:endIdx])
			if err != nil {
				return nil, err
			}

			_, err = lexer.reader.Read(bytes[:endIdx])
			if err != nil {
				panic(err) // should never happen
			}
		}
	}

	if buffer.Len() == 0 {
		return nil, nil
	}

	value := string(buffer.Bytes())

	// unescape $
	value = strings.ReplaceAll(value, "$$", "$")

	// "escape" ` since the text block is output as raw string literal.
	// Note that there's no clean way to escape ` in raw string.
	// https://github.com/golang/go/issues/24475
	value = strings.ReplaceAll(value, "`", "`+\"`\"+`")

	return NewAtom(TextToken, Location(loc), value, false, false), nil
}

func (lexer *rawBodyLexer) tokenizeNonSubstituteDirective() (BodyToken, error) {
	content, loc, err := readDirective(lexer.reader, 2, "]]")
	if err != nil {
		return nil, err
	}
	content = content[2 : len(content)-2]

	trimLeading := false
	if len(content) > 0 && content[0] == '-' {
		trimLeading = true
		content = content[1:]
	}

	trimTrailing := false
	if len(content) > 0 && content[len(content)-1] == '-' {
		trimTrailing = true
		content = content[:len(content)-1]
	}

	directiveReader := parseutil.NewLocationReader(
		"",
		bytes.NewBuffer(content))
	directiveReader.Location = parseutil.Location(loc)

	err = parseutil.StripLeadingWhitespaces(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	symbol, _, err := parseutil.MaybeTokenizeSymbol(
		directiveReader,
		directiveSymbols)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if symbol != nil {
		content, err = ioutil.ReadAll(directiveReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		return NewAtom(
			SymbolId(symbol.Id),
			loc,
			string(content),
			trimLeading,
			trimTrailing), nil
	}

	id, _, err := parseutil.MaybeTokenizeIdentifier(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	err = parseutil.StripLeadingWhitespaces(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	param := ""
	// check for "else if" compound identifier
	if id == "else" {
		second, _, err := parseutil.MaybeTokenizeIdentifier(directiveReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if second == "if" {
			id = "else if"

			err = parseutil.StripLeadingWhitespaces(directiveReader)
			if err != nil && err != io.EOF {
				return nil, err
			}
		} else {
			param = string(second)
		}
	}

	content, err = ioutil.ReadAll(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	param += string(content)

	_, ok := parameterlessOnlyDirectives[id]
	if ok && len(param) > 0 {
		return nil, fmt.Errorf(
			"unexpected parameter specified in [[%s]] directive (%s)",
			id,
			loc)
	}

	_, ok = parameteredOnlyDirectives[id]
	if ok && param == "" {
		return nil, fmt.Errorf(
			"expected parameter not specified in [[%s]] directive (%s)",
			id,
			loc)
	}

	switch id {
	case "":
		return nil, fmt.Errorf(
			"invalid directive. directive type not specified (%s)", loc)
	case "end":
		return NewTToken(EndToken, loc, trimLeading, trimTrailing), nil
	case "default":
		return NewTToken(DefaultToken, loc, trimLeading, trimTrailing), nil
	case "else":
		return NewTToken(ElseToken, loc, trimLeading, trimTrailing), nil
	case "for":
		return NewValue(ForToken, loc, param, trimLeading, trimTrailing), nil
	case "switch":
		return NewValue(SwitchToken, loc, param, trimLeading, trimTrailing), nil
	case "case":
		return NewValue(CaseToken, loc, param, trimLeading, trimTrailing), nil
	case "if":
		return NewValue(IfToken, loc, param, trimLeading, trimTrailing), nil
	case "else if":
		return NewValue(ElseIfToken, loc, param, trimLeading, trimTrailing), nil
	case "continue":
		return NewAtom(
			ContinueToken,
			loc,
			param,
			trimLeading,
			trimTrailing), nil
	case "break":
		return NewAtom(BreakToken, loc, param, trimLeading, trimTrailing), nil
	case "return":
		return NewAtom(ReturnToken, loc, param, trimLeading, trimTrailing), nil
	case "error":
		return NewAtom(ErrorToken, loc, param, trimLeading, trimTrailing), nil
	case "embed":
		return NewAtom(EmbedToken, loc, param, trimLeading, trimTrailing), nil
	}

	return nil, fmt.Errorf(
		"invalid directive. unknown directive type %s (%s)",
		id,
		loc)
}

func (lexer *rawBodyLexer) maybeTokenizeDirective() (BodyToken, error) {
	content, err := lexer.reader.Peek(2)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if len(content) < 2 {
		return nil, nil
	}

	if content[0] == '[' {
		if content[1] != '[' {
			panic("Programming error")
		}

		return lexer.tokenizeNonSubstituteDirective()

	} else if content[0] == '$' {
		if 'a' <= content[1] && content[1] <= 'z' ||
			'A' <= content[1] && content[1] <= 'Z' ||
			content[1] == '_' {
			// $<identifier>

			loc := lexer.reader.Location

			_, err := lexer.reader.ReadByte()
			if err != nil {
				panic(err) // Should never happen
			}

			value, _, err := parseutil.MaybeTokenizeIdentifier(lexer.reader)
			if err != nil {
				return nil, err
			}

			return NewAtom(
				SubstitutionToken,
				Location(loc),
				string(value),
				false,
				false), nil

		} else if content[1] == '(' {

			content, loc, err := readDirective(lexer.reader, 2, ")")
			if err != nil {
				return nil, err
			}

			value := string(content[2 : len(content)-1])
			return NewAtom(SubstitutionToken, loc, value, false, false), nil

		} else if content[1] == '$' {
			panic("Programming error")
		} else {
            return nil, fmt.Errorf(
                "invalid substitute directive (%s)",
			    lexer.reader.Location)
        }
	}

	panic("Programming error")
}

// Search for the first occurence of terminal starting from startIdx.  The
// terminal must start with one of the following characters: ']' ')' '}' '\n'.
// This respect golang scoping, string, char and comments, i.e.,
// {} [] () `` "" '' /**/ //
func readDirective(
	reader *parseutil.LocationReader,
	startIdx int,
	terminal string) (
	[]byte,
	Location,
	error) {

	if terminal == "" {
		return nil, Location{}, fmt.Errorf("Invalid terminal")
	}

	if terminal[0] != '}' &&
		terminal[0] != ')' &&
		terminal[0] != ']' &&
		terminal[0] != '\n' {
		return nil, Location{}, fmt.Errorf("Invalid terminal: " + terminal)
	}

	currentScope := []string{"#root"}

	peekRange := 64
	prevLen := 0
	checkIdx := startIdx

	bytes, err := reader.Peek(peekRange)
	if err != nil && err != io.EOF {
		return nil, Location{}, err
	}

	if len(bytes) <= startIdx {
		return nil, Location{}, fmt.Errorf(
			"lex error: \"%s\" not found (%s)",
			terminal,
			reader.Location)
	}

	for {
		needMoreBytes := false
		for !needMoreBytes && checkIdx < len(bytes) {
			char := bytes[checkIdx]

			top := currentScope[len(currentScope)-1]
			switch top {
			case "#root", "{", "[", "(": // nestable scopes
				switch char {
				case '{', '[', '(', '`', '"', '\'':
					currentScope = append(currentScope, string(char))
				case '/':
					if len(bytes) <= checkIdx+1 {
						needMoreBytes = true
						continue
					} else if bytes[checkIdx+1] == '*' ||
						bytes[checkIdx+1] == '/' {

						currentScope = append(currentScope, string(bytes[checkIdx:checkIdx+2]))
						checkIdx += 1
					}
				case '}', ']', ')', '\n':
					if (top == "(" && char == ')') ||
						(top == "{" && char == '}') ||
						(top == "[" && char == ']') {

						currentScope = currentScope[:len(currentScope)-1]
						checkIdx += 1
						continue
					}

					if top == "#root" {
						if len(bytes) <= checkIdx+len(terminal)-1 {
							needMoreBytes = true
							continue
						}

						if terminal == string(bytes[checkIdx:checkIdx+len(terminal)]) {
							loc := Location(reader.Location)

							bytes = make([]byte, checkIdx+len(terminal))
							n, err := reader.Read(bytes)
							if len(bytes) != n || err != nil {
								panic(err) // should never happen
							}

							return bytes, loc, nil
						}
					}

					if char != '\n' {
						bytes = bytes[:checkIdx]

						n, err := reader.Read(bytes)
						if len(bytes) != n || err != nil {
							panic(err) // should never happen
						}

						return nil, Location{}, fmt.Errorf(
							"lex error: no matching pair for %c (%s)",
							char,
							reader.Location)
					}
				}
			case "//":
				if char == '\n' {
					currentScope = currentScope[:len(currentScope)-1]
					// don't include '\n' since it could be part of the terminal
					checkIdx -= 1
				}
			case "/*":
				if char == '*' {
					if len(bytes) <= checkIdx+1 {
						needMoreBytes = true
						continue
					} else if bytes[checkIdx+1] == '/' {
						checkIdx += 1
						currentScope = currentScope[:len(currentScope)-1]
					}
				}
			case "`", "'", "\"":
				if char == '\\' {
					checkIdx += 1 // skip escaped char
				} else if char == top[0] {
					currentScope = currentScope[:len(currentScope)-1]
				}
			default:
				panic(fmt.Sprintf("Programming error: %v", currentScope))
			}

			if !needMoreBytes {
				checkIdx += 1
			}
		}

		peekRange *= 2
		prevLen = len(bytes)

		bytes, err = reader.Peek(peekRange)
		if err != nil && err != io.EOF {
			return nil, Location{}, err
		}

		if prevLen == len(bytes) { // not found
			return nil, Location{}, fmt.Errorf(
				"lex error: \"%s\" not found (%s)",
				terminal,
				reader.Location)
		}
	}
}
