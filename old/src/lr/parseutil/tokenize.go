package parseutil

import (
	"io"
)

func IsWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

func stripLeadingWhitespaces(reader *LocationReader) (bool, error) {
	consumedBytes := false

	for {
		bytes, err := reader.Peek(1)
		if err != nil {
			return consumedBytes, err
		}

		char := bytes[0]
		if IsWhitespace(char) {
			_, err = reader.ReadByte()
			if err != nil {
				panic(err) // should never happen
			}

			consumedBytes = true
		} else {
			break
		}
	}

	return consumedBytes, nil
}

// Strip all leading whitespaces
func StripLeadingWhitespaces(reader *LocationReader) error {
	_, err := stripLeadingWhitespaces(reader)
	return err
}

// Strip all leading whitespaces as well as golang style comments
// (i.e., /* */  and //)
func StripLeadingWhitespacesAndComments(reader *LocationReader) error {
	modified := true
	for modified {
		modified = false

		var err error
		modified, err = stripLeadingWhitespaces(reader)
		if err != nil {
			return err
		}

		bytes, _ := reader.Peek(2)

		if string(bytes) == "//" {
			for {
				char, err := reader.ReadByte()
				if err != nil {
					return err
				}

				if char == '\n' {
					break
				}
			}

			modified = true
			continue
		}

		if string(bytes) == "/*" {
			n, err := reader.Read(bytes)
			if n != 2 || err != nil {
				panic(err) // should never happen
			}

			for {
				bytes, err = reader.Peek(2)
				if err != nil {
					return err
				}

				if string(bytes) == "*/" {
					n, err := reader.Read(bytes)
					if n != 2 || err != nil {
						panic(err) // should never happen
					}

					break
				}

				_, err = reader.ReadByte()
				if err != nil {
					panic(err)
				}
			}

			modified = true
		}
	}

	return nil
}

type Symbol struct {
	Value string // can't be empty string
	Id    int
}

type Symbols []Symbol

func (s Symbols) Len() int { return len(s) }

func (s Symbols) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }

// longer symbol before shorter symbol
// symbol value as tie break
func (s Symbols) Less(i int, j int) bool {
	if len(s[i].Value) == len(s[j].Value) {
		return s[i].Value < s[j].Value
	}

	return len(s[i].Value) > len(s[j].Value)
}

func MaybeTokenizeSymbol(
	reader *LocationReader,
	sortedSymbols Symbols) (
	*Symbol,
	Location,
	error) {

	var bytes []byte
	var err error
	for idx, symbol := range sortedSymbols {
		if idx == 0 {
			bytes, err = reader.Peek(len(symbol.Value))
			if err != nil {
				if err != io.EOF || len(bytes) == 0 {
					return nil, Location{}, err
				}
			}
		}

		if len(bytes) < len(symbol.Value) {
			continue
		}

		if string(bytes[:len(symbol.Value)]) == symbol.Value {
			loc := reader.Location

			bytes = bytes[:len(symbol.Value)]
			n, err := reader.Read(bytes)
			if len(bytes) != n || err != nil {
				panic(err) // should never happen
			}

			return &symbol, loc, nil
		}
	}

	return nil, Location{}, nil
}

// If the reader's leading bytes are ascii character of the form 'x' (or
// '\t' for escaped characters), pop those bytes off the reader and return
// the value.  Otherwise, return a nil slice.
func MaybeTokenizeCharacter(reader *LocationReader) (string, Location, error) {
	bytes, err := reader.Peek(4)
	if err != nil {
		return "", Location{}, err
	}

	if len(bytes) < 3 {
		return "", Location{}, nil
	}

	if bytes[0] != '\'' {
		return "", Location{}, nil
	}

	numBytes := 3
	if bytes[1] == '\\' { // c escape
		if len(bytes) < 4 || bytes[3] != '\'' {
			return "", Location{}, nil
		}

		switch bytes[2] {
		case 't', 'n', '\'', '\\':
		default:
			return "", Location{}, nil
		}

		numBytes = 4
	} else {
		if bytes[2] != '\'' {
			return "", Location{}, nil
		}

		valid := false

		char := bytes[1]
		if ('a' <= char && char <= 'z') ||
			('A' <= char && char <= 'Z') ||
			('0' <= char && char <= '9') {

			valid = true
		} else {
			switch char {
			case '`', '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
				'-', '_', '=', '+', '[', '{', ']', '}', '|', ';', ':', '"',
				',', '<', '.', '>', '/', '?', ' ':
				valid = true

			}
		}

		if !valid {
			return "", Location{}, nil
		}
	}

	loc := reader.Location

	bytes = bytes[:numBytes]
	n, err := reader.Read(bytes)
	if len(bytes) != n || err != nil {
		panic(err) // should never happen
	}

	return string(bytes), loc, nil
}

// If the reader's leading bytes are identifer of the form [a-zA-Z_]\w* ,
// pop those bytes off the reader and return the value.  Otherwise, return a
// nil slice.
func MaybeTokenizeIdentifier(reader *LocationReader) (string, Location, error) {
	peekRange := 32
	prevLen := 0
	checkIdx := 0

	var bytes []byte
	var err error
	for {
		bytes, err = reader.Peek(peekRange)
		if err != nil && err != io.EOF {
			return "", Location{}, err
		}

		if len(bytes) == 0 {
			return "", Location{}, nil
		}

		if checkIdx == 0 {
			char := bytes[0]
			if !(('a' <= char && char <= 'z') ||
				('A' <= char && char <= 'Z') ||
				char == '_') {

				return "", Location{}, nil
			}

			checkIdx = 1
		}

		if prevLen == len(bytes) { // ran out of bytes to read
			break
		}

		foundEnd := false
		for checkIdx < len(bytes) {
			char := bytes[checkIdx]
			if !(('a' <= char && char <= 'z') ||
				('A' <= char && char <= 'Z') ||
				('0' <= char && char <= '9') ||
				char == '_') {

				foundEnd = true
				break
			}
			checkIdx += 1
		}

		if foundEnd {
			break
		}

		prevLen = len(bytes)
		peekRange *= 2
	}

	loc := reader.Location

	bytes = bytes[:checkIdx]
	n, err := reader.Read(bytes)
	if len(bytes) != n || err != nil {
		panic(err) // should never happen
	}

	return string(bytes), loc, nil
}
