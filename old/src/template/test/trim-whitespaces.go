// Auto-generated from source: trim-whitespaces.template

package main

import (
	_fmt "fmt"
	_io "io"
)

type TrimWhitespacesTemplate struct {
}

func (TrimWhitespacesTemplate) Name() string { return "TrimWhitespacesTemplate" }

func (template *TrimWhitespacesTemplate) writeValue(
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

func (_template *TrimWhitespacesTemplate) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	// trim-whitespaces.template:5:0
	{
		_n, _err := _output.Write([]byte(`list 1: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// trim-whitespaces.template:6:0
	for idx, i := range []int{1, 2, 3, 4} {
		// trim-whitespaces.template:7:4
		if idx != 3 {
			// trim-whitespaces.template:8:0
			{
				_n, _err := _template.writeValue(
					_output,
					(i),
					"trim-whitespaces.template:8:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// trim-whitespaces.template:8:2
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// trim-whitespaces.template:10:0
			{
				_n, _err := _template.writeValue(
					_output,
					(i),
					"trim-whitespaces.template:10:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// trim-whitespaces.template:12:8
	{
		_n, _err := _output.Write([]byte(`

list 2:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// trim-whitespaces.template:15:0
	for idx, i := range []int{1, 2, 3, 4} {
		// trim-whitespaces.template:16:4
		if idx != 3 {
			// trim-whitespaces.template:16:22
			{
				_n, _err := _output.Write([]byte(` `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// trim-whitespaces.template:17:1
			{
				_n, _err := _template.writeValue(
					_output,
					(i),
					"trim-whitespaces.template:17:1")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// trim-whitespaces.template:17:3
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// trim-whitespaces.template:18:14
			{
				_n, _err := _output.Write([]byte(` `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// trim-whitespaces.template:19:1
			{
				_n, _err := _template.writeValue(
					_output,
					(i),
					"trim-whitespaces.template:19:1")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// trim-whitespaces.template:21:8
	{
		_n, _err := _output.Write([]byte(`

`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// trim-whitespaces.template:23:0
	for _, c := range "abcde" {
		// trim-whitespaces.template:23:32
		{
			_n, _err := _template.writeValue(
				_output,
				(string(c)),
				"trim-whitespaces.template:23:32")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// trim-whitespaces.template:23:53
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
