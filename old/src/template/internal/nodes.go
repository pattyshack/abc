package template

// Note: in case of text, we need to modify the value and escape ` and /
// correctly

// in lexer, we can enforce scoping character.  If we prevent unpaired ), we
// can ensure that ($(stmt.Value)) always is an valid expression

type Statement interface {
	IsStatement()

	Id() SymbolId
	Loc() Location
}

type TToken struct {
	GenericSymbol

	// When true, and the previous statement is text, remove the whitespaces
	// in the text that are adjacent to this statement, potentially up to and
	// including the previous line's newline character.
	trimLeadingWhitespaces bool

	// When true, and the next statement is text, remove the whitespaces
	// in the text taht are adjacent to this statement, potentially up to and
	// including the current line's newline character.
	trimTrailingWhitespaces bool
}

func NewTToken(
	id SymbolId,
	loc Location,
	trimLeading bool,
	trimTrailing bool) *TToken {

	return &TToken{
		GenericSymbol{id, loc},
		trimLeading,
		trimTrailing,
	}
}

func (token *TToken) TrimLeadingWhitespaces() bool {
	return token.trimLeadingWhitespaces
}

func (token *TToken) TrimTrailingWhitespaces() bool {
	return token.trimTrailingWhitespaces
}

type Value struct {
	*TToken

	Value string
}

func NewValue(
	id SymbolId,
	loc Location,
	val string,
	trimLeading bool,
	trimTrailing bool) *Value {

	return &Value{NewTToken(id, loc, trimLeading, trimTrailing), val}
}

type Atom struct {
	*TToken

	Value string
}

func NewAtom(
	id SymbolId,
	loc Location,
	val string,
	trimLeading bool,
	trimTrailing bool) *Atom {

	return &Atom{NewTToken(id, loc, trimLeading, trimTrailing), val}
}

func (Atom) IsStatement() {}

type Branch struct {
	Predicate *Value
	Body      []Statement
}

type For struct {
	Branch
}

func (For) IsStatement() {}

func (For) Id() SymbolId { return ForType }

func (f *For) Loc() Location {
	return f.Predicate.Loc()
}

type Switch struct {
	Switch  *Value
	Cases   []*Branch
	Default *Branch
}

func (Switch) IsStatement() {}

func (Switch) Id() SymbolId { return SwitchType }

func (s *Switch) Loc() Location {
	return s.Switch.Loc()
}

type If struct {
	If      Branch
	ElseIfs []*Branch
	Else    *Branch
}

func (If) IsStatement() {}

func (If) Id() SymbolId { return IfType }

func (i *If) Loc() Location {
	return i.If.Predicate.Loc()
}

type Argument struct {
	Name string
	Type string
}

type TemplateDeclaration struct {
	GenericSymbol

	TemplateName string
	Arguments    []Argument
}

func NewTemplateDeclaration(
	loc Location,
	name string,
	args []Argument) *TemplateDeclaration {

	return &TemplateDeclaration{
		GenericSymbol{TemplateDeclToken, loc},
		name,
		args,
	}
}

type File struct {
	PackageName string

	Imports string

	*TemplateDeclaration

	Body []Statement
}
