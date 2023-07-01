// Auto-generated from source: substitute.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type Substitute struct {
	ind  string
	stmt *template.Atom
}

func (Substitute) Name() string { return "Substitute" }

func (template *Substitute) writeValue(
	output _io.Writer,
	value interface{},
	loc string) (
	int,
	error) {

	var valueBytes []byte
	switch val := value.(type) {
	case _fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *Substitute) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// substitute.template:14:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:14:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:14:4
	{
		_n, _err := _output.Write([]byte(`// `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:14:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"substitute.template:14:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:14:20
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:15:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:15:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:15:4
	{
		_n, _err := _output.Write([]byte(`{
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:16:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:16:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:16:4
	{
		_n, _err := _output.Write([]byte(`	_n, _err := _template.writeValue(
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:17:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:17:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:17:4
	{
		_n, _err := _output.Write([]byte(`		_output,
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:18:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:18:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:18:4
	{
		_n, _err := _output.Write([]byte(`		(`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:18:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Value),
			"substitute.template:18:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:18:20
	{
		_n, _err := _output.Write([]byte(`),
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:19:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:19:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:19:4
	{
		_n, _err := _output.Write([]byte(`		"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:19:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"substitute.template:19:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:19:20
	{
		_n, _err := _output.Write([]byte(`")
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:20:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:20:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:20:4
	{
		_n, _err := _output.Write([]byte(`	_numWritten += int64(_n)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:21:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:21:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:21:4
	{
		_n, _err := _output.Write([]byte(`	if _err != nil {
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:22:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:22:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:22:4
	{
		_n, _err := _output.Write([]byte(`		return _numWritten, _err
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:23:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:23:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:23:4
	{
		_n, _err := _output.Write([]byte(`	}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:24:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"substitute.template:24:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// substitute.template:24:4
	{
		_n, _err := _output.Write([]byte(`}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
