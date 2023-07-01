package py_cst

type Expression interface {
	NodeInfo

	IsExpression()
}

type isExpr struct {
	Node
}

func (isExpr) IsExpression() {}

func DottedNameToExpr(names []*Token) Expression {
	var expr Expression
	for i, n := range names {
		if i == 0 {
			expr = NewIdentifier(n)
		} else {
			expr = &DotExpr{
				Parent: expr,
				Field:  n,
			}
		}
	}

	return expr
}

func NewFactorExpr(factorSigns []*Token, expr Expression) Expression {
	for i := len(factorSigns) - 1; i >= 0; i-- {
		expr = &UnaryExpr{
			Op:    factorSigns[i],
			Value: expr,
		}
	}

	return expr
}

type ComprehensionType int

const (
	UnknownComprehension                     = 0
	GeneratorComprehension ComprehensionType = 1
	ListComprehension      ComprehensionType = 2
	SetComprehension       ComprehensionType = 3
	DictComprehension      ComprehensionType = 4
)

type ComprehensionExpr struct {
	isExpr

	Type      ComprehensionType
	Key       Expression // Set by dict / set comprehension
	Value     Expression // Set by list / generator / dict comprehension
	Iterators []*Iterator
}

func (c *ComprehensionExpr) str(prefix string) string {
	val := prefix + "Comprehension:" + c.Node.str(prefix)
	val += "\n" + prefix + "  Type: "
	switch c.Type {
	case UnknownComprehension:
		val += "unknown"
	case GeneratorComprehension:
		val += "generator"
	case ListComprehension:
		val += "list"
	case SetComprehension:
		val += "set"
	case DictComprehension:
		val += "dict"
	}

	if c.Key != nil {
		val += "\n" + prefix + "  Key:\n"
		val += c.Key.str(prefix + "    ")
	}

	if c.Value != nil {
		val += "\n" + prefix + "  Value:\n"
		val += c.Value.str(prefix + "    ")
	}

	val += "\n" + prefix + "  Iterators:"
	for _, iter := range c.Iterators {
		val += "\n" + iter.str(prefix+"    ")
	}

	return val
}

func (c *ComprehensionExpr) String() string {
	return c.str("")
}

type CollectionType int

const (
	UnknownCollection = 0
	TupleCollection   = 1
	ListCollection    = 2
	SetCollection     = 3
	DictCollection    = 4
)

type CollectionExpr struct {
	isExpr

	Type  CollectionType
	Items []Expression
}

func (c *CollectionExpr) str(prefix string) string {
	v := prefix + "Collection: " + c.Node.str(prefix)
	v += "\n" + prefix + "  Type: "
	switch c.Type {
	case UnknownCollection:
		v += "unknown"
	case TupleCollection:
		v += "tuple"
	case ListCollection:
		v += "list"
	case SetCollection:
		v += "set"
	case DictCollection:
		v += "dict"
	}

	v += "\n" + prefix + "  Items:"
	for _, item := range c.Items {
		v += "\n" + item.str(prefix+"    ")
	}

	return v
}

func (c *CollectionExpr) String() string {
	return c.str("")
}

type UnaryExpr struct {
	isExpr

	Op    *Token
	Value Expression
}

func (u *UnaryExpr) str(prefix string) string {
	v := prefix + "Unary: " + u.Node.str(prefix)
	v += "\n" + prefix + "  Op:\n"
	v += u.Op.str(prefix + "    ")

	v += "\n" + prefix + "  Value:\n"
	v += u.Value.str(prefix + "    ")

	return v
}

func (u *UnaryExpr) String() string {
	return u.str("")
}

type BinaryExpr struct {
	isExpr

	Left  Expression
	Op    string
	Right Expression
}

func NewBinaryExpr(lhs Expression, op *Token, rhs Expression) *BinaryExpr {
	expr := &BinaryExpr{
		Left:  lhs,
		Op:    op.Value,
		Right: rhs,
	}

	expr.MergeFrom(lhs.NodeInfo())
	expr.MergeFrom(&op.Node)
	expr.MergeFrom(rhs.NodeInfo())

	return expr
}

func (b *BinaryExpr) str(prefix string) string {
	v := prefix + "Binary: " + b.Node.str(prefix)
	v += "\n" + prefix + "  Left:\n"
	v += b.Left.str(prefix + "    ")

	v += "\n" + prefix + "  Op: " + b.Op

	v += "\n" + prefix + "  Right:\n"
	v += b.Right.str(prefix + "    ")

	return v
}

func (b *BinaryExpr) String() string {
	return b.str("")
}

// Base type for None, Number, and Identifer.
type TokenExpr struct {
	isExpr

	DebugType string

	Value string
}

func NewTokenExpr(t *Token, debugType string) TokenExpr {
	expr := TokenExpr{
		DebugType: debugType,
		Value:     t.Value,
	}

	expr.Node = t.Node
	return expr
}

func (t *TokenExpr) str(prefix string) string {
	v := prefix + t.DebugType + ": " + t.Node.str(prefix)
	v += "\n" + prefix + "  Value: " + t.Value
	return v
}

func (t *TokenExpr) String() string {
	return t.str("")
}

type None struct {
	TokenExpr
}

func NewNone(t *Token) *None {
	return &None{NewTokenExpr(t, "None")}
}

type Number struct {
	TokenExpr
}

func NewNumber(t *Token) *Number {
	return &Number{NewTokenExpr(t, "Number")}
}

type Identifier struct {
	TokenExpr
}

func NewIdentifier(t *Token) *Identifier {
	return &Identifier{NewTokenExpr(t, "Identifier")}
}

type String struct {
	isExpr

	Pieces []*Token
}

func (s *String) str(prefix string) string {
	v := prefix + "String: " + s.Node.str(prefix)
	v += "\n" + prefix + "  Pieces:"
	for _, p := range s.Pieces {
		v += "\n" + p.str(prefix+"    ")
	}

	return v
}

func (s *String) String() string {
	return s.str("")
}

// <parent>.<field>
type DotExpr struct {
	isExpr

	Parent Expression
	Field  *Token
}

func (f *DotExpr) str(prefix string) string {
	v := prefix + "DotExpr: " + f.Node.str(prefix)
	v += "\n" + prefix + "  Parent:\n"
	v += f.Parent.str(prefix + "    ")

	v += "\n" + prefix + "  Field:\n"
	v += f.Field.str(prefix + "    ")
	return v
}

func (f *DotExpr) String() string {
	return f.str("")
}

// <method>(<args>)
type CallExpr struct {
	isExpr

	Method    Expression
	Arguments []*Argument
}

func NewCallExpr(method Expression, args *ArgumentList) *CallExpr {
	expr := &CallExpr{
		Method:    method,
		Arguments: args.Args,
	}

	expr.MergeFrom(method.NodeInfo())
	expr.MergeFrom(&args.Node)

	return expr
}

func (i *CallExpr) str(prefix string) string {
	v := prefix + "CallExpr: " + i.Node.str(prefix)
	v += "\n" + prefix + "  Method:\n"
	v += i.Method.str(prefix + "    ")

	v += "\n" + prefix + "  Arguments:"
	for _, arg := range i.Arguments {
		v += "\n" + arg.str(prefix+"    ")
	}

	return v
}

func (i *CallExpr) String() string {
	return i.str("")
}

func NewExpressionFromTrailers(
	atom Expression,
	trailers []interface{}) Expression {

	expr := atom
	for _, t := range trailers {
		switch val := t.(type) {
		case *Token: // field access
			expr = &DotExpr{
				Parent: expr,
				Field:  val,
			}
		case []*Argument:
			// TODO
			expr = &CallExpr{
				Method:    expr,
				Arguments: val,
			}
		case []*Subscript:
			expr = &SubscriptExpr{
				Source:     expr,
				Subscripts: val,
			}
		default:
			panic("Unexpected trailer type")
		}
	}

	return expr
}

type EvalExpr struct {
	Expression
}

func (e *EvalExpr) str(prefix string) string {
	v := prefix + "Eval: " + e.NodeInfo().str(prefix)
	v += "\n" + prefix + "  Expression:\n"
	v += e.Expression.str(prefix + "    ")
	return v
}

func (e *EvalExpr) String() string {
	return e.str("")
}

type YieldExpr struct {
	Expression // nil means implicit None
}

func (y *YieldExpr) str(prefix string) string {
	v := prefix + "YieldExpr:" + y.NodeInfo().str(prefix)
	if y.Expression != nil {
		v += "\n" + prefix + "  Expression:\n"
		v += y.Expression.str(prefix + "    ")
	} else {
		v += " (Implicit None)"
	}

	return v
}

func (y *YieldExpr) String() string {
	return y.str("")
}

// Decorator is a pseudo-Expression
type Decorator struct {
	Expression
}

func NewDecorator(at *Token, expr Expression, newline *Token) *Decorator {
	d := &Decorator{
		Expression: expr,
	}

	d.NodeInfo().MergeLeadingFrom(&at.Node)
	d.NodeInfo().MergeTrailingFrom(&newline.Node)
	return d
}

func (d *Decorator) str(prefix string) string {
	v := prefix + "Decorator: " + d.NodeInfo().str(prefix)
	v += "\n" + prefix + "  Expression:\n"
	v += d.Expression.str(prefix + "    ")
	return v
}

func (d *Decorator) String() string {
	return d.str("")
}

type LambdaExpr struct {
	isExpr

	Arguments []*Argument
	Value     Expression
}

func (l *LambdaExpr) str(prefix string) string {
	v := prefix + "Lambda: " + l.Node.str(prefix)
	v += "\n" + prefix + "  Arguments:"

	for _, a := range l.Arguments {
		v += "\n" + a.str(prefix+"    ")
	}

	v += "\n" + prefix + "  Value:\n"
	v += l.Value.str(prefix + "    ")

	return v
}

func (l *LambdaExpr) String() string {
	return l.str("")
}

type ConditionExpr struct {
	isExpr

	Predicate Expression
	True      Expression
	False     Expression
}

func (c *ConditionExpr) str(prefix string) string {
	v := prefix + "ConditionExpr: " + c.Node.str(prefix)
	v += "\n" + prefix + "  Predicate:\n"
	v += c.Predicate.str(prefix + "    ")
	v += prefix + "  True:\n"
	v += c.True.str(prefix + "    ")
	v += prefix + "  False:\n"
	v += c.False.str(prefix + "    ")

	return v
}

func (c *ConditionExpr) String() string {
	return c.str("")
}

type SubscriptExpr struct {
	isExpr

	Source     Expression
	Subscripts []*Subscript
}

func (s *SubscriptExpr) str(prefix string) string {
	v := prefix + "SubscriptExpr " + s.Node.str(prefix)
	v += "\n" + prefix + "  Source:\n"
	v += s.Source.str(prefix + "    ")

	v += prefix + "  Subscripts:"
	for _, sub := range s.Subscripts {
		v += "\n" + sub.str(prefix+"    ")
	}

	return v
}

func (s *SubscriptExpr) String() string {
	return s.str("")
}
