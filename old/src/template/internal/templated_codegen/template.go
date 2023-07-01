package templated_codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"

	"github.com/pattyshack/abc/src/template/internal"
)

type Template struct {
	source string
	file   *template.File

	shouldFormat bool
}

func NewTemplate(source string, file *template.File) io.WriterTo {
	return &Template{source, file, true}
}

func (template *Template) WriteTo(output io.Writer) (int64, error) {
	buffer := bytes.NewBuffer(nil)

	_, err := (&File{template.source, template.file}).WriteTo(buffer)
	if err != nil {
		return 0, err
	}

	formatted := buffer.Bytes()
	if template.shouldFormat {
		formatted, err = format.Source(buffer.Bytes())
		if err != nil {
			return 0, fmt.Errorf(
				"Failed to format (%s) generated code:\n%s",
				err,
				buffer.Bytes())
		}
	}

	n, err := output.Write(formatted)
	return int64(n), err
}
