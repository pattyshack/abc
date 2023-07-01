// Auto-generated from source: switch.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type Switch struct {
	ind  string
	stmt *template.Switch
}

func (Switch) Name() string { return "Switch" }

func (template *Switch) writeValue(
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

func (_template *Switch) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// switch.template:13:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"switch.template:13:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:13:4
	{
		_n, _err := _output.Write([]byte(`// `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:13:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"switch.template:13:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:13:20
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:14:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"switch.template:14:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:14:6
	{
		_n, _err := _output.Write([]byte(`switch `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:14:13
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Switch.Value),
			"switch.template:14:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:14:33
	{
		_n, _err := _output.Write([]byte(` {
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:15:0
	for _, branch := range stmt.Cases {
		// switch.template:16:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"switch.template:16:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:16:6
		{
			_n, _err := _output.Write([]byte(`case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:16:11
		{
			_n, _err := _template.writeValue(
				_output,
				(branch.Predicate.Value),
				"switch.template:16:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:16:36
		{
			_n, _err := _output.Write([]byte(`:
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:18:4
		{
			_n, _err := (&Body{ind + "\t", branch.Body}).WriteTo(_output)
			_numWritten += _n
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// switch.template:22:0
	if stmt.Default != nil {
		// switch.template:23:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"switch.template:23:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:23:6
		{
			_n, _err := _output.Write([]byte(`default:
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:25:4
		{
			_n, _err := (&Body{ind + "\t", stmt.Default.Body}).WriteTo(_output)
			_numWritten += _n
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// switch.template:28:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"switch.template:28:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:28:4
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
