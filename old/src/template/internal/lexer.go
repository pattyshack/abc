package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pattyshack/abc/src/lr/parseutil"
)

var (
	sectionMarker = parseutil.Symbols{
		{"%%\n", 0},
		{"%%", 0},
	}
)

type currentLexer interface {
	Next() (Token, error)
}

type LexerImpl struct {
	reader *parseutil.LocationReader

	currentLexer
}

func NewLexer(filename string, input io.Reader) (Lexer, error) {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	content = stripHeaderComments(content)

	reader := parseutil.NewLocationReader(filename, bytes.NewBuffer(content))

	return &LexerImpl{
		reader:       reader,
		currentLexer: &headerLexer{reader},
	}, nil
}

func (lexer *LexerImpl) CurrentLocation() Location {
	return Location(lexer.reader.Location)
}

func (lexer *LexerImpl) Next() (Token, error) {
	token, err := lexer.currentLexer.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() == SectionMarkerToken {
		lexer.currentLexer = &bodyLexer{
			raw: &rawBodyLexer{lexer.reader},
		}
	}

	return token, nil
}

type headerLexer struct {
	reader *parseutil.LocationReader
}

func (lexer *headerLexer) Next() (Token, error) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	val, loc, err := parseutil.MaybeTokenizeIdentifier(lexer.reader)
	if err != nil {
		return nil, err
	}

	switch string(val) {
	case "package":
		return lexer.tokenizePackage(Location(loc))
	case "import":
		return lexer.tokenizeImport(Location(loc))
	case "template":
		return lexer.tokenizeTemplateDecl(Location(loc))
	case "":
		// try to tokenize symbol below
	default:
		return nil, fmt.Errorf(
			"Unexpected IDENTIFIER %s (%s)",
			string(val),
			loc)
	}

	symbol, loc, err := parseutil.MaybeTokenizeSymbol(
		lexer.reader,
		sectionMarker)
	if err != nil {
		return nil, err
	}

	if symbol != nil {
		return &GenericSymbol{SectionMarkerToken, Location(loc)}, nil
	}

	return nil, fmt.Errorf("Unexpected character at %s", lexer.reader.Location)
}

func (lexer *headerLexer) tokenizePackage(pkgLoc Location) (Token, error) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	val, _, err := parseutil.MaybeTokenizeIdentifier(lexer.reader)
	if err != nil {
		return nil, err
	}

	if val != "" {
		return NewValue(PackageToken, pkgLoc, val, false, false), nil
	}

	return nil, fmt.Errorf("Unexpected character at %s", lexer.reader.Location)
}

func (lexer *headerLexer) tokenizeImport(importLoc Location) (Token, error) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	symbol, _, err := parseutil.MaybeTokenizeSymbol(
		lexer.reader,
		parseutil.Symbols{{"(", 0}})
	if err != nil {
		return nil, err
	}

	if symbol == nil {
		return nil, fmt.Errorf(
			"Unexpected character at %s",
			lexer.reader.Location)
	}

	value, _, err := readDirective(lexer.reader, 0, ")")
	if err != nil {
		return nil, err
	}

	value = value[:len(value)-1]

	return NewValue(ImportToken, importLoc, string(value), false, false), nil
}

func (lexer *headerLexer) tokenizeTemplateDecl(
	declLoc Location) (
	Token,
	error) {

	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	templateName, _, err := parseutil.MaybeTokenizeIdentifier(lexer.reader)
	if err != nil {
		return nil, err
	}

	if templateName == "" {
		return nil, fmt.Errorf(
			"Unexpected character at %s",
			lexer.reader.Location)
	}

	err = parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	lcurl, _, err := parseutil.MaybeTokenizeSymbol(
		lexer.reader,
		parseutil.Symbols{{"{", 0}})
	if err != nil {
		return nil, err
	}

	if lcurl == nil {
		return nil, fmt.Errorf(
			"Unexpected character at %s",
			lexer.reader.Location)
	}

	body, loc, err := readDirective(lexer.reader, 0, "}")
	if err != nil {
		return nil, err
	}

	body = body[:len(body)-1]

	buffer := bytes.NewBuffer(body)
	declReader := parseutil.NewLocationReader("", buffer)
	declReader.Location = parseutil.Location(loc)

	args := []Argument{}
	for {
		err := parseutil.StripLeadingWhitespaces(declReader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		head, err := declReader.Peek(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(head) == 0 {
			break
		}

		line, loc, err := readDirective(declReader, 0, "\n")
		if err != nil {
			return nil, err
		}

		line = line[:len(line)-1]
		lineReader := parseutil.NewLocationReader("", bytes.NewBuffer(line))
		lineReader.Location = parseutil.Location(loc)

		argName, _, err := parseutil.MaybeTokenizeIdentifier(lineReader)
		if err != nil {
			return nil, err
		}

		if argName == "" {
			return nil, fmt.Errorf(
				"Expecting argument name (%s)",
				lineReader.Location)
		}

		err = parseutil.StripLeadingWhitespaces(lineReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		typeName, err := ioutil.ReadAll(lineReader)
		if err != nil {
			return nil, err
		}

		if len(typeName) == 0 {
			return nil, fmt.Errorf(
				"Expecting argument type (%s)",
				lineReader.Location)
		}

		args = append(args, Argument{string(argName), string(typeName)})
	}

	return NewTemplateDeclaration(declLoc, templateName, args), nil
}

// Strip all comments from the header section of template.  Note that
// the content is modify in place.
func stripHeaderComments(template []byte) []byte {
	curr := 0
	shifted := 0

	skip := func() {
		curr += 1
	}

	shift := func() {
		template[shifted] = template[curr]
		shifted += 1
		curr += 1
	}

	// XXX maybe handle `` string and '' char as well
	inString := false
	inLineComment := false
	inBlockComment := false

	for curr < len(template) {
		char := template[curr]

		if inString {
			shift()

			if char == '\\' {
				if curr < len(template) {
					shift() // shift escaped character
				}
			} else if char == '"' {
				inString = false
			}
		} else if inLineComment {
			if char == '\n' { // preserve '\n' in the code
				inLineComment = false
			} else {
				skip()
			}
		} else if inBlockComment {
			if char == '\n' {
				// We need to preserve \n to ensure token locations are correct
				shift()
			} else {
				skip()
			}

			if char == '*' &&
				curr < len(template) &&
				template[curr] == '/' {

				skip()
				inBlockComment = false
			}
		} else {
			if char == '/' && curr+1 < len(template) {
				if template[curr+1] == '/' {
					skip()
					skip()
					inLineComment = true
				} else if template[curr+1] == '*' {
					skip()
					skip()
					inBlockComment = true
				} else {
					shift()
				}
			} else if char == '"' {
				shift()
				inString = true
			} else if char == '%' &&
				curr+1 < len(template) &&
				template[curr+1] == '%' {

				// Reached section marker.  Leave the body unmodified.
				for curr < len(template) {
					shift()
				}
			} else {
				shift()
			}
		}
	}

	return template[:shifted]
}

type bodyLexer struct {
	raw *rawBodyLexer

	lookAhead    []BodyToken
	lookAheadErr error
}

func (lexer *bodyLexer) fillTokensAndMaybeTrimWhitespaces() {
	if lexer.lookAheadErr != nil {
		return
	}

	for len(lexer.lookAhead) < 2 {
		token, err := lexer.raw.Next()
		if err != nil {
			lexer.lookAheadErr = err
			return
		}

		lexer.lookAhead = append(lexer.lookAhead, token)
	}

	if lexer.lookAhead[0].Id() == TextToken &&
		lexer.lookAhead[1].TrimLeadingWhitespaces() {

		text := lexer.lookAhead[0].(*Atom)
		length := len(text.Value)

		for length > 0 {
			char := text.Value[length-1]
			if char == ' ' || char == '\t' {
				length -= 1
				continue
			}

			if char == '\n' {
				length -= 1
			}

			// windows style \r\n.  Don't know why I even bothered checking ...
			if char == '\r' {
				length -= 1
			}

			break
		}

		text.Value = text.Value[:length]
	}

	if lexer.lookAhead[0].TrimTrailingWhitespaces() &&
		lexer.lookAhead[1].Id() == TextToken {

		text := lexer.lookAhead[1].(*Atom)
		start := 0

		for start < len(text.Value) {
			char := text.Value[start]
			if char == ' ' || char == '\t' || char == '\r' {
				start += 1
				continue
			}

			if char == '\n' {
				start += 1
			}

			break
		}

		text.Value = text.Value[start:]
	}
}

func (lexer *bodyLexer) Next() (Token, error) {
	lexer.fillTokensAndMaybeTrimWhitespaces()

	if len(lexer.lookAhead) > 0 {
		ret := lexer.lookAhead[0]
		lexer.lookAhead = lexer.lookAhead[1:]
		return ret, nil
	}

	if lexer.lookAheadErr == nil {
		panic("Programming error")
	}

	return nil, lexer.lookAheadErr
}
