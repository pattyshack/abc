package py_cst

import (
	"fmt"
)

type Statement interface {
	NodeInfo

	IsStatement()
}

type isStmt struct {
	Node
}

func (isStmt) IsStatement() {}

type ExprStmt struct {
	isStmt

	Expression
}

func NewExprStmt(expr Expression) *ExprStmt {
	stmt := &ExprStmt{
		Expression: expr,
	}

	stmt.MergeFrom(expr.NodeInfo())

	return stmt
}

func (e *ExprStmt) str(prefix string) string {
	v := prefix + "ExprStmt: " + e.Node.str(prefix)
	v += "\n" + e.Expression.str(prefix+"    ")
	return v
}

func (e *ExprStmt) String() string {
	return e.str("")
}

type PassStmt struct {
	isStmt

	// newlines and comment blocks are implicitly pass statments
	IsImplicit bool
}

func NewPassStmt(token *Token, implicit bool) *PassStmt {
	s := &PassStmt{
		IsImplicit: implicit,
	}

	s.Node = token.Node

	return s
}

func (p *PassStmt) str(prefix string) string {
	return fmt.Sprintf(
		"%sPassStmt: %s\n%s  Implicit: %v",
		prefix,
		p.Node.str(prefix),
		prefix,
		p.IsImplicit)
}

func (p *PassStmt) String() string {
	return p.str("")
}

// TODO doc string ...
type FuncDef struct {
	isStmt

	Decorators []*Decorator

	Name       string
	Arguments  []*Argument
	Statements []Statement
}

func NewFuncDef(
	def *Token,
	name *Token,
	args *ArgumentList,
	colon *Token,
	stmts []Statement) *FuncDef {

	f := &FuncDef{
		Name:       name.Value,
		Arguments:  args.Args,
		Statements: stmts,
	}

	f.Node = def.Node
	f.MergeFrom(&name.Node)
	f.MergeFrom(&args.Node)
	f.MergeTrailingFrom(&colon.Node)

	return f
}

func (f *FuncDef) SetDecorators(decorators []*Decorator) {
	f.Decorators = decorators

	for i := len(decorators) - 1; i >= 0; i-- {
		f.MergeLeadingFrom(decorators[i].NodeInfo())
	}
}

func (f *FuncDef) str(prefix string) string {
	v := prefix + "FuncDef: " + f.Node.str(prefix)
	v += "\n" + prefix + "  Name: " + f.Name

	v += "\n" + prefix + "  Decorators:"
	for _, d := range f.Decorators {
		v += "\n" + d.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Arguments:"
	for _, a := range f.Arguments {
		v += "\n" + a.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Statements:"
	for _, s := range f.Statements {
		v += "\n" + s.str(prefix+"    ")
	}

	return v
}

func (f *FuncDef) String() string {
	return f.str("")
}

// TODO doc string ...
type ClassDef struct {
	isStmt

	Decorators []*Decorator

	Name          *Token
	ParentClasses []Expression
	Statements    []Statement
}

func (c *ClassDef) SetDecorators(decorators []*Decorator) {
	c.Decorators = decorators

	for i := len(decorators) - 1; i >= 0; i-- {
		c.MergeLeadingFrom(decorators[i].NodeInfo())
	}
}

func (c *ClassDef) str(prefix string) string {
	v := prefix + "ClassDef: " + c.Node.str(prefix)
	v += "\n" + prefix + "  Name:\n"
	v += c.Name.str(prefix + "    ")

	v += "\n" + prefix + "  Decorators:"
	for _, d := range c.Decorators {
		v += "\n" + d.str(prefix+"    ")
	}

	v += "\n" + prefix + "  ParentClasses:"
	for _, p := range c.ParentClasses {
		v += "\n" + p.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Statements:"
	for _, s := range c.Statements {
		v += "\n" + s.str(prefix+"    ")
	}

	return v
}

func (c *ClassDef) String() string {
	return c.str("")
}

type GlobalStmt struct {
	isStmt

	Names []*Token
}

func (g *GlobalStmt) str(prefix string) string {
	v := prefix + "GlobalStmt: " + g.Node.str(prefix)
	v += "\n" + prefix + "  Names:"

	for _, n := range g.Names {
		v += "\n" + n.str(prefix+"    ")
	}

	return v
}

func (g *GlobalStmt) String() string {
	return g.str("")
}

type ImportStmt struct {
	isStmt

	ModulePath []*Token
	Alias      *Token // Optional
}

func (i *ImportStmt) str(prefix string) string {
	v := prefix + "ImportStmt: " + i.Node.str(prefix)
	v += "\n" + prefix + "  ModulePath:"

	for _, p := range i.ModulePath {
		v += "\n" + p.str(prefix+"    ")
	}

	if i.Alias != nil {
		v += "\n" + prefix + "  Alias:\n"
		v += i.Alias.str(prefix + "    ")
	}
	return v
}

func (i *ImportStmt) String() string {
	return i.str("")
}

type FromStmt struct {
	isStmt

	DotPrefixCount int
	ModulePath     []*Token

	// Empty list implies "import *"
	Imports []*ImportClause
}

func (f *FromStmt) str(prefix string) string {
	v := fmt.Sprintf(
		"%sFromStmt: %s\n%s  DotPrefixCount: %v\n%sModulePath:",
		prefix,
		f.Node.str(prefix),
		prefix,
		f.DotPrefixCount,
		prefix)

	for _, p := range f.ModulePath {
		v += "\n" + p.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Imports:"
	if len(f.Imports) == 0 {
		v += " ALL"
	} else {
		for _, i := range f.Imports {
			v += "\n" + i.str(prefix+"    ")
		}
	}

	return v
}

func (f *FromStmt) String() string {
	return f.str("")
}

type PrintStmt struct {
	isStmt

	Output           Expression   // Optional
	Values           []Expression // Optional
	LineContinuation bool
}

func NewPrintStmt(
	prnt *Token,
	shift *Token,
	out Expression,
	values *ExprList) *PrintStmt {

	stmt := &PrintStmt{
		Output: out,
	}

	if values != nil {
		stmt.Values = values.Expressions
		stmt.LineContinuation = values.ExplicitCollection
	}

	stmt.MergeFrom(&prnt.Node)

	if shift != nil {
		stmt.MergeFrom(&shift.Node)
	}

	if out != nil {
		stmt.MergeFrom(out.NodeInfo())
	}

	return stmt
}

func (p *PrintStmt) str(prefix string) string {
	v := prefix + "PrintStmt: " + p.Node.str(prefix)
	v += "\n" + prefix + "  Values:"
	for _, val := range p.Values {
		v += "\n" + val.str(prefix+"    ")
	}
	if p.LineContinuation {
		v += "\n" + prefix + "  LineContinuation: true"
	}

	return v
}

func (p *PrintStmt) String() string {
	return p.str("")
}

// Base type for break & continue
type TokenStmt struct {
	isStmt

	DebugType string
}

func NewTokenStmt(t *Token, debugType string) TokenStmt {
	stmt := TokenStmt{
		DebugType: debugType,
	}

	stmt.Node = t.Node
	return stmt
}

func (t *TokenStmt) str(prefix string) string {
	v := prefix + t.DebugType + ": " + t.Node.str(prefix)
	return v
}

func (t *TokenStmt) String() string {
	return t.str("")
}

type BreakStmt struct {
	TokenStmt
}

func NewBreakStmt(t *Token) *BreakStmt {
	return &BreakStmt{NewTokenStmt(t, "BreakStmt")}
}

type ContinueStmt struct {
	TokenStmt
}

func NewContinueStmt(t *Token) *ContinueStmt {
	return &ContinueStmt{NewTokenStmt(t, "ContinueStmt")}
}

type ConditionStmt struct {
	isStmt

	If   *ConditionClause
	Elif []*ConditionClause // Optional
	Else []Statement        // Optional
}

func (c *ConditionStmt) str(prefix string) string {
	v := prefix + "ConditionStmt: " + c.Node.str(prefix)
	v += "\n" + prefix + "  If:\n"
	v += c.If.str(prefix + "    ")

	if len(c.Elif) > 0 {
		v += "\n" + prefix + "  Elif:\n"
		for _, clause := range c.Elif {
			v += "\n" + clause.str(prefix+"    ")
		}
	}

	if len(c.Else) > 0 {
		v += "\n" + prefix + "  Else:\n"
		for _, s := range c.Else {
			v += "\n" + s.str(prefix+"    ")
		}
	}

	return v
}

func (c *ConditionStmt) String() string {
	return c.str("")
}

type WhileStmt struct {
	isStmt

	Loop *ConditionClause

	Else []Statement // Optional
}

func (w *WhileStmt) str(prefix string) string {
	v := prefix + "WhileStmt: " + w.Node.str(prefix)
	v += "\n" + prefix + "  Loop:\n"
	v += w.Loop.str(prefix + "    ")

	if len(w.Else) > 0 {
		v += "\n" + prefix + "  Else:"
		for _, s := range w.Else {
			v += "\n" + s.str(prefix+"    ")
		}
	}

	return v
}

func (w *WhileStmt) String() string {
	return w.str("")
}

type ForStmt struct {
	isStmt

	Iterator *Iterator

	Loop []Statement
	Else []Statement // Optional
}

func (f *ForStmt) str(prefix string) string {
	v := prefix + "ForStmt: " + f.Node.str(prefix)
	v += "\n" + prefix + "  Iterator:\n"
	v += f.Iterator.str(prefix + "    ")

	v += "\n" + prefix + "  Loop:"
	for _, s := range f.Loop {
		v += "\n" + s.str(prefix+"    ")
	}

	if len(f.Else) > 0 {
		v += "\n" + prefix + "  Else:"
		for _, s := range f.Else {
			v += "\n" + s.str(prefix+"    ")
		}
	}

	return v
}

func (f *ForStmt) String() string {
	return f.str("")
}

type TryStmt struct {
	isStmt

	Try     []Statement
	Except  []*ConditionClause // Optional
	Else    []Statement        // Optional
	Finally []Statement        // Optional
}

func (t *TryStmt) str(prefix string) string {
	v := prefix + "TryStmt: " + t.Node.str(prefix)
	v += "\n" + prefix + "  Try:"
	for _, s := range t.Try {
		v += "\n" + s.str(prefix+"    ")
	}

	if len(t.Except) > 0 {
		v += "\n" + prefix + "  Except:"
		for _, c := range t.Except {
			v += "\n" + c.str(prefix+"    ")
		}
	}

	if len(t.Else) > 0 {
		v += "\n" + prefix + "  Else:"
		for _, s := range t.Else {
			v += "\n" + s.str(prefix+"    ")
		}
	}

	if len(t.Finally) > 0 {
		v += "\n" + prefix + "  Finally:"
		for _, s := range t.Finally {
			v += "\n" + s.str(prefix+"    ")
		}
	}

	return v
}

func (t *TryStmt) String() string {
	return t.str("")
}

type ReturnStmt struct {
	isStmt

	Value Expression
}

func NewReturnStmt(ret *Token, value Expression) *ReturnStmt {
	stmt := &ReturnStmt{
		Value: value,
	}

	stmt.Node = ret.Node

	if value != nil {
		stmt.MergeFrom(value.NodeInfo())
	} else {
		stmt.Value = NewNone(ret)
	}

	return stmt
}

func (r *ReturnStmt) str(prefix string) string {
	v := prefix + "ReturnStmt: " + r.Node.str(prefix)
	v += "\n" + prefix + "  Value:\n"
	v += r.Value.str(prefix + "    ")
	return v
}

func (r *ReturnStmt) String() string {
	return r.str("")
}

type DelStmt struct {
	isStmt

	List []Expression
}

func NewDelStmt(del *Token, list *ExprList) *DelStmt {
	stmt := &DelStmt{
		List: list.Expressions,
	}
	stmt.Node = del.Node

	return stmt
}

func (d *DelStmt) str(prefix string) string {
	v := prefix + "DelStmt: " + d.Node.str(prefix)
	v += "\n" + prefix + "  List:"
	for _, e := range d.List {
		v += "\n" + e.str(prefix+"    ")
	}

	return v
}

func (d *DelStmt) String() string {
	return d.str("")
}

type RaiseStmt struct {
	isStmt

	First  Expression
	Second Expression
	Third  Expression
}

func NewRaiseStmt(
	r *Token,
	first Expression,
	comma1 *Token,
	second Expression,
	comma2 *Token,
	third Expression) *RaiseStmt {

	stmt := &RaiseStmt{
		First:  first,
		Second: second,
		Third:  third,
	}

	stmt.Node = r.Node

	if first != nil {
		stmt.MergeFrom(first.NodeInfo())
	}

	if comma1 != nil {
		stmt.MergeFrom(&comma1.Node)
	}

	if second != nil {
		stmt.MergeFrom(second.NodeInfo())
	}

	if comma2 != nil {
		stmt.MergeFrom(&comma2.Node)
	}

	if third != nil {
		stmt.MergeFrom(third.NodeInfo())
	}

	return stmt
}

func (r *RaiseStmt) str(prefix string) string {
	v := prefix + "RaiseStmt: " + r.Node.str(prefix)
	v += "\n" + prefix + "  First:\n"
	v += r.First.str(prefix + "    ")
	if r.Second != nil {
		v += "\n" + prefix + "  Second\n" + r.Second.str(prefix+"    ")
	}
	if r.Third != nil {
		v += "\n" + prefix + "  Third\n" + r.Third.str(prefix+"    ")
	}

	return v
}

func (r *RaiseStmt) String() string {
	return r.str("")
}

type ExecStmt struct {
	isStmt

	Expr   Expression
	Global Expression
	Local  Expression
}

func (e *ExecStmt) str(prefix string) string {
	v := prefix + "ExecStmt: " + e.Node.str(prefix)
	v += "\n" + prefix + "  Expr:\n"
	v += e.Expr.str(prefix + "    ")
	if e.Global != nil {
		v += "\n" + prefix + "  Global\n" + e.Global.str(prefix+"    ")
	}
	if e.Local != nil {
		v += "\n" + prefix + "  Local\n" + e.Local.str(prefix+"    ")
	}

	return v
}

func (e *ExecStmt) String() string {
	return e.str("")
}

type AssertStmt struct {
	isStmt

	Expr  Expression
	Debug Expression // Optional
}

func (a *AssertStmt) str(prefix string) string {
	v := prefix + "AssertStmt: " + a.Node.str(prefix)
	v += "\n" + prefix + "  Expr:\n"
	v += a.Expr.str(prefix + "    ")
	if a.Debug != nil {
		v += "\n" + prefix + "  Debug\n" + a.Debug.str(prefix+"    ")
	}

	return v
}

func (a *AssertStmt) String() string {
	return a.str("")
}

type WithStmt struct {
	isStmt

	WithClauses []*WithClause
	Statements  []Statement
}

func (w *WithStmt) str(prefix string) string {
	v := prefix + "WithStmt: " + w.Node.str(prefix)
	v += "\n" + prefix + "  WithClause:"

	for _, c := range w.WithClauses {
		v += "\n" + c.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Statement:"
	for _, s := range w.Statements {
		v += "\n" + s.str(prefix+"    ")
	}

	return v
}

func (w *WithStmt) String() string {
	return w.str("")
}
