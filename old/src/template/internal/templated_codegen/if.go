// Auto-generated from source: if.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type If struct {
	ind  string
	stmt *template.If
}

func (If) Name() string { return "If" }

func (template *If) writeValue(
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

func (_template *If) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// if.template:13:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"if.template:13:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:13:4
	{
		_n, _err := _output.Write([]byte(`// `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:13:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"if.template:13:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:13:20
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:14:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"if.template:14:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:14:6
	{
		_n, _err := _output.Write([]byte(`if `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:14:9
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.If.Predicate.Value),
			"if.template:14:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:14:35
	{
		_n, _err := _output.Write([]byte(` {
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:16:0
	{
		_n, _err := (&Body{ind + "\t", stmt.If.Body}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:18:0
	for _, branch := range stmt.ElseIfs {
		// if.template:18:41
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:20:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"if.template:20:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:20:4
		{
			_n, _err := _output.Write([]byte(`} else if `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:20:14
		{
			_n, _err := _template.writeValue(
				_output,
				(branch.Predicate.Value),
				"if.template:20:14")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:20:39
		{
			_n, _err := _output.Write([]byte(` {
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:22:4
		{
			_n, _err := (&Body{ind + "\t", branch.Body}).WriteTo(_output)
			_numWritten += _n
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:26:0
	if stmt.Else != nil {
		// if.template:27:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"if.template:27:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:27:4
		{
			_n, _err := _output.Write([]byte(`} else {
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// if.template:29:4
		{
			_n, _err := (&Body{ind + "\t", stmt.Else.Body}).WriteTo(_output)
			_numWritten += _n
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:32:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"if.template:32:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:32:4
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
