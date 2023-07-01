// Auto-generated from source: if.template

package main

import (
	_fmt "fmt"
	_io "io"
)

type IfTemplate struct {
}

func (IfTemplate) Name() string { return "IfTemplate" }

func (template *IfTemplate) writeValue(
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

func (_template *IfTemplate) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	// if.template:6:0

	val := 3

	pred := func() bool { return false }

	// if.template:10:3
	{
		_n, _err := _output.Write([]byte(`
If Without Else If / Else 1
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:13:0
	if val*7 == 21 {
		// if.template:13:20
		{
			_n, _err := _output.Write([]byte(`val * 7 = 21`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:13:39
	{
		_n, _err := _output.Write([]byte(`

If Without Else If / Else 2
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:16:0
	if val == 3 && pred() {
		// if.template:16:27
		{
			_n, _err := _output.Write([]byte(`sadness`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:16:41
	{
		_n, _err := _output.Write([]byte(`

With Else Branch
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:19:0
	if val != 3 {
		// if.template:19:15
		{
			_n, _err := _output.Write([]byte(`val not equal 3`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else {
		// if.template:19:38
		{
			_n, _err := _output.Write([]byte(`val equals 3`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:19:57
	{
		_n, _err := _output.Write([]byte(`

With Else If Branches 1
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:22:0
	if val == 0 {
		// if.template:22:16
		{
			_n, _err := _output.Write([]byte(`0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 1 {
		// if.template:24:21
		{
			_n, _err := _output.Write([]byte(`1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 2 {
		// if.template:26:21
		{
			_n, _err := _output.Write([]byte(`2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:28:8
	{
		_n, _err := _output.Write([]byte(`
With Else If Branches 2
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:31:0
	if val == 0 {
		// if.template:31:16
		{
			_n, _err := _output.Write([]byte(`0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 1 {
		// if.template:33:21
		{
			_n, _err := _output.Write([]byte(`1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 2 {
		// if.template:35:21
		{
			_n, _err := _output.Write([]byte(`2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 3 {
		// if.template:37:21
		{
			_n, _err := _output.Write([]byte(`3
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// if.template:39:8
	{
		_n, _err := _output.Write([]byte(`
With Else If And Else Branches
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// if.template:42:0
	if val == 0 {
		// if.template:42:16
		{
			_n, _err := _output.Write([]byte(`0
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 1 {
		// if.template:44:21
		{
			_n, _err := _output.Write([]byte(`1
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}

	} else if val == 2 {
		// if.template:46:21
		{
			_n, _err := _output.Write([]byte(`2
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	} else {
		// if.template:48:9
		{
			_n, _err := _output.Write([]byte(`other
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	return _numWritten, nil
}
