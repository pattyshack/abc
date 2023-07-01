package py_cst

import (
	"fmt"
)

type CommentsFormatting struct {
	// Same idea as go's comment map:
	// A comment group g is associated with a node n if:
	// - g starts on the same line as n ends
	// - g starts on the line immediately following n, and there is
	//   at least one empty line after g and before the next node
	// - g starts before n and is not associated to the node before n
	//   via the previous rules
	LeadingCommentGroups     [][]string
	TrailingCommentGroups    [][]string
	TrailingNewLine          bool
	TrailingLineContinuation bool
}

type Node struct {
	Location

	CommentsFormatting
}

func (n *Node) MergeFrom(other *Node) {
	if n.Location.Line == 0 { // i.e., uninitialized
		n.Location = other.Location
	}

	if len(other.LeadingCommentGroups) > 0 ||
		len(other.TrailingCommentGroups) > 0 {
		n.LeadingCommentGroups = append(
			n.LeadingCommentGroups,
			n.TrailingCommentGroups...)

		n.LeadingCommentGroups = append(
			n.LeadingCommentGroups,
			other.LeadingCommentGroups...)

		n.TrailingCommentGroups = other.TrailingCommentGroups

		other.LeadingCommentGroups = nil
		other.TrailingCommentGroups = nil
	}

	n.TrailingNewLine = other.TrailingNewLine
	n.TrailingLineContinuation = other.TrailingLineContinuation
}

func (n *Node) MergeTrailingFrom(other *Node) {
	if n.Location.Line == 0 { // i.e., uninitialized
		n.Location = other.Location
	}

	n.TrailingCommentGroups = append(
		n.TrailingCommentGroups,
		other.LeadingCommentGroups...)

	n.TrailingCommentGroups = append(
		n.TrailingCommentGroups,
		other.TrailingCommentGroups...)

	other.LeadingCommentGroups = nil
	other.TrailingCommentGroups = nil

	n.TrailingNewLine = other.TrailingNewLine
	n.TrailingLineContinuation = other.TrailingLineContinuation
}

func (n *Node) MergeLeadingFrom(other *Node) {
	n.Location = other.Location

	leading := other.LeadingCommentGroups
	leading = append(leading, other.TrailingCommentGroups...)

	other.LeadingCommentGroups = nil
	other.TrailingCommentGroups = nil

	n.LeadingCommentGroups = append(leading, n.LeadingCommentGroups...)
}

func (n *Node) NodeInfo() *Node {
	return n
}

func (n *Node) str(prefix string) string {
	v := n.Location.ShortString()
	for _, g := range n.LeadingCommentGroups {
		v += "\n" + prefix + "  LeadingCommentGroup:"
		for _, l := range g {
			v += "\n" + prefix + "    " + l
		}
	}

	for _, g := range n.TrailingCommentGroups {
		v += "\n" + prefix + "  TrailingCommentGroup:"
		for _, l := range g {
			v += "\n" + prefix + "    " + l
		}
	}

	if n.TrailingNewLine {
		v += "\n" + prefix + "  TrailingNewLine: true"
	}

	if n.TrailingLineContinuation {
		v += "\n" + prefix + "  TrailingLineContinuation: true"
	}

	return v
}

func (n *Node) String() string {
	return n.str("")
}

// TODO: Add Loc, comment, extra
type NodeInfo interface {
	str(string) string

	String() string

	NodeInfo() *Node
}

type Token struct {
	Node

	TokenType int
	Value     string
}

func (t *Token) str(prefix string) string {
	// HACK: go yacc token id generation seems broken
	val := fmt.Sprintf("%vToken %v: %v",
		prefix,
		yyTokname(t.TokenType-57342),
		t.Node.str(prefix))
	val += "\n" + prefix + "    Value: " + t.Value

	return val
}

func (t *Token) String() string {
	return t.str("")
}

// Intermediate node.  Do not appear in the parsed tree.
// there are many ways to interpret exprlist / testlist / etc =...(
// (i) and (i,) have different meanings
type ExprList struct {
	Expressions        []Expression
	ExplicitCollection bool
}

func NewExprList(list []Expression) *ExprList {
	v := &ExprList{
		Expressions: list,
	}

	return v
}

func (e *ExprList) ConvertToExpr() Expression {
	if len(e.Expressions) == 1 && !e.ExplicitCollection {
		return e.Expressions[0]
	}

	return &CollectionExpr{
		Type:  TupleCollection,
		Items: e.Expressions,
	}
}

type Argument struct {
	Node

	// Optional when used by CallExpr
	Name Expression

	// Optional when used by FuncDef / LambdaExpr
	Value Expression

	PositionVarArg bool // '*'
	KeywordVarArg  bool // '**'
}

func NewArgument(name Expression, assign *Token, value Expression) *Argument {
	arg := &Argument{
		Name:  name,
		Value: value,
	}

	if name != nil {
		arg.MergeFrom(name.NodeInfo())
	}

	if assign != nil {
		arg.MergeFrom(&assign.Node)
	}

	if value != nil {
		arg.MergeFrom(value.NodeInfo())
	}

	return arg
}

// *arg in the context of definitions.
func NewPositionVarParam(star *Token, name Expression) *Argument {
	arg := &Argument{
		Node:           star.Node,
		Name:           name,
		PositionVarArg: true,
	}

	arg.MergeFrom(name.NodeInfo())
	return arg
}

// **kwarg in the context of defintions.
func NewKeywordVarParam(starStar *Token, name Expression) *Argument {
	arg := &Argument{
		Node:          starStar.Node,
		Name:          name,
		KeywordVarArg: true,
	}

	arg.MergeFrom(name.NodeInfo())
	return arg
}

// *arg in the context of invocation.
func NewPositionVarArg(star *Token, value Expression) *Argument {
	arg := &Argument{
		Node:           star.Node,
		Value:          value,
		PositionVarArg: true,
	}

	arg.MergeFrom(value.NodeInfo())
	return arg
}

// **kwarg in the context of invocation.
func NewKeywordVarArg(starStar *Token, value Expression) *Argument {
	arg := &Argument{
		Node:          starStar.Node,
		Value:         value,
		KeywordVarArg: true,
	}

	arg.MergeFrom(value.NodeInfo())
	return arg
}

func (a *Argument) str(prefix string) string {
	val := prefix + "Argument:" + a.Node.str(prefix) + "\n"
	if a.Name != nil {
		val += prefix + "  Name:\n"
		val += a.Name.str(prefix+"    ") + "\n"
	}

	val += prefix + "  Value:\n"
	if a.Value != nil {
		val += a.Value.str(prefix + "    ")
	} else {
		val += prefix + "    unknown"
	}

	return val
}

func (a *Argument) String() string {
	return a.str("")
}

// intermediate node.  does not appear in parsed tree.
type ArgumentList struct {
	Node
	Args []*Argument
}

func NewArgumentList(
	leftParen *Token,
	args []*Argument,
	rightParen *Token) *ArgumentList {

	list := &ArgumentList{
		Node: leftParen.Node,
		Args: args,
	}

	if len(list.TrailingCommentGroups) > 0 && len(args) > 0 {
		args[0].LeadingCommentGroups = append(
			list.TrailingCommentGroups,
			args[0].LeadingCommentGroups...)

		list.TrailingCommentGroups = nil
	}

	if len(rightParen.LeadingCommentGroups) > 0 {
		if len(args) > 0 {
			args[len(args)-1].TrailingCommentGroups = append(
				args[len(args)-1].TrailingCommentGroups,
				rightParen.LeadingCommentGroups...)
		} else {
			list.TrailingCommentGroups = append(
				list.TrailingCommentGroups,
				rightParen.LeadingCommentGroups...)
		}
	}

	list.TrailingCommentGroups = append(
		list.TrailingCommentGroups,
		rightParen.TrailingCommentGroups...)

	return list
}

// Used by both for loops and comprehensions
// for <binded variables> in <source> if <filter>
type Iterator struct {
	Node

	BoundVariables Expression
	Source         Expression
	Filters        []Expression // Optional
}

func (i *Iterator) str(prefix string) string {
	v := prefix + "Iterator:" + i.Node.str(prefix)
	v += "\n" + prefix + "  BoundVariables:\n"
	v += i.BoundVariables.str(prefix + "    ")

	v += "\n" + prefix + "  Source:\n"
	v += i.Source.str(prefix + "    ")

	if len(i.Filters) > 0 {
		v += "\n" + prefix + "  Filters:\n"
		for _, f := range i.Filters {
			v += f.str(prefix + "    ")
		}
	}

	return v
}

func (i *Iterator) String() string {
	return i.str("")
}

type ImportClause struct {
	Node

	Name  *Token
	Alias *Token
}

func (i *ImportClause) str(prefix string) string {
	v := prefix + "ImportClause:\n" + prefix + "  Name:\n"
	v += i.Name.str(prefix + "    ")

	if i.Alias != nil {
		v += "\n  Alias:\n" + i.Alias.str(prefix+"    ")
	}

	return v
}

func (i *ImportClause) String() string {
	return i.str("")
}

type ConditionClause struct {
	// predicate for if/while stmt, exceptions for except clauses.
	// nil Matches in except clause means unconditional matching.
	Matches Expression

	Alias Expression // Optional.  Only used by except clauses

	Branch []Statement
}

func (c *ConditionClause) str(prefix string) string {
	v := prefix + "ConditionClause:"
	if c.Matches != nil {
		v += "\n" + prefix + "  Matches:\n" + c.Matches.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Branch:"
	for _, s := range c.Branch {
		v += "\n" + s.str(prefix+"    ")
	}

	return v
}

func (c *ConditionClause) String() string {
	return c.str("")
}

// This is a "union" of Ellipsis, Index, and (Left:Middle:Right)
type Subscript struct {
	// dot dot dot
	Ellipsis bool // When true, Index/Left/Middle/Right must be nil

	// Exact index.
	Index Expression

	// <left>:<middle>:<right>
	Left   Expression
	Middle Expression
	Right  Expression
}

func (s *Subscript) str(prefix string) string {
	v := prefix + "Subscript:\n"
	if s.Ellipsis {
		v += prefix + "  Ellipsis: true"
	} else if s.Index != nil {
		v += prefix + "  Index:\n" + s.Index.str(prefix+"    ")
	} else {
		v += prefix + "  Left:\n"
		if s.Left == nil {
			v += " <nil>"
		} else {
			v += "\n" + s.Left.str(prefix+"    ")
		}

		v += "\n" + prefix + "  Middle:"
		if s.Middle == nil {
			v += " <nil>"
		} else {
			v += "\n" + s.Middle.str(prefix+"    ")
		}

		v += "\n" + prefix + "  Middle:"
		if s.Middle == nil {
			v += " <nil>"
		} else {
			v += "\n" + s.Middle.str(prefix+"    ")
		}
	}

	return v
}

func (s *Subscript) String() string {
	return s.str("")
}

type WithClause struct {
	Value         Expression
	BoundVariable Expression
}

func (w *WithClause) str(prefix string) string {
	v := prefix + "WithClause:\n" + prefix + "  Value:\n"
	v += w.Value.str(prefix + "    ")
	if w.BoundVariable != nil {
		v += "\n" + prefix + "  BoundVariables:\n"
		v += w.BoundVariable.str(prefix + "    ")
	}
	return v
}

func (w *WithClause) String() string {
	return w.str("")
}
