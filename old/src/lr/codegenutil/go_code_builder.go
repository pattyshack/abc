package codegenutil

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"regexp"
)

var (
	nameRe = regexp.MustCompile(`((?:(?:\[\])*\**)*)(?:(.+)\.)?(\w+(?:{})?)$`)
)

type importEntry struct {
	path  string
	alias string
}

type importObject struct {
	*importEntry

	prefix string // * or &
	name   string
}

func (obj *importObject) String() string {
	if obj.alias == "" {
		return obj.prefix + obj.name
	}

	return obj.prefix + obj.alias + "." + obj.name
}

type GoImports struct {
	imports map[string]*importEntry
}

func NewGoImports() *GoImports {
	return &GoImports{
		imports: map[string]*importEntry{},
	}
}

// This supports accessing objects of the form:
//     *(\[\])*(\*)*)*(<full module path>\.)?<object>({})?
// map objects are not supported
func (imports *GoImports) Obj(fullName string) *importObject {
	match := nameRe.FindStringSubmatch(fullName)
	if match == nil {
		panic("Invalid fullName: " + fullName)
	}

	prefix := match[1]
	modulePath := match[2]
	name := match[3]

	if modulePath == "" {
		return &importObject{&importEntry{}, prefix, name}
	}

	entry, ok := imports.imports[modulePath]
	if !ok {
		entry = &importEntry{
			path: modulePath,
		}

		imports.imports[modulePath] = entry
	}

	return &importObject{entry, prefix, name}
}

func (imports *GoImports) assignAlias() error {
	aliasCount := map[string]int{}
	for pkg, entry := range imports.imports {
		base := filepath.Base(pkg)
		if base == "." || base == "/" {
			return fmt.Errorf("Invalid import path: %s", pkg)
		}

		alias := base
		aliasCount[alias] += 1
		if aliasCount[alias] > 1 {
			alias = fmt.Sprintf("%s%d", alias, aliasCount[alias])
		}

		entry.alias = alias
	}

	return nil
}

func (imports *GoImports) WriteTo(output io.Writer) (int64, error) {
	err := imports.assignAlias()
	if err != nil {
		return 0, err
	}

	builder := NewCodeBuilder()
	if len(imports.imports) > 0 {
		builder.Line("import (")
		builder.PushIndent()
		// Maybe separate built-in packages from other packages
		for _, entry := range imports.imports {
			builder.Line("%s \"%s\"", entry.alias, entry.path)
		}
		builder.PopIndent()
		builder.Line(")")
		builder.Line("")
	}

	return builder.WriteTo(output)
}

type FormattedGoSource struct {
	Source io.WriterTo
}

func NewFormattedGoSource(source io.WriterTo) io.WriterTo {
	return &FormattedGoSource{source}
}

func (fgs *FormattedGoSource) WriteTo(output io.Writer) (int64, error) {
	buffer := bytes.NewBuffer(nil)

	_, err := fgs.Source.WriteTo(buffer)
	if err != nil {
		return 0, err
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to format (%s) generated code:\n%s",
			err,
			buffer.Bytes())
	}

	n, err := output.Write(formatted)
	return int64(n), err
}
