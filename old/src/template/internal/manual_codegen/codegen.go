package manual_codegen

import (
	"io"

    "github.com/pattyshack/abc/src/lr/codegenutil"
	"github.com/pattyshack/abc/src/template/internal"
)

type Template struct {
    source string
	spec *template.File
}

func NewTemplate(source string, spec *template.File) io.WriterTo {
	return &Template{source, spec}
}

func (temp *Template) WriteTo(output io.Writer) (int64, error) {
	spec := temp.spec
	builder := codegenutil.NewCodeBuilder()
    imports := codegenutil.NewGoImports()

	l := builder.Line

    l("// Auto-generated from source: %s", temp.source)
    l("")
    l("package %s", spec.PackageName)
    l("")
    builder.Embed(imports)
    l("")

	if spec.Imports != "" {
		l("import (")
		l("%s", spec.Imports)
		l(")")
	}

	l("type %s struct {", spec.TemplateName)
	for _, arg := range spec.Arguments {
		l("    %s %s", arg.Name, arg.Type)
	}
	l("}")
	l("")

	l("func (%s) Name() string { return \"%s\" }",
		spec.TemplateName,
		spec.TemplateName)
	l("")

	l("func (template *%s) writeValue(", spec.TemplateName)
	l("    output %v, value interface{}, loc string) (int, error) {",
		imports.Obj("io.Writer"))
	builder.PushIndent()
	l("var valueBytes []byte")
	l("switch val := value.(type) {")
	l("case %v: valueBytes = []byte(val.String())",
		imports.Obj("fmt.Stringer"))
	l("case string: valueBytes = []byte(val)")
	l("case []byte: valueBytes = val")
	for _, primitive := range template.OutputablePrimitiveTypes {
		l("case %s: valueBytes = []byte(%v(\"%%v\", val))",
			primitive,
			imports.Obj("fmt.Sprintf"))
	}
	l("default:	return 0, %v(\"Unsupported output value type (%%s): %%v\", loc, value)",
		imports.Obj("fmt.Errorf"))
	l("}")
	l("")

	l("return output.Write(valueBytes)")

	builder.PopIndent()
	l("}")
	l("")

	l("func (_template *%s) WriteTo(_output %s) (int64, error) {",
		spec.TemplateName,
		imports.Obj("io.Writer"))
	builder.PushIndent()
	l("_numWritten := int64(0)")
	l("")

	for _, arg := range spec.Arguments {
		l("%s := _template.%s", arg.Name, arg.Name)
	}

	buildBody(spec.Body, builder, imports)

	l("")
	l("return _numWritten, nil")
	builder.PopIndent()
	l("}")

    return codegenutil.NewFormattedGoSource(builder).WriteTo(output)
}

func buildBody(
    body []template.Statement,
    builder *codegenutil.CodeBuilder,
    imports *codegenutil.GoImports) {

	l := builder.Line
	push := builder.PushIndent
	pop := builder.PopIndent

	for _, statement := range body {
		l("")
		if statement.Id() != template.CommentToken &&
			statement.Id() != template.TextToken {

			l("// %s", statement.Loc())
		}

		switch stmt := statement.(type) {
		case *template.Atom:
			switch stmt.Id() {
			case template.CommentToken:
				// do nothing
			case template.TextToken:
				if stmt.Value != "" {
					l("// %s", statement.Loc())
					l("{")
					push()
					// Note: This assumes the value is already golang string
					// escaped
					l("_n, _err := _output.Write([]byte(`%s`))", stmt.Value)
					l("_numWritten += int64(_n)")
					l("if _err != nil {")
					l("    return _numWritten, _err")
					l("}")
					pop()
					l("}")
				}
			case template.SubstitutionToken:
				l("{")
				push()
				l("_n, _err := _template.writeValue(")
				l("    _output,")
				l("    (%s),", stmt.Value)
				l("    \"%s\")", stmt.Loc())
				l("_numWritten += int64(_n)")
				l("if _err != nil {")
				l("    return _numWritten, _err")
				l("}")
				pop()
				l("}")
			case template.EmbedToken:
				l("{")
				push()
				l("_n, _err := (%s).WriteTo(_output)", stmt.Value)
				l("_numWritten += _n")
				l("if _err != nil {")
				l("    return _numWritten, _err")
				l("}")
				pop()
				l("}")
			case template.CopySectionToken:
				l("%s", stmt.Value)
			case template.ContinueToken:
				l("continue")
			case template.BreakToken:
				l("break")
			case template.ReturnToken:
				l("return _numWritten, nil")
			case template.ErrorToken:
				l("{")
				push()
				l("_err := (%s)", stmt.Value)
				l("if _err == nil {")
				l("    _err = %v(\"Unexpected error (%s)\")",
					imports.Obj("fmt.Errorf"),
					stmt.Loc())
				l("}")
				l("return _numWritten, _err")
				pop()
				l("}")
			default:
				l("COMPILE ERROR: Bug in template generation code")
				l("Unexpected atom type: %s", stmt.Id())
			}
		case *template.For:
			l("for %s {", stmt.Predicate.Value)
			buildBody(stmt.Body, builder, imports)
			l("}")
		case *template.Switch:
			l("switch %s {", stmt.Switch.Value)
			for _, branch := range stmt.Cases {
				l("case %s:", branch.Predicate.Value)
				buildBody(branch.Body, builder, imports)
			}

			if stmt.Default != nil {
				l("default:")
				buildBody(stmt.Default.Body, builder, imports)
			}
			l("}")
		case *template.If:
			l("if %s {", stmt.If.Predicate.Value)
			buildBody(stmt.If.Body, builder, imports)
			for _, branch := range stmt.ElseIfs {
				l("} else if %s {", branch.Predicate.Value)
				buildBody(branch.Body, builder, imports)
			}
			if stmt.Else != nil {
				l("} else {")
				buildBody(stmt.Else.Body, builder, imports)
			}
			l("}")
		default:
			l("{")
			l("    // compile error.  bug in template generation code")
			l("    Unexpected statement type: %s", stmt.Id())
			l("}")
		}
	}
}
