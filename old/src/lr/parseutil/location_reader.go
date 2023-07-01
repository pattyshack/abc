package parseutil

import (
	"bufio"
	"fmt"
	"io"
)

// maybe track end line / end column?
type Location struct {
	FileName string
	Line     int // one-base
	Column   int // zero-base
}

func (l Location) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l Location) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

// Not efficient, but probably good enough
type LocationReader struct {
	Location

	reader *bufio.Reader
}

func NewLocationReader(fileName string, reader io.Reader) *LocationReader {
	return &LocationReader{
		Location: Location{
			FileName: fileName,
			Line:     1,
		},
		reader: bufio.NewReaderSize(reader, 1024*1024),
	}
}

func (r *LocationReader) Peek(n int) ([]byte, error) {
	return r.reader.Peek(n)
}

func (r *LocationReader) ReadByte() (byte, error) {
	b, err := r.reader.ReadByte()
	if err != nil {
		return 0, err
	}

	if b == '\n' {
		r.Line += 1
		r.Column = 0
	} else {
		r.Column += 1
	}

	return b, err
}

func (r *LocationReader) Read(buf []byte) (int, error) {
	n, err := r.reader.Read(buf)

	for i := 0; i < n; i++ {
		b := buf[i]
		if b == '\n' {
			r.Line += 1
			r.Column = 0
		} else {
			r.Column += 1
		}
	}

	return n, err
}
func (r *LocationReader) ReadString(delim byte) (string, error) {
	s, err := r.reader.ReadString(delim)

	for _, b := range s {
		if b == '\n' {
			r.Line += 1
			r.Column = 0
		} else {
			r.Column += 1
		}
	}

	return s, err
}
