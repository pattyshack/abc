package code_gen

import (
	"fmt"
	"io"

	lr "github.com/pattyshack/abc/src/lr/internal"
)

const (
	GoLang = "go"
)

func GenerateLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates,
	targetLanguage string) (
	io.WriterTo,
	error) {

	switch targetLanguage {
	case GoLang:
		return GenerateGoLRCode(grammar, states)
	}

	return nil, fmt.Errorf("Unsupported target language: %s", targetLanguage)
}
