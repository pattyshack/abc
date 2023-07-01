// Auto-generated from source: for.template

package main

import (
	_fmt "fmt"
	_io "io"

	fmt2 "fmt"
)

type ForTemplate struct {
	Chan  <-chan int
	Count int
}

func (ForTemplate) Name() string { return "ForTemplate" }

func (template *ForTemplate) writeValue(
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

func (_template *ForTemplate) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	Chan := _template.Chan
	Count := _template.Count

	// for.template:16:3
	{
		_n, _err := _output.Write([]byte(`Infinite Loop:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:18:0
	for {
		// for.template:19:4
		Count += 1
		// for.template:20:4
		if Count%3 == 0 {
			// for.template:20:27
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:22:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:22:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:22:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 0`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:23:8
			continue

		} else if Count%3 == 1 {
			// for.template:24:32
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:26:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:26:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:26:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 1`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// for.template:27:14
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:29:2
			{
				_n, _err := _template.writeValue(
					_output,
					(Count),
					"for.template:29:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:29:8
			{
				_n, _err := _output.Write([]byte(` % 3 == 2`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// for.template:30:8
			if Count > 10 {
				// for.template:31:12
				break
			}
		}
	}
	// for.template:34:7
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:38:3
	{
		_n, _err := _output.Write([]byte(`
Predicate Loop
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:41:0
	for Count < 30 {
		// for.template:41:19
		{
			_n, _err := _output.Write([]byte(`  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:42:2
		{
			_n, _err := _template.writeValue(
				_output,
				(Count),
				"for.template:42:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:42:8
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:44:4

		Count = (Count +
			1) * 2

	}
	// for.template:52:3
	{
		_n, _err := _output.Write([]byte(`
Counter Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:55:0
	for i := 0; i < 5; i++ {
		// for.template:55:27
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:56:2
		{
			_n, _err := _template.writeValue(
				_output,
				(i),
				"for.template:56:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:56:4
		{
			_n, _err := _output.Write([]byte(`.0`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// for.template:57:8
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:61:3
	{
		_n, _err := _output.Write([]byte(`
Slice Range Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:64:0
	for _, item := range []string{"foo", "bar"} {
		// for.template:64:48
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:65:2
		{
			_n, _err := _template.writeValue(
				_output,
				(fmt2.Sprintf("item: %s", item)),
				"for.template:65:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// for.template:66:8
	{
		_n, _err := _output.Write([]byte(`

Map Range Loop`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:69:0
	for key, val := range map[string]int{"key1": 1, "key2": 2} {
		// for.template:69:61
		{
			_n, _err := _output.Write([]byte(`
  `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:70:2
		{
			_n, _err := _template.writeValue(
				_output,
				(key),
				"for.template:70:2")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:70:6
		{
			_n, _err := _output.Write([]byte(` -> `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:70:10
		{
			_n, _err := _template.writeValue(
				_output,
				(val),
				"for.template:70:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// for.template:71:8
	{
		_n, _err := _output.Write([]byte(`

Channel Range Loop
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:74:0
	for val := range Chan {
		// for.template:74:27
		{
			_n, _err := _template.writeValue(
				_output,
				(val),
				"for.template:74:27")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:74:31
		{
			_n, _err := _output.Write([]byte(`,`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// for.template:74:39
	{
		_n, _err := _output.Write([]byte(`

Infinite Loop 2:
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// for.template:77:0
	cnt := 0
	// for.template:78:0
	for {
		// for.template:79:4
		if cnt > 3 {
			// for.template:80:8
			return _numWritten, nil
		}
		// for.template:81:14
		{
			_n, _err := _output.Write([]byte(`Iteration: `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:82:11
		{
			_n, _err := _template.writeValue(
				_output,
				(cnt),
				"for.template:82:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:82:15
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// for.template:84:4
		cnt += 1
	}
	// for.template:85:8
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
