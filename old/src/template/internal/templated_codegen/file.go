// Auto-generated from source: file.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type File struct {
	source string
	spec   *template.File
}

func (File) Name() string { return "File" }

func (template *File) writeValue(
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

func (_template *File) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	source := _template.source
	spec := _template.spec

	// file.template:13:0
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:13:31
	{
		_n, _err := _template.writeValue(
			_output,
			(source),
			"file.template:13:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:13:40
	{
		_n, _err := _output.Write([]byte(`

package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:15:8
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.PackageName),
			"file.template:15:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:15:27
	{
		_n, _err := _output.Write([]byte(`

import (
	_fmt "fmt"
	_io "io"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:20:0
	if spec.Imports != "" {
		// file.template:20:26
		{
			_n, _err := _output.Write([]byte(`

`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:22:0
		{
			_n, _err := _template.writeValue(
				_output,
				(spec.Imports),
				"file.template:22:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:23:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:26:5
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:26:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:26:25
	{
		_n, _err := _output.Write([]byte(` struct {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:27:0
	for _, arg := range spec.Arguments {
		// file.template:27:39
		{
			_n, _err := _output.Write([]byte(`
	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:28:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:28:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:28:12
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:28:13
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Type),
				"file.template:28:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:29:8
	{
		_n, _err := _output.Write([]byte(`
}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:32:6
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:32:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:32:26
	{
		_n, _err := _output.Write([]byte(`) Name() string { return "`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:32:52
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:32:52")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:32:72
	{
		_n, _err := _output.Write([]byte(`" }

func (template *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:34:16
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:34:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:34:36
	{
		_n, _err := _output.Write([]byte(`) writeValue(
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
		valueBytes = val`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:49:0
	for _, primitive := range template.OutputablePrimitiveTypes {
		// file.template:49:64
		{
			_n, _err := _output.Write([]byte(`
	case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:50:6
		{
			_n, _err := _template.writeValue(
				_output,
				(primitive),
				"file.template:50:6")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:50:16
		{
			_n, _err := _output.Write([]byte(`:
		valueBytes = []byte(_fmt.Sprintf("%v", val))`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:52:9
	{
		_n, _err := _output.Write([]byte(`
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:64:17
	{
		_n, _err := _template.writeValue(
			_output,
			(spec.TemplateName),
			"file.template:64:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:64:37
	{
		_n, _err := _output.Write([]byte(`) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:71:0
	for idx, arg := range spec.Arguments {
		// file.template:72:4
		if idx == 0 {
			// file.template:72:20
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:74:13
		{
			_n, _err := _output.Write([]byte(`	`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:75:1
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:75:1")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:75:12
		{
			_n, _err := _output.Write([]byte(` := _template.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:75:26
		{
			_n, _err := _template.writeValue(
				_output,
				(arg.Name),
				"file.template:75:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:75:37
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:76:7
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:0
	{
		_n, _err := (&Body{"\t", spec.Body}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:34
	{
		_n, _err := _output.Write([]byte(`
	return _numWritten, nil
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
