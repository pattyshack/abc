package codegenutil

import (
	"fmt"
	"io"
)

type line struct {
	indent      string
	indentLevel int
	format      string
	args        []interface{}
}

func (l *line) WriteTo(writer io.Writer) (int64, error) {
	total := int64(0)

	if l.format != "" {
		for i := 0; i < l.indentLevel; i++ {
			n, err := writer.Write([]byte(l.indent))
			total += int64(n)

			if err != nil {
				return total, err
			}
		}

		n, err := fmt.Fprintf(writer, l.format, l.args...)
		total += int64(n)

		if err != nil {
			return total, err
		}
	}

	n, err := writer.Write([]byte{'\n'})
	total += int64(n)
	if err != nil {
		return total, err
	}

	return total, nil
}

type CodeBuilder struct {
	chunks []io.WriterTo

	indent      string
	indentLevel int
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		chunks:      nil,
		indent:      "    ",
		indentLevel: 0,
	}
}

func (cb *CodeBuilder) PushIndent() {
	cb.indentLevel += 1
}

func (cb *CodeBuilder) PopIndent() {
	if cb.indentLevel > 0 {
		cb.indentLevel -= 1
	}
}

func (cb *CodeBuilder) Line(format string, args ...interface{}) {
	cb.chunks = append(
		cb.chunks,
		&line{cb.indent, cb.indentLevel, format, args})
}

func (cb *CodeBuilder) Embed(template io.WriterTo) {
	cb.chunks = append(cb.chunks, template)
}

func (cb *CodeBuilder) WriteTo(output io.Writer) (int64, error) {
	total := int64(0)
	for _, chunk := range cb.chunks {
		n, err := chunk.WriteTo(output)
		total += n

		if err != nil {
			return total, err
		}
	}

	return total, nil
}
