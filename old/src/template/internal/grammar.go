// Auto-generated from source: grammar.lr

package template

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	SectionMarkerToken = SymbolId(256)
	PackageToken       = SymbolId(257)
	ImportToken        = SymbolId(258)
	TemplateDeclToken  = SymbolId(259)
	ForToken           = SymbolId(260)
	SwitchToken        = SymbolId(261)
	CaseToken          = SymbolId(262)
	IfToken            = SymbolId(263)
	ElseIfToken        = SymbolId(264)
	DefaultToken       = SymbolId(265)
	ElseToken          = SymbolId(266)
	EndToken           = SymbolId(267)
	TextToken          = SymbolId(268)
	SubstitutionToken  = SymbolId(269)
	EmbedToken         = SymbolId(270)
	CopySectionToken   = SymbolId(271)
	CommentToken       = SymbolId(272)
	ContinueToken      = SymbolId(273)
	BreakToken         = SymbolId(274)
	ReturnToken        = SymbolId(275)
	ErrorToken         = SymbolId(276)
)

type Location struct {
	FileName string
	Line     int
	Column   int
}

func (l Location) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l Location) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type Token interface {
	Id() SymbolId
	Loc() Location
}

type GenericSymbol struct {
	SymbolId
	Location
}

func (t *GenericSymbol) Id() SymbolId { return t.SymbolId }

func (t *GenericSymbol) Loc() Location { return t.Location }

type Lexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *GenericSymbol
	Next() (Token, error)

	CurrentLocation() Location
}

type Reducer interface {
	// 22:8: file -> ...
	ToFile(Package_ *Value, OptionalImports_ *Value, TemplateDecl_ *TemplateDeclaration, SectionMarker_ *GenericSymbol, Body_ []Statement) (*File, error)

	// 29:4: optional_imports -> imports: ...
	ImportsToOptionalImports(Import_ *Value) (*Value, error)

	// 30:4: optional_imports -> nil: ...
	NilToOptionalImports() (*Value, error)

	// 33:4: body -> add: ...
	AddToBody(Body_ []Statement, Statement_ Statement) ([]Statement, error)

	// 34:4: body -> nil: ...
	NilToBody() ([]Statement, error)

	// 37:4: statement -> atom: ...
	AtomToStatement(Atom_ Statement) (Statement, error)

	// 38:4: statement -> for: ...
	ForToStatement(For_ Statement) (Statement, error)

	// 39:4: statement -> switch: ...
	SwitchToStatement(Switch_ Statement) (Statement, error)

	// 40:4: statement -> if: ...
	IfToStatement(If_ Statement) (Statement, error)

	// 43:4: atom -> text: ...
	TextToAtom(Text_ *Atom) (Statement, error)

	// 44:4: atom -> substitution: ...
	SubstitutionToAtom(Substitution_ *Atom) (Statement, error)

	// 45:4: atom -> embed: ...
	EmbedToAtom(Embed_ *Atom) (Statement, error)

	// 46:4: atom -> copy_section: ...
	CopySectionToAtom(CopySection_ *Atom) (Statement, error)

	// 47:4: atom -> comment: ...
	CommentToAtom(Comment_ *Atom) (Statement, error)

	// 48:4: atom -> continue: ...
	ContinueToAtom(Continue_ *Atom) (Statement, error)

	// 49:4: atom -> break: ...
	BreakToAtom(Break_ *Atom) (Statement, error)

	// 50:4: atom -> return: ...
	ReturnToAtom(Return_ *Atom) (Statement, error)

	// 51:4: atom -> error: ...
	ErrorToAtom(Error_ *Atom) (Statement, error)

	// 53:7: for -> ...
	ToFor(For_ *Value, Body_ []Statement, End_ *TToken) (Statement, error)

	// 57:4: switch -> with_whitespace: ...
	WithWhitespaceToSwitch(Switch_ *Value, Text_ *Atom, CaseList_ []*Branch, OptionalDefault_ *Branch, End_ *TToken) (Statement, error)

	// 58:4: switch -> without_whitespace: ...
	WithoutWhitespaceToSwitch(Switch_ *Value, CaseList_ []*Branch, OptionalDefault_ *Branch, End_ *TToken) (Statement, error)

	// 61:4: case_list -> add: ...
	AddToCaseList(CaseList_ []*Branch, Case_ *Value, Body_ []Statement) ([]*Branch, error)

	// 62:4: case_list -> case: ...
	CaseToCaseList(Case_ *Value, Body_ []Statement) ([]*Branch, error)

	// 65:4: optional_default -> default: ...
	DefaultToOptionalDefault(Default_ *TToken, Body_ []Statement) (*Branch, error)

	// 66:4: optional_default -> nil: ...
	NilToOptionalDefault() (*Branch, error)

	// 68:6: if -> ...
	ToIf(If_ *Value, Body_ []Statement, ElseIfList_ []*Branch, OptionalElse_ *Branch, End_ *TToken) (Statement, error)

	// 71:4: else_if_list -> add: ...
	AddToElseIfList(ElseIfList_ []*Branch, ElseIf_ *Value, Body_ []Statement) ([]*Branch, error)

	// 72:4: else_if_list -> nil: ...
	NilToElseIfList() ([]*Branch, error)

	// 75:4: optional_else -> else: ...
	ElseToOptionalElse(Else_ *TToken, Body_ []Statement) (*Branch, error)

	// 76:4: optional_else -> nil: ...
	NilToOptionalElse() (*Branch, error)
}

type ParseErrorHandler interface {
	Error(nextToken Token, parseStack _Stack) error
}

type DefaultParseErrorHandler struct{}

func (DefaultParseErrorHandler) Error(nextToken Token, stack _Stack) error {
	return fmt.Errorf(
		"Syntax error: unexpected symbol %v. Expecting %v (%v)",
		nextToken.Id(),
		ExpectedTerminals(stack[len(stack)-1].StateId),
		nextToken.Loc())
}

func ExpectedTerminals(id _StateId) []SymbolId {
	result := []SymbolId{}
	for key, _ := range _ActionTable {
		if key._StateId != id {
			continue
		}
		result = append(result, key.SymbolId)
	}

	sort.Slice(result, func(i int, j int) bool { return result[i] < result[j] })
	return result
}

func Parse(lexer Lexer, reducer Reducer) (*File, error) {

	return ParseWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*File,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal *File
		return errRetVal, err
	}
	return item.File, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _Parse(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
	startState _StateId) (
	*_StackItem,
	error) {

	stateStack := _Stack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_StackItem{startState, nil},
	}

	symbolStack := &_PseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _ActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _ShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _ReduceAction {
			var reduceSymbol *Symbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _AcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i SymbolId) String() string {
	switch i {
	case _EndMarker:
		return "$"
	case _WildcardMarker:
		return "*"
	case SectionMarkerToken:
		return "SECTION_MARKER"
	case PackageToken:
		return "PACKAGE"
	case ImportToken:
		return "IMPORT"
	case TemplateDeclToken:
		return "TEMPLATE_DECL"
	case ForToken:
		return "FOR"
	case SwitchToken:
		return "SWITCH"
	case CaseToken:
		return "CASE"
	case IfToken:
		return "IF"
	case ElseIfToken:
		return "ELSE_IF"
	case DefaultToken:
		return "DEFAULT"
	case ElseToken:
		return "ELSE"
	case EndToken:
		return "END"
	case TextToken:
		return "TEXT"
	case SubstitutionToken:
		return "SUBSTITUTION"
	case EmbedToken:
		return "EMBED"
	case CopySectionToken:
		return "COPY_SECTION"
	case CommentToken:
		return "COMMENT"
	case ContinueToken:
		return "CONTINUE"
	case BreakToken:
		return "BREAK"
	case ReturnToken:
		return "RETURN"
	case ErrorToken:
		return "ERROR"
	case FileType:
		return "file"
	case OptionalImportsType:
		return "optional_imports"
	case BodyType:
		return "body"
	case StatementType:
		return "statement"
	case AtomType:
		return "atom"
	case ForType:
		return "for"
	case SwitchType:
		return "switch"
	case CaseListType:
		return "case_list"
	case OptionalDefaultType:
		return "optional_default"
	case IfType:
		return "if"
	case ElseIfListType:
		return "else_if_list"
	case OptionalElseType:
		return "optional_else"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	FileType            = SymbolId(277)
	OptionalImportsType = SymbolId(278)
	BodyType            = SymbolId(279)
	StatementType       = SymbolId(280)
	AtomType            = SymbolId(281)
	ForType             = SymbolId(282)
	SwitchType          = SymbolId(283)
	CaseListType        = SymbolId(284)
	OptionalDefaultType = SymbolId(285)
	IfType              = SymbolId(286)
	ElseIfListType      = SymbolId(287)
	OptionalElseType    = SymbolId(288)
)

type _ActionType int

const (
	// NOTE: error action is implicit
	_ShiftAction  = _ActionType(0)
	_ReduceAction = _ActionType(1)
	_AcceptAction = _ActionType(2)
)

func (i _ActionType) String() string {
	switch i {
	case _ShiftAction:
		return "shift"
	case _ReduceAction:
		return "reduce"
	case _AcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _ReduceType int

const (
	_ReduceToFile                    = _ReduceType(1)
	_ReduceImportsToOptionalImports  = _ReduceType(2)
	_ReduceNilToOptionalImports      = _ReduceType(3)
	_ReduceAddToBody                 = _ReduceType(4)
	_ReduceNilToBody                 = _ReduceType(5)
	_ReduceAtomToStatement           = _ReduceType(6)
	_ReduceForToStatement            = _ReduceType(7)
	_ReduceSwitchToStatement         = _ReduceType(8)
	_ReduceIfToStatement             = _ReduceType(9)
	_ReduceTextToAtom                = _ReduceType(10)
	_ReduceSubstitutionToAtom        = _ReduceType(11)
	_ReduceEmbedToAtom               = _ReduceType(12)
	_ReduceCopySectionToAtom         = _ReduceType(13)
	_ReduceCommentToAtom             = _ReduceType(14)
	_ReduceContinueToAtom            = _ReduceType(15)
	_ReduceBreakToAtom               = _ReduceType(16)
	_ReduceReturnToAtom              = _ReduceType(17)
	_ReduceErrorToAtom               = _ReduceType(18)
	_ReduceToFor                     = _ReduceType(19)
	_ReduceWithWhitespaceToSwitch    = _ReduceType(20)
	_ReduceWithoutWhitespaceToSwitch = _ReduceType(21)
	_ReduceAddToCaseList             = _ReduceType(22)
	_ReduceCaseToCaseList            = _ReduceType(23)
	_ReduceDefaultToOptionalDefault  = _ReduceType(24)
	_ReduceNilToOptionalDefault      = _ReduceType(25)
	_ReduceToIf                      = _ReduceType(26)
	_ReduceAddToElseIfList           = _ReduceType(27)
	_ReduceNilToElseIfList           = _ReduceType(28)
	_ReduceElseToOptionalElse        = _ReduceType(29)
	_ReduceNilToOptionalElse         = _ReduceType(30)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceToFile:
		return "ToFile"
	case _ReduceImportsToOptionalImports:
		return "ImportsToOptionalImports"
	case _ReduceNilToOptionalImports:
		return "NilToOptionalImports"
	case _ReduceAddToBody:
		return "AddToBody"
	case _ReduceNilToBody:
		return "NilToBody"
	case _ReduceAtomToStatement:
		return "AtomToStatement"
	case _ReduceForToStatement:
		return "ForToStatement"
	case _ReduceSwitchToStatement:
		return "SwitchToStatement"
	case _ReduceIfToStatement:
		return "IfToStatement"
	case _ReduceTextToAtom:
		return "TextToAtom"
	case _ReduceSubstitutionToAtom:
		return "SubstitutionToAtom"
	case _ReduceEmbedToAtom:
		return "EmbedToAtom"
	case _ReduceCopySectionToAtom:
		return "CopySectionToAtom"
	case _ReduceCommentToAtom:
		return "CommentToAtom"
	case _ReduceContinueToAtom:
		return "ContinueToAtom"
	case _ReduceBreakToAtom:
		return "BreakToAtom"
	case _ReduceReturnToAtom:
		return "ReturnToAtom"
	case _ReduceErrorToAtom:
		return "ErrorToAtom"
	case _ReduceToFor:
		return "ToFor"
	case _ReduceWithWhitespaceToSwitch:
		return "WithWhitespaceToSwitch"
	case _ReduceWithoutWhitespaceToSwitch:
		return "WithoutWhitespaceToSwitch"
	case _ReduceAddToCaseList:
		return "AddToCaseList"
	case _ReduceCaseToCaseList:
		return "CaseToCaseList"
	case _ReduceDefaultToOptionalDefault:
		return "DefaultToOptionalDefault"
	case _ReduceNilToOptionalDefault:
		return "NilToOptionalDefault"
	case _ReduceToIf:
		return "ToIf"
	case _ReduceAddToElseIfList:
		return "AddToElseIfList"
	case _ReduceNilToElseIfList:
		return "NilToElseIfList"
	case _ReduceElseToOptionalElse:
		return "ElseToOptionalElse"
	case _ReduceNilToOptionalElse:
		return "NilToOptionalElse"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _StateId int

func (id _StateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_State1  = _StateId(1)
	_State2  = _StateId(2)
	_State3  = _StateId(3)
	_State4  = _StateId(4)
	_State5  = _StateId(5)
	_State6  = _StateId(6)
	_State7  = _StateId(7)
	_State8  = _StateId(8)
	_State9  = _StateId(9)
	_State10 = _StateId(10)
	_State11 = _StateId(11)
	_State12 = _StateId(12)
	_State13 = _StateId(13)
	_State14 = _StateId(14)
	_State15 = _StateId(15)
	_State16 = _StateId(16)
	_State17 = _StateId(17)
	_State18 = _StateId(18)
	_State19 = _StateId(19)
	_State20 = _StateId(20)
	_State21 = _StateId(21)
	_State22 = _StateId(22)
	_State23 = _StateId(23)
	_State24 = _StateId(24)
	_State25 = _StateId(25)
	_State26 = _StateId(26)
	_State27 = _StateId(27)
	_State28 = _StateId(28)
	_State29 = _StateId(29)
	_State30 = _StateId(30)
	_State31 = _StateId(31)
	_State32 = _StateId(32)
	_State33 = _StateId(33)
	_State34 = _StateId(34)
	_State35 = _StateId(35)
	_State36 = _StateId(36)
	_State37 = _StateId(37)
	_State38 = _StateId(38)
	_State39 = _StateId(39)
	_State40 = _StateId(40)
	_State41 = _StateId(41)
	_State42 = _StateId(42)
	_State43 = _StateId(43)
	_State44 = _StateId(44)
	_State45 = _StateId(45)
	_State46 = _StateId(46)
	_State47 = _StateId(47)
	_State48 = _StateId(48)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ *GenericSymbol

	Atom         *Atom
	Branch       *Branch
	Branches     []*Branch
	File         *File
	Statement    Statement
	Statements   []Statement
	TemplateDecl *TemplateDeclaration
	Token        *TToken
	Value        *Value
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case TextToken, SubstitutionToken, EmbedToken, CopySectionToken, CommentToken, ContinueToken, BreakToken, ReturnToken, ErrorToken:
		val, ok := token.(*Atom)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *Atom (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Atom = val
	case _EndMarker, SectionMarkerToken:
		val, ok := token.(*GenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *GenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	case TemplateDeclToken:
		val, ok := token.(*TemplateDeclaration)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *TemplateDeclaration (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.TemplateDecl = val
	case DefaultToken, ElseToken, EndToken:
		val, ok := token.(*TToken)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *TToken (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Token = val
	case PackageToken, ImportToken, ForToken, SwitchToken, CaseToken, IfToken, ElseIfToken:
		val, ok := token.(*Value)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *Value (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Value = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *Symbol) Id() SymbolId {
	return s.SymbolId_
}

func (s *Symbol) Loc() Location {
	type locator interface{ Loc() Location }
	switch s.SymbolId_ {
	case TextToken, SubstitutionToken, EmbedToken, CopySectionToken, CommentToken, ContinueToken, BreakToken, ReturnToken, ErrorToken:
		loc, ok := interface{}(s.Atom).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalDefaultType, OptionalElseType:
		loc, ok := interface{}(s.Branch).(locator)
		if ok {
			return loc.Loc()
		}
	case CaseListType, ElseIfListType:
		loc, ok := interface{}(s.Branches).(locator)
		if ok {
			return loc.Loc()
		}
	case FileType:
		loc, ok := interface{}(s.File).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementType, AtomType, ForType, SwitchType, IfType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case BodyType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case TemplateDeclToken:
		loc, ok := interface{}(s.TemplateDecl).(locator)
		if ok {
			return loc.Loc()
		}
	case DefaultToken, ElseToken, EndToken:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Loc()
		}
	case PackageToken, ImportToken, ForToken, SwitchToken, CaseToken, IfToken, ElseIfToken, OptionalImportsType:
		loc, ok := interface{}(s.Value).(locator)
		if ok {
			return loc.Loc()
		}
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return Location{}
}

type _PseudoSymbolStack struct {
	lexer Lexer
	top   []*Symbol
}

func (stack *_PseudoSymbolStack) Top() (*Symbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &GenericSymbol{_EndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_PseudoSymbolStack) Push(symbol *Symbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_PseudoSymbolStack) Pop() (*Symbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _StackItem struct {
	StateId _StateId

	*Symbol
}

type _Stack []*_StackItem

type _Action struct {
	ActionType _ActionType

	ShiftStateId _StateId
	ReduceType   _ReduceType
}

func (act *_Action) ShiftItem(symbol *Symbol) *_StackItem {
	return &_StackItem{StateId: act.ShiftStateId, Symbol: symbol}
}

func (act *_Action) ReduceSymbol(
	reducer Reducer,
	stack _Stack) (
	_Stack,
	*Symbol,
	error) {

	var err error
	symbol := &Symbol{}
	switch act.ReduceType {
	case _ReduceToFile:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FileType
		symbol.File, err = reducer.ToFile(args[0].Value, args[1].Value, args[2].TemplateDecl, args[3].Generic_, args[4].Statements)
	case _ReduceImportsToOptionalImports:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalImportsType
		symbol.Value, err = reducer.ImportsToOptionalImports(args[0].Value)
	case _ReduceNilToOptionalImports:
		symbol.SymbolId_ = OptionalImportsType
		symbol.Value, err = reducer.NilToOptionalImports()
	case _ReduceAddToBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BodyType
		symbol.Statements, err = reducer.AddToBody(args[0].Statements, args[1].Statement)
	case _ReduceNilToBody:
		symbol.SymbolId_ = BodyType
		symbol.Statements, err = reducer.NilToBody()
	case _ReduceAtomToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.AtomToStatement(args[0].Statement)
	case _ReduceForToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.ForToStatement(args[0].Statement)
	case _ReduceSwitchToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.SwitchToStatement(args[0].Statement)
	case _ReduceIfToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.IfToStatement(args[0].Statement)
	case _ReduceTextToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.TextToAtom(args[0].Atom)
	case _ReduceSubstitutionToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.SubstitutionToAtom(args[0].Atom)
	case _ReduceEmbedToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.EmbedToAtom(args[0].Atom)
	case _ReduceCopySectionToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.CopySectionToAtom(args[0].Atom)
	case _ReduceCommentToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.CommentToAtom(args[0].Atom)
	case _ReduceContinueToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.ContinueToAtom(args[0].Atom)
	case _ReduceBreakToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.BreakToAtom(args[0].Atom)
	case _ReduceReturnToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.ReturnToAtom(args[0].Atom)
	case _ReduceErrorToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Statement, err = reducer.ErrorToAtom(args[0].Atom)
	case _ReduceToFor:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ForType
		symbol.Statement, err = reducer.ToFor(args[0].Value, args[1].Statements, args[2].Token)
	case _ReduceWithWhitespaceToSwitch:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = SwitchType
		symbol.Statement, err = reducer.WithWhitespaceToSwitch(args[0].Value, args[1].Atom, args[2].Branches, args[3].Branch, args[4].Token)
	case _ReduceWithoutWhitespaceToSwitch:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = SwitchType
		symbol.Statement, err = reducer.WithoutWhitespaceToSwitch(args[0].Value, args[1].Branches, args[2].Branch, args[3].Token)
	case _ReduceAddToCaseList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CaseListType
		symbol.Branches, err = reducer.AddToCaseList(args[0].Branches, args[1].Value, args[2].Statements)
	case _ReduceCaseToCaseList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CaseListType
		symbol.Branches, err = reducer.CaseToCaseList(args[0].Value, args[1].Statements)
	case _ReduceDefaultToOptionalDefault:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Branch, err = reducer.DefaultToOptionalDefault(args[0].Token, args[1].Statements)
	case _ReduceNilToOptionalDefault:
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Branch, err = reducer.NilToOptionalDefault()
	case _ReduceToIf:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfType
		symbol.Statement, err = reducer.ToIf(args[0].Value, args[1].Statements, args[2].Branches, args[3].Branch, args[4].Token)
	case _ReduceAddToElseIfList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ElseIfListType
		symbol.Branches, err = reducer.AddToElseIfList(args[0].Branches, args[1].Value, args[2].Statements)
	case _ReduceNilToElseIfList:
		symbol.SymbolId_ = ElseIfListType
		symbol.Branches, err = reducer.NilToElseIfList()
	case _ReduceElseToOptionalElse:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalElseType
		symbol.Branch, err = reducer.ElseToOptionalElse(args[0].Token, args[1].Statements)
	case _ReduceNilToOptionalElse:
		symbol.SymbolId_ = OptionalElseType
		symbol.Branch, err = reducer.NilToOptionalElse()
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _ActionTableKey struct {
	_StateId
	SymbolId
}

type _ActionTableType map[_ActionTableKey]*_Action

func (table _ActionTableType) Get(
	stateId _StateId,
	symbolId SymbolId) (
	*_Action,
	bool) {

	action, ok := table[_ActionTableKey{stateId, symbolId}]
	if ok {
		return action, ok
	}

	action, ok = table[_ActionTableKey{stateId, _WildcardMarker}]
	return action, ok
}

var (
	_GotoState1Action                      = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                      = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                      = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                      = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                      = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                      = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                      = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                      = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                      = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                     = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                     = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                     = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                     = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                     = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                     = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                     = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                     = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                     = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                     = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                     = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                     = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                     = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                     = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                     = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                     = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                     = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                     = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                     = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                     = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                     = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                     = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                     = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                     = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                     = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                     = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                     = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                     = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                     = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                     = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                     = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                     = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                     = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                     = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                     = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                     = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                     = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                     = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                     = &_Action{_ShiftAction, _State48, 0}
	_ReduceToFileAction                    = &_Action{_ReduceAction, 0, _ReduceToFile}
	_ReduceImportsToOptionalImportsAction  = &_Action{_ReduceAction, 0, _ReduceImportsToOptionalImports}
	_ReduceNilToOptionalImportsAction      = &_Action{_ReduceAction, 0, _ReduceNilToOptionalImports}
	_ReduceAddToBodyAction                 = &_Action{_ReduceAction, 0, _ReduceAddToBody}
	_ReduceNilToBodyAction                 = &_Action{_ReduceAction, 0, _ReduceNilToBody}
	_ReduceAtomToStatementAction           = &_Action{_ReduceAction, 0, _ReduceAtomToStatement}
	_ReduceForToStatementAction            = &_Action{_ReduceAction, 0, _ReduceForToStatement}
	_ReduceSwitchToStatementAction         = &_Action{_ReduceAction, 0, _ReduceSwitchToStatement}
	_ReduceIfToStatementAction             = &_Action{_ReduceAction, 0, _ReduceIfToStatement}
	_ReduceTextToAtomAction                = &_Action{_ReduceAction, 0, _ReduceTextToAtom}
	_ReduceSubstitutionToAtomAction        = &_Action{_ReduceAction, 0, _ReduceSubstitutionToAtom}
	_ReduceEmbedToAtomAction               = &_Action{_ReduceAction, 0, _ReduceEmbedToAtom}
	_ReduceCopySectionToAtomAction         = &_Action{_ReduceAction, 0, _ReduceCopySectionToAtom}
	_ReduceCommentToAtomAction             = &_Action{_ReduceAction, 0, _ReduceCommentToAtom}
	_ReduceContinueToAtomAction            = &_Action{_ReduceAction, 0, _ReduceContinueToAtom}
	_ReduceBreakToAtomAction               = &_Action{_ReduceAction, 0, _ReduceBreakToAtom}
	_ReduceReturnToAtomAction              = &_Action{_ReduceAction, 0, _ReduceReturnToAtom}
	_ReduceErrorToAtomAction               = &_Action{_ReduceAction, 0, _ReduceErrorToAtom}
	_ReduceToForAction                     = &_Action{_ReduceAction, 0, _ReduceToFor}
	_ReduceWithWhitespaceToSwitchAction    = &_Action{_ReduceAction, 0, _ReduceWithWhitespaceToSwitch}
	_ReduceWithoutWhitespaceToSwitchAction = &_Action{_ReduceAction, 0, _ReduceWithoutWhitespaceToSwitch}
	_ReduceAddToCaseListAction             = &_Action{_ReduceAction, 0, _ReduceAddToCaseList}
	_ReduceCaseToCaseListAction            = &_Action{_ReduceAction, 0, _ReduceCaseToCaseList}
	_ReduceDefaultToOptionalDefaultAction  = &_Action{_ReduceAction, 0, _ReduceDefaultToOptionalDefault}
	_ReduceNilToOptionalDefaultAction      = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}
	_ReduceToIfAction                      = &_Action{_ReduceAction, 0, _ReduceToIf}
	_ReduceAddToElseIfListAction           = &_Action{_ReduceAction, 0, _ReduceAddToElseIfList}
	_ReduceNilToElseIfListAction           = &_Action{_ReduceAction, 0, _ReduceNilToElseIfList}
	_ReduceElseToOptionalElseAction        = &_Action{_ReduceAction, 0, _ReduceElseToOptionalElse}
	_ReduceNilToOptionalElseAction         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalElse}
)

var _ActionTable = _ActionTableType{
	{_State2, _EndMarker}:           &_Action{_AcceptAction, 0, 0},
	{_State1, PackageToken}:         _GotoState3Action,
	{_State1, FileType}:             _GotoState2Action,
	{_State3, ImportToken}:          _GotoState4Action,
	{_State3, OptionalImportsType}:  _GotoState5Action,
	{_State5, TemplateDeclToken}:    _GotoState6Action,
	{_State6, SectionMarkerToken}:   _GotoState7Action,
	{_State7, BodyType}:             _GotoState8Action,
	{_State8, ForToken}:             _GotoState15Action,
	{_State8, SwitchToken}:          _GotoState19Action,
	{_State8, IfToken}:              _GotoState16Action,
	{_State8, TextToken}:            _GotoState20Action,
	{_State8, SubstitutionToken}:    _GotoState18Action,
	{_State8, EmbedToken}:           _GotoState13Action,
	{_State8, CopySectionToken}:     _GotoState12Action,
	{_State8, CommentToken}:         _GotoState10Action,
	{_State8, ContinueToken}:        _GotoState11Action,
	{_State8, BreakToken}:           _GotoState9Action,
	{_State8, ReturnToken}:          _GotoState17Action,
	{_State8, ErrorToken}:           _GotoState14Action,
	{_State8, StatementType}:        _GotoState24Action,
	{_State8, AtomType}:             _GotoState21Action,
	{_State8, ForType}:              _GotoState22Action,
	{_State8, SwitchType}:           _GotoState25Action,
	{_State8, IfType}:               _GotoState23Action,
	{_State15, BodyType}:            _GotoState26Action,
	{_State16, BodyType}:            _GotoState27Action,
	{_State19, CaseToken}:           _GotoState28Action,
	{_State19, TextToken}:           _GotoState29Action,
	{_State19, CaseListType}:        _GotoState30Action,
	{_State26, ForToken}:            _GotoState15Action,
	{_State26, SwitchToken}:         _GotoState19Action,
	{_State26, IfToken}:             _GotoState16Action,
	{_State26, EndToken}:            _GotoState31Action,
	{_State26, TextToken}:           _GotoState20Action,
	{_State26, SubstitutionToken}:   _GotoState18Action,
	{_State26, EmbedToken}:          _GotoState13Action,
	{_State26, CopySectionToken}:    _GotoState12Action,
	{_State26, CommentToken}:        _GotoState10Action,
	{_State26, ContinueToken}:       _GotoState11Action,
	{_State26, BreakToken}:          _GotoState9Action,
	{_State26, ReturnToken}:         _GotoState17Action,
	{_State26, ErrorToken}:          _GotoState14Action,
	{_State26, StatementType}:       _GotoState24Action,
	{_State26, AtomType}:            _GotoState21Action,
	{_State26, ForType}:             _GotoState22Action,
	{_State26, SwitchType}:          _GotoState25Action,
	{_State26, IfType}:              _GotoState23Action,
	{_State27, ForToken}:            _GotoState15Action,
	{_State27, SwitchToken}:         _GotoState19Action,
	{_State27, IfToken}:             _GotoState16Action,
	{_State27, TextToken}:           _GotoState20Action,
	{_State27, SubstitutionToken}:   _GotoState18Action,
	{_State27, EmbedToken}:          _GotoState13Action,
	{_State27, CopySectionToken}:    _GotoState12Action,
	{_State27, CommentToken}:        _GotoState10Action,
	{_State27, ContinueToken}:       _GotoState11Action,
	{_State27, BreakToken}:          _GotoState9Action,
	{_State27, ReturnToken}:         _GotoState17Action,
	{_State27, ErrorToken}:          _GotoState14Action,
	{_State27, StatementType}:       _GotoState24Action,
	{_State27, AtomType}:            _GotoState21Action,
	{_State27, ForType}:             _GotoState22Action,
	{_State27, SwitchType}:          _GotoState25Action,
	{_State27, IfType}:              _GotoState23Action,
	{_State27, ElseIfListType}:      _GotoState32Action,
	{_State28, BodyType}:            _GotoState33Action,
	{_State29, CaseToken}:           _GotoState28Action,
	{_State29, CaseListType}:        _GotoState34Action,
	{_State30, CaseToken}:           _GotoState35Action,
	{_State30, DefaultToken}:        _GotoState36Action,
	{_State30, OptionalDefaultType}: _GotoState37Action,
	{_State32, ElseIfToken}:         _GotoState39Action,
	{_State32, ElseToken}:           _GotoState38Action,
	{_State32, OptionalElseType}:    _GotoState40Action,
	{_State33, ForToken}:            _GotoState15Action,
	{_State33, SwitchToken}:         _GotoState19Action,
	{_State33, IfToken}:             _GotoState16Action,
	{_State33, TextToken}:           _GotoState20Action,
	{_State33, SubstitutionToken}:   _GotoState18Action,
	{_State33, EmbedToken}:          _GotoState13Action,
	{_State33, CopySectionToken}:    _GotoState12Action,
	{_State33, CommentToken}:        _GotoState10Action,
	{_State33, ContinueToken}:       _GotoState11Action,
	{_State33, BreakToken}:          _GotoState9Action,
	{_State33, ReturnToken}:         _GotoState17Action,
	{_State33, ErrorToken}:          _GotoState14Action,
	{_State33, StatementType}:       _GotoState24Action,
	{_State33, AtomType}:            _GotoState21Action,
	{_State33, ForType}:             _GotoState22Action,
	{_State33, SwitchType}:          _GotoState25Action,
	{_State33, IfType}:              _GotoState23Action,
	{_State34, CaseToken}:           _GotoState35Action,
	{_State34, DefaultToken}:        _GotoState36Action,
	{_State34, OptionalDefaultType}: _GotoState41Action,
	{_State35, BodyType}:            _GotoState42Action,
	{_State36, BodyType}:            _GotoState43Action,
	{_State37, EndToken}:            _GotoState44Action,
	{_State38, BodyType}:            _GotoState45Action,
	{_State39, BodyType}:            _GotoState46Action,
	{_State40, EndToken}:            _GotoState47Action,
	{_State41, EndToken}:            _GotoState48Action,
	{_State42, ForToken}:            _GotoState15Action,
	{_State42, SwitchToken}:         _GotoState19Action,
	{_State42, IfToken}:             _GotoState16Action,
	{_State42, TextToken}:           _GotoState20Action,
	{_State42, SubstitutionToken}:   _GotoState18Action,
	{_State42, EmbedToken}:          _GotoState13Action,
	{_State42, CopySectionToken}:    _GotoState12Action,
	{_State42, CommentToken}:        _GotoState10Action,
	{_State42, ContinueToken}:       _GotoState11Action,
	{_State42, BreakToken}:          _GotoState9Action,
	{_State42, ReturnToken}:         _GotoState17Action,
	{_State42, ErrorToken}:          _GotoState14Action,
	{_State42, StatementType}:       _GotoState24Action,
	{_State42, AtomType}:            _GotoState21Action,
	{_State42, ForType}:             _GotoState22Action,
	{_State42, SwitchType}:          _GotoState25Action,
	{_State42, IfType}:              _GotoState23Action,
	{_State43, ForToken}:            _GotoState15Action,
	{_State43, SwitchToken}:         _GotoState19Action,
	{_State43, IfToken}:             _GotoState16Action,
	{_State43, TextToken}:           _GotoState20Action,
	{_State43, SubstitutionToken}:   _GotoState18Action,
	{_State43, EmbedToken}:          _GotoState13Action,
	{_State43, CopySectionToken}:    _GotoState12Action,
	{_State43, CommentToken}:        _GotoState10Action,
	{_State43, ContinueToken}:       _GotoState11Action,
	{_State43, BreakToken}:          _GotoState9Action,
	{_State43, ReturnToken}:         _GotoState17Action,
	{_State43, ErrorToken}:          _GotoState14Action,
	{_State43, StatementType}:       _GotoState24Action,
	{_State43, AtomType}:            _GotoState21Action,
	{_State43, ForType}:             _GotoState22Action,
	{_State43, SwitchType}:          _GotoState25Action,
	{_State43, IfType}:              _GotoState23Action,
	{_State45, ForToken}:            _GotoState15Action,
	{_State45, SwitchToken}:         _GotoState19Action,
	{_State45, IfToken}:             _GotoState16Action,
	{_State45, TextToken}:           _GotoState20Action,
	{_State45, SubstitutionToken}:   _GotoState18Action,
	{_State45, EmbedToken}:          _GotoState13Action,
	{_State45, CopySectionToken}:    _GotoState12Action,
	{_State45, CommentToken}:        _GotoState10Action,
	{_State45, ContinueToken}:       _GotoState11Action,
	{_State45, BreakToken}:          _GotoState9Action,
	{_State45, ReturnToken}:         _GotoState17Action,
	{_State45, ErrorToken}:          _GotoState14Action,
	{_State45, StatementType}:       _GotoState24Action,
	{_State45, AtomType}:            _GotoState21Action,
	{_State45, ForType}:             _GotoState22Action,
	{_State45, SwitchType}:          _GotoState25Action,
	{_State45, IfType}:              _GotoState23Action,
	{_State46, ForToken}:            _GotoState15Action,
	{_State46, SwitchToken}:         _GotoState19Action,
	{_State46, IfToken}:             _GotoState16Action,
	{_State46, TextToken}:           _GotoState20Action,
	{_State46, SubstitutionToken}:   _GotoState18Action,
	{_State46, EmbedToken}:          _GotoState13Action,
	{_State46, CopySectionToken}:    _GotoState12Action,
	{_State46, CommentToken}:        _GotoState10Action,
	{_State46, ContinueToken}:       _GotoState11Action,
	{_State46, BreakToken}:          _GotoState9Action,
	{_State46, ReturnToken}:         _GotoState17Action,
	{_State46, ErrorToken}:          _GotoState14Action,
	{_State46, StatementType}:       _GotoState24Action,
	{_State46, AtomType}:            _GotoState21Action,
	{_State46, ForType}:             _GotoState22Action,
	{_State46, SwitchType}:          _GotoState25Action,
	{_State46, IfType}:              _GotoState23Action,
	{_State3, TemplateDeclToken}:    _ReduceNilToOptionalImportsAction,
	{_State4, TemplateDeclToken}:    _ReduceImportsToOptionalImportsAction,
	{_State7, _WildcardMarker}:      _ReduceNilToBodyAction,
	{_State8, _EndMarker}:           _ReduceToFileAction,
	{_State9, _WildcardMarker}:      _ReduceBreakToAtomAction,
	{_State10, _WildcardMarker}:     _ReduceCommentToAtomAction,
	{_State11, _WildcardMarker}:     _ReduceContinueToAtomAction,
	{_State12, _WildcardMarker}:     _ReduceCopySectionToAtomAction,
	{_State13, _WildcardMarker}:     _ReduceEmbedToAtomAction,
	{_State14, _WildcardMarker}:     _ReduceErrorToAtomAction,
	{_State15, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State16, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State17, _WildcardMarker}:     _ReduceReturnToAtomAction,
	{_State18, _WildcardMarker}:     _ReduceSubstitutionToAtomAction,
	{_State20, _WildcardMarker}:     _ReduceTextToAtomAction,
	{_State21, _WildcardMarker}:     _ReduceAtomToStatementAction,
	{_State22, _WildcardMarker}:     _ReduceForToStatementAction,
	{_State23, _WildcardMarker}:     _ReduceIfToStatementAction,
	{_State24, _WildcardMarker}:     _ReduceAddToBodyAction,
	{_State25, _WildcardMarker}:     _ReduceSwitchToStatementAction,
	{_State27, _WildcardMarker}:     _ReduceNilToElseIfListAction,
	{_State28, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State30, EndToken}:            _ReduceNilToOptionalDefaultAction,
	{_State31, _WildcardMarker}:     _ReduceToForAction,
	{_State32, EndToken}:            _ReduceNilToOptionalElseAction,
	{_State33, _WildcardMarker}:     _ReduceCaseToCaseListAction,
	{_State34, EndToken}:            _ReduceNilToOptionalDefaultAction,
	{_State35, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State36, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State38, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State39, _WildcardMarker}:     _ReduceNilToBodyAction,
	{_State42, _WildcardMarker}:     _ReduceAddToCaseListAction,
	{_State43, EndToken}:            _ReduceDefaultToOptionalDefaultAction,
	{_State44, _WildcardMarker}:     _ReduceWithoutWhitespaceToSwitchAction,
	{_State45, EndToken}:            _ReduceElseToOptionalElseAction,
	{_State46, _WildcardMarker}:     _ReduceAddToElseIfListAction,
	{_State47, _WildcardMarker}:     _ReduceToIfAction,
	{_State48, _WildcardMarker}:     _ReduceWithWhitespaceToSwitchAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.file
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 3
      file -> State 2

  State 2:
    Kernel Items:
      #accept: ^ file., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 3:
    Kernel Items:
      file: PACKAGE.optional_imports TEMPLATE_DECL SECTION_MARKER body
    Reduce:
      TEMPLATE_DECL -> [optional_imports]
    Goto:
      IMPORT -> State 4
      optional_imports -> State 5

  State 4:
    Kernel Items:
      optional_imports: IMPORT., TEMPLATE_DECL
    Reduce:
      TEMPLATE_DECL -> [optional_imports]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      file: PACKAGE optional_imports.TEMPLATE_DECL SECTION_MARKER body
    Reduce:
      (nil)
    Goto:
      TEMPLATE_DECL -> State 6

  State 6:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL.SECTION_MARKER body
    Reduce:
      (nil)
    Goto:
      SECTION_MARKER -> State 7

  State 7:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL SECTION_MARKER.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 8

  State 8:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL SECTION_MARKER body., $
      body: body.statement
    Reduce:
      $ -> [file]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 9:
    Kernel Items:
      atom: BREAK., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      atom: COMMENT., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      atom: CONTINUE., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      atom: COPY_SECTION., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      atom: EMBED., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      atom: ERROR., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      for: FOR.body END
    Reduce:
      * -> [body]
    Goto:
      body -> State 26

  State 16:
    Kernel Items:
      if: IF.body else_if_list optional_else END
    Reduce:
      * -> [body]
    Goto:
      body -> State 27

  State 17:
    Kernel Items:
      atom: RETURN., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      atom: SUBSTITUTION., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      switch: SWITCH.TEXT case_list optional_default END
      switch: SWITCH.case_list optional_default END
    Reduce:
      (nil)
    Goto:
      CASE -> State 28
      TEXT -> State 29
      case_list -> State 30

  State 20:
    Kernel Items:
      atom: TEXT., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      statement: atom., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      statement: for., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      statement: if., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      body: body statement., *
    Reduce:
      * -> [body]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      statement: switch., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      body: body.statement
      for: FOR body.END
    Reduce:
      (nil)
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      END -> State 31
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 27:
    Kernel Items:
      body: body.statement
      if: IF body.else_if_list optional_else END
    Reduce:
      * -> [else_if_list]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23
      else_if_list -> State 32

  State 28:
    Kernel Items:
      case_list: CASE.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 33

  State 29:
    Kernel Items:
      switch: SWITCH TEXT.case_list optional_default END
    Reduce:
      (nil)
    Goto:
      CASE -> State 28
      case_list -> State 34

  State 30:
    Kernel Items:
      switch: SWITCH case_list.optional_default END
      case_list: case_list.CASE body
    Reduce:
      END -> [optional_default]
    Goto:
      CASE -> State 35
      DEFAULT -> State 36
      optional_default -> State 37

  State 31:
    Kernel Items:
      for: FOR body END., *
    Reduce:
      * -> [for]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      if: IF body else_if_list.optional_else END
      else_if_list: else_if_list.ELSE_IF body
    Reduce:
      END -> [optional_else]
    Goto:
      ELSE_IF -> State 39
      ELSE -> State 38
      optional_else -> State 40

  State 33:
    Kernel Items:
      body: body.statement
      case_list: CASE body., *
    Reduce:
      * -> [case_list]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 34:
    Kernel Items:
      switch: SWITCH TEXT case_list.optional_default END
      case_list: case_list.CASE body
    Reduce:
      END -> [optional_default]
    Goto:
      CASE -> State 35
      DEFAULT -> State 36
      optional_default -> State 41

  State 35:
    Kernel Items:
      case_list: case_list CASE.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 42

  State 36:
    Kernel Items:
      optional_default: DEFAULT.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 43

  State 37:
    Kernel Items:
      switch: SWITCH case_list optional_default.END
    Reduce:
      (nil)
    Goto:
      END -> State 44

  State 38:
    Kernel Items:
      optional_else: ELSE.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 45

  State 39:
    Kernel Items:
      else_if_list: else_if_list ELSE_IF.body
    Reduce:
      * -> [body]
    Goto:
      body -> State 46

  State 40:
    Kernel Items:
      if: IF body else_if_list optional_else.END
    Reduce:
      (nil)
    Goto:
      END -> State 47

  State 41:
    Kernel Items:
      switch: SWITCH TEXT case_list optional_default.END
    Reduce:
      (nil)
    Goto:
      END -> State 48

  State 42:
    Kernel Items:
      body: body.statement
      case_list: case_list CASE body., *
    Reduce:
      * -> [case_list]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 43:
    Kernel Items:
      body: body.statement
      optional_default: DEFAULT body., END
    Reduce:
      END -> [optional_default]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 44:
    Kernel Items:
      switch: SWITCH case_list optional_default END., *
    Reduce:
      * -> [switch]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      body: body.statement
      optional_else: ELSE body., END
    Reduce:
      END -> [optional_else]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 46:
    Kernel Items:
      body: body.statement
      else_if_list: else_if_list ELSE_IF body., *
    Reduce:
      * -> [else_if_list]
    Goto:
      FOR -> State 15
      SWITCH -> State 19
      IF -> State 16
      TEXT -> State 20
      SUBSTITUTION -> State 18
      EMBED -> State 13
      COPY_SECTION -> State 12
      COMMENT -> State 10
      CONTINUE -> State 11
      BREAK -> State 9
      RETURN -> State 17
      ERROR -> State 14
      statement -> State 24
      atom -> State 21
      for -> State 22
      switch -> State 25
      if -> State 23

  State 47:
    Kernel Items:
      if: IF body else_if_list optional_else END., *
    Reduce:
      * -> [if]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      switch: SWITCH TEXT case_list optional_default END., *
    Reduce:
      * -> [switch]
    Goto:
      (nil)

Number of states: 48
Number of shift actions: 169
Number of reduce actions: 39
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
