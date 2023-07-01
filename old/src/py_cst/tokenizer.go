package py_cst

import (
	"fmt"
	"io"
)

type Tokenizer struct {
	reader *LocationReader

	indentLevels []int

	openness int // how many levels of open '{' '(' '['

	tryTokenizeIndentation bool

	reachedEOF bool
}

func newTokenizer(fileName string, reader io.Reader) *Tokenizer {
	return &Tokenizer{
		reader:                 NewLocationReader(fileName, reader),
		indentLevels:           []int{0},
		tryTokenizeIndentation: true,
	}
}

func (c *Tokenizer) CurrentLocation() Location {
	return c.reader.Location
}

func (c *Tokenizer) node() Node {
	return Node{Location: c.reader.Location}
}

// hand rolling a tokenizer =...(
func (c *Tokenizer) Next() (*Token, error) {
	// must token space first to strip spaces
	token, err := c.stripSpacesAndMaybeTokenizeIndentation()
	if err != nil {
		return nil, err
	}

	if token != nil {
		return token, nil
	}

	// strings must be tokenized before names due to byte / unicode / raw
	// prefixes
	token, err = c.maybeTokenizeString()
	if err != nil {
		return nil, err
	}

	if token != nil {
		return token, nil
	}

	// TODO handle bytes
	token = c.maybeTokenizeName()
	if token != nil {
		return token, nil
	}

	// numbers must be tokenized before symbols due to negation signs.
	token = c.maybeTokenizeNumber()
	if token != nil {
		return token, nil
	}

	token = c.maybeTokenizeSymbol()
	if token != nil {
		c.tryTokenizeIndentation = (token.TokenType == NEWLINE)
		return token, nil
	}

	token = c.maybeTokenizeComment()
	if token != nil {
		c.tryTokenizeIndentation = true
		return token, nil
	}

	b, err := c.reader.Peek(1)
	if err != nil {
		if err == io.EOF {
			if !c.reachedEOF { // first time reaching EOF
				c.reachedEOF = true

				// return NEWLINE to end simple_stmt
				return &Token{c.node(), NEWLINE, ""}, nil
			}

			// DEDENT to close suite scopes
			for len(c.indentLevels) > 1 {
				c.indentLevels = c.indentLevels[:len(c.indentLevels)-1]

				return &Token{c.node(), DEDENT, ""}, nil
			}
		}

		return nil, err
	}

	return nil, fmt.Errorf(
		"Unsupported token: %s (%v)",
		string(b),
		c.reader.Location)
}

func (c *Tokenizer) stripSpacesAndMaybeTokenizeIndentation() (*Token, error) {
	tryTokenizeIndentation := c.tryTokenizeIndentation
	c.tryTokenizeIndentation = false

	// TODO: handle more gracefully.  we should scan until we reach the first
	// non-space character.
	bytes, _ := c.reader.Peek(128)
	if len(bytes) == 0 {
		return nil, nil
	}

	match := spaceRe.Find(bytes)

	// don't tokenize indentation if the line has no syntactic meaning.
	remaining := len(bytes) - len(match)
	if remaining > 0 &&
		(bytes[len(match)] == '\n' || bytes[len(match)] == '#') {
		tryTokenizeIndentation = false
	} else if remaining > 1 &&
		bytes[len(match)] == '\r' &&
		bytes[len(match)+1] == '\n' {

		tryTokenizeIndentation = false
	}

	var token *Token

	if tryTokenizeIndentation && c.openness == 0 {
		indentLevel := 0
		for _, b := range match {
			if b == '\t' {
				indentLevel += 8 - (indentLevel % 8)
			} else {
				indentLevel += 1
			}
		}

		prev := c.indentLevels[len(c.indentLevels)-1]

		if prev > indentLevel {
			fmt.Println("INDENTS", c.indentLevels, indentLevel, c.reader.Location)
			c.indentLevels = c.indentLevels[:len(c.indentLevels)-1]

			if indentLevel > c.indentLevels[len(c.indentLevels)-1] {
				return nil, fmt.Errorf(
					"Unexpected indent: %v",
					c.reader.Location)
			}

			// Don't strip whitespace yet since we may need to emit more
			// DEDENT tokens
			c.tryTokenizeIndentation = true

			return &Token{c.node(), DEDENT, ""}, nil

		} else if prev < indentLevel {
			c.indentLevels = append(c.indentLevels, indentLevel)
			token = &Token{c.node(), INDENT, ""}
		}
	}

	if len(match) == 0 {
		return nil, nil
	}

	n, err := c.reader.Read(match)
	if len(match) != n || err != nil {
		panic(err) // should never happen
	}

	return token, nil
}

func (c *Tokenizer) maybeTokenizeString() (*Token, error) {
	bytes, _ := c.reader.Peek(5)

	prefix := string(bytes)

	pos := 1
	if len(prefix) > 1 &&
		(prefix[0] == 'u' || prefix[0] == 'U' ||
			prefix[0] == 'b' || prefix[0] == 'B') {

		prefix = prefix[1:]
		pos += 1
	}

	if len(prefix) > 1 && (prefix[0] == 'r' || prefix[0] == 'R') {
		prefix = prefix[1:]
		pos += 1
	}

	if len(prefix) == 0 || (prefix[0] != '\'' && prefix[0] != '"') {
		return nil, nil
	}

	long := false
	delimiter := prefix[0]
	preSize := 1024

	if len(prefix) >= 3 && (prefix[:3] == "\"\"\"" || prefix[:3] == "'''") {
		long = true
		preSize = 128 * 1024
		pos += 2
	}

	startLocation := c.reader.Location
	consumed := make([]byte, 0, preSize)

	peekSize := 256

	end := -1

	skip := false
	matchCount := 0

	bytes, _ = c.reader.Peek(peekSize)
	for {
		for ; pos < len(bytes); pos++ {
			if skip {
				skip = false
				continue
			}

			if bytes[pos] == '\\' { // skip escaped character
				matchCount = 0
				skip = true
				continue
			}

			if bytes[pos] != delimiter {
				matchCount = 0
				continue
			}

			if !long {
				end = pos + 1
				break
			}

			matchCount += 1
			if matchCount >= 3 {
				end = pos + 1
				break
			}
		}

		if end > 0 {
			consumed = append(consumed, bytes[:end]...)
			n, err := c.reader.Read(bytes[:end])
			if n != end || err != nil {
				panic(err) // should never happen
			}

			break
		}

		consumed = append(consumed, bytes...)
		n, err := c.reader.Read(bytes)
		if n != len(bytes) || err != nil {
			panic(err) // should never happen
		}

		if peekSize < 1024*1024 {
			peekSize *= 2
		}

		pos = 0
		bytes, _ = c.reader.Peek(peekSize)
		if len(bytes) == 0 {
			break
		}
	}

	if end < 0 {
		return nil, fmt.Errorf(
			"Unsupported token: %s (%v)",
			string(delimiter),
			startLocation)
	}

	return &Token{
		Node{Location: startLocation},
		STRING,
		string(consumed),
	}, nil
}

func (c *Tokenizer) maybeTokenizeName() *Token {
	// TODO: handle more gracefully.  Scan until the first \W
	bytes, _ := c.reader.Peek(128)

	name := nameRe.Find(bytes)
	if name == nil {
		return nil
	}

	token := &Token{c.node(), NAME, ""}

	n, err := c.reader.Read(name)
	if len(name) != n || err != nil {
		panic(err) // should never happen
	}

	token.Value = string(name)

	code, ok := keywords[token.Value]
	if ok {
		token.TokenType = code
	}

	return token
}

func (c *Tokenizer) maybeTokenizeNumber() *Token {
	// TODO: handle more gracefully.  This may cut off a really long number ...
	bytes, _ := c.reader.Peek(128)

	// NOTE: float parsing must parsed before int (in order to capture
	// decimal & exponent).
	floatVal := floatRe.Find(bytes)
	if floatVal != nil {
		token := &Token{c.node(), FLOAT, string(floatVal)}

		n, err := c.reader.Read(floatVal)
		if len(floatVal) != n || err != nil {
			panic(err) // should never happen
		}

		return token
	}

	intVal := intRe.Find(bytes)
	if intVal == nil {
		return nil
	}

	token := &Token{c.node(), INTEGER, string(intVal)}

	n, err := c.reader.Read(intVal)
	if len(intVal) != n || err != nil {
		panic(err) // should never happen
	}

	return token
}

func (c *Tokenizer) maybeTokenizeSymbol() *Token {
	bytes, _ := c.reader.Peek(3)

	var token *Token
	for i := len(bytes); i > 0; i-- {
		val := string(bytes[:i])
		code, ok := symbols[val]
		if ok {
			token = &Token{c.node(), code, val}

			n, err := c.reader.Read(bytes[:i])
			if i != n || err != nil {
				panic(err) // should never happen
			}

			break
		}
	}

	if token == nil {
		return nil
	}

	if token.TokenType == LEFT_PARENTHESIS ||
		token.TokenType == LEFT_BRACKET ||
		token.TokenType == LEFT_BRACE {

		c.openness += 1
	}

	if token.TokenType == RIGHT_PARENTHESIS ||
		token.TokenType == RIGHT_BRACKET ||
		token.TokenType == RIGHT_BRACE {

		c.openness -= 1
	}

	return token
}

func (c *Tokenizer) maybeTokenizeComment() *Token {
	bytes, _ := c.reader.Peek(1)

	if len(bytes) == 0 {
		return nil
	}

	firstByte := bytes[0]

	if firstByte == '#' {
		token := &Token{c.node(), COMMENT_NEWLINE, ""}

		comment, err := c.reader.ReadString('\n')
		if err != nil {
			return nil
		}

		token.Value = comment[:len(comment)-1]
		if len(token.Value) > 0 &&
			token.Value[len(token.Value)-1] == '\r' {
			token.Value = token.Value[:len(token.Value)-1]
		}

		return token
	}

	return nil
}
