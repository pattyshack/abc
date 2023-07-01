// Auto-generated from source: body.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/abc/src/template/internal"
)

type Body struct {
	ind  string
	body []template.Statement
}

func (Body) Name() string { return "Body" }

func (template *Body) writeValue(
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

func (_template *Body) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	body := _template.body

	// body.template:14:0
	for _, statement := range body {
		// body.template:15:4
		switch stmt := statement.(type) {
		case *template.Atom:
			// body.template:18:8
			switch stmt.Id() {
			case template.CommentToken:
			case template.TextToken:
				// body.template:23:12
				{
					_n, _err := (&Text{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.SubstitutionToken:
				// body.template:26:12
				{
					_n, _err := (&Substitute{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.EmbedToken:
				// body.template:29:12
				{
					_n, _err := (&Embed{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.CopySectionToken:
				// body.template:32:12
				{
					_n, _err := (&CopySection{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.ContinueToken:
				// body.template:35:12
				{
					_n, _err := (&Continue{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.BreakToken:
				// body.template:38:12
				{
					_n, _err := (&Break{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.ReturnToken:
				// body.template:41:12
				{
					_n, _err := (&Return{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			case template.ErrorToken:
				// body.template:44:12
				{
					_n, _err := (&Error{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}
			default:
				// body.template:46:20
				{
					_n, _err := _output.Write([]byte(`
            // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// body.template:47:15
				{
					_n, _err := _template.writeValue(
						_output,
						(statement.Loc()),
						"body.template:47:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// body.template:47:33
				{
					_n, _err := _output.Write([]byte(`
            COMPILE ERROR: bug in template generation code
            Unexpected atom type: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// body.template:49:34
				{
					_n, _err := _template.writeValue(
						_output,
						(stmt.Id()),
						"body.template:49:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		case *template.For:
			// body.template:53:8
			{
				_n, _err := (&For{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}
		case *template.Switch:
			// body.template:56:8
			{
				_n, _err := (&Switch{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}
		case *template.If:
			// body.template:59:8
			{
				_n, _err := (&If{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}
		default:
			// body.template:61:16
			{
				_n, _err := _output.Write([]byte(`
        // `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// body.template:62:11
			{
				_n, _err := _template.writeValue(
					_output,
					(statement.Loc()),
					"body.template:62:11")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// body.template:62:29
			{
				_n, _err := _output.Write([]byte(`
        COMPILE ERROR: bug in template generation code
        Unexpected statement type: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// body.template:64:35
			{
				_n, _err := _template.writeValue(
					_output,
					(statement.Id()),
					"body.template:64:35")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}

	return _numWritten, nil
}
