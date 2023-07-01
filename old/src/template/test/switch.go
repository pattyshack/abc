// Auto-generated from source: switch.template

package main

import (
	_fmt "fmt"
	_io "io"
)

type SwitchTemplate struct {
	TypeSwitch  interface{}
	ValueSwitch string
}

func (SwitchTemplate) Name() string { return "SwitchTemplate" }

func (template *SwitchTemplate) writeValue(
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

func (_template *SwitchTemplate) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	TypeSwitch := _template.TypeSwitch
	ValueSwitch := _template.ValueSwitch

	// switch.template:9:0
	{
		_n, _err := _output.Write([]byte(`
Type Switch
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:11:0
	switch v := TypeSwitch.(type) {
	case int:
		// switch.template:12:12
		{
			_n, _err := _output.Write([]byte(`int: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:12:17
		{
			_n, _err := _template.writeValue(
				_output,
				(v),
				"switch.template:12:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:12:19
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	case string:
		// switch.template:13:15
		{
			_n, _err := _output.Write([]byte(`string: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:13:23
		{
			_n, _err := _template.writeValue(
				_output,
				(v),
				"switch.template:13:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:13:25
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	default:
		// switch.template:14:11
		{
			_n, _err := _output.Write([]byte(`other: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:14:18
		{
			_n, _err := _template.writeValue(
				_output,
				(v),
				"switch.template:14:18")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// switch.template:14:20
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// switch.template:15:8
	{
		_n, _err := _output.Write([]byte(`
Value Switch
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// switch.template:18:0
	switch ValueSwitch {
	case "hello":
		// switch.template:18:38
		{
			_n, _err := _output.Write([]byte(`you say hello, and I say goodbye
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	case "goodbye":
		// switch.template:19:18
		{
			_n, _err := _output.Write([]byte(`you say goodbye, and I say hello
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	case "other":
	case "other2":
	}
	// switch.template:20:40
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
