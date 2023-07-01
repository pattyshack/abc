// Auto-generated from source: output-types.template

package main

import (
	_fmt "fmt"
	_io "io"
)

type OutputTypesTemplate struct {
	Custom interface{}
}

func (OutputTypesTemplate) Name() string { return "OutputTypesTemplate" }

func (template *OutputTypesTemplate) writeValue(
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

func (_template *OutputTypesTemplate) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	Custom := _template.Custom

	// output-types.template:8:0
	{
		_n, _err := _output.Write([]byte(`stringer: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:8:10
	{
		_n, _err := _template.writeValue(
			_output,
			(Custom),
			"output-types.template:8:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:8:17
	{
		_n, _err := _output.Write([]byte(`

string: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:10:8
	{
		_n, _err := _template.writeValue(
			_output,
			("abcd"),
			"output-types.template:10:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:10:17
	{
		_n, _err := _output.Write([]byte(`
bytes: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:11:7
	{
		_n, _err := _template.writeValue(
			_output,
			(append([]byte("hello"), []byte(" world")...)),
			"output-types.template:11:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:11:54
	{
		_n, _err := _output.Write([]byte(`

bool: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:13:6
	{
		_n, _err := _template.writeValue(
			_output,
			(true),
			"output-types.template:13:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:13:11
	{
		_n, _err := _output.Write([]byte(` / `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:13:14
	{
		_n, _err := _template.writeValue(
			_output,
			(false),
			"output-types.template:13:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:13:20
	{
		_n, _err := _output.Write([]byte(`

uint8 (aka byte): `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:15:18
	{
		_n, _err := _template.writeValue(
			_output,
			('a'),
			"output-types.template:15:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:15:24
	{
		_n, _err := _output.Write([]byte(`
int32 (aka rune): `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:16:18
	{
		_n, _err := _template.writeValue(
			_output,
			(rune(97)),
			"output-types.template:16:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:16:29
	{
		_n, _err := _output.Write([]byte(`

uint: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:18:6
	{
		_n, _err := _template.writeValue(
			_output,
			(uint(11)),
			"output-types.template:18:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:18:17
	{
		_n, _err := _output.Write([]byte(`
int: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:19:5
	{
		_n, _err := _template.writeValue(
			_output,
			(int(-11)),
			"output-types.template:19:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:19:16
	{
		_n, _err := _output.Write([]byte(`

uint64: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:21:8
	{
		_n, _err := _template.writeValue(
			_output,
			(uint64(17)),
			"output-types.template:21:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:21:21
	{
		_n, _err := _output.Write([]byte(`
int64: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:22:7
	{
		_n, _err := _template.writeValue(
			_output,
			(int64(-17)),
			"output-types.template:22:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:22:20
	{
		_n, _err := _output.Write([]byte(`

float32: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:24:9
	{
		_n, _err := _template.writeValue(
			_output,
			(float32(3.14159)),
			"output-types.template:24:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:24:28
	{
		_n, _err := _output.Write([]byte(`
float64: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:25:9
	{
		_n, _err := _template.writeValue(
			_output,
			(float64(2.71828)),
			"output-types.template:25:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:25:28
	{
		_n, _err := _output.Write([]byte(`

complex: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:27:9
	{
		_n, _err := _template.writeValue(
			_output,
			(complex(0, 1)),
			"output-types.template:27:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:27:25
	{
		_n, _err := _output.Write([]byte(`

invalid: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:29:9
	{
		_n, _err := _template.writeValue(
			_output,
			([]string{"foo", "bar"}),
			"output-types.template:29:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// output-types.template:29:34
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
