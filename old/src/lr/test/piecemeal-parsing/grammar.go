// Auto-generated from source: grammar.lr

package main

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	IdToken    = SymbolId(256)
	ErrorToken = SymbolId(257)
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
	// 18:4: expr_list -> add: ...
	AddToExprList(ExprList_ []Expr, Expr_ Expr) ([]Expr, error)

	// 19:4: expr_list -> nil: ...
	NilToExprList() ([]Expr, error)

	// 22:4: atom -> id: ...
	IdToAtom(Id_ *Id) (Expr, error)

	// 23:4: atom -> error: ...
	ErrorToAtom(Error_ *Err) (Expr, error)

	// 24:4: atom -> block: ...
	BlockToAtom(Block_ *Block) (Expr, error)

	// 27:4: expr -> atom: ...
	AtomToExpr(Atom_ Expr) (Expr, error)

	// 28:4: expr -> binary: ...
	BinaryToExpr(Expr_ Expr, Op_ *GenericSymbol, Atom_ Expr) (Expr, error)

	// 31:4: op -> plus: ...
	PlusToOp(char *GenericSymbol) (*GenericSymbol, error)

	// 32:4: op -> minus: ...
	MinusToOp(char *GenericSymbol) (*GenericSymbol, error)

	// 34:9: block -> ...
	ToBlock(char *GenericSymbol, ExprList_ []Expr, char2 *GenericSymbol) (*Block, error)
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

func ParseExprList(lexer Lexer, reducer Reducer) ([]Expr, error) {

	return ParseExprListWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExprListWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	[]Expr,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal []Expr
		return errRetVal, err
	}
	return item.ExprList, nil
}

func ParseBlock(lexer Lexer, reducer Reducer) (*Block, error) {

	return ParseBlockWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseBlockWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*Block,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal *Block
		return errRetVal, err
	}
	return item.Block, nil
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
	case '+':
		return "'+'"
	case '-':
		return "'-'"
	case '{':
		return "'{'"
	case '}':
		return "'}'"
	case IdToken:
		return "ID"
	case ErrorToken:
		return "ERROR"
	case ExprListType:
		return "expr_list"
	case AtomType:
		return "atom"
	case ExprType:
		return "expr"
	case OpType:
		return "op"
	case BlockType:
		return "block"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	ExprListType = SymbolId(262)
	AtomType     = SymbolId(263)
	ExprType     = SymbolId(264)
	OpType       = SymbolId(265)
	BlockType    = SymbolId(266)
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
	_ReduceAddToExprList = _ReduceType(1)
	_ReduceNilToExprList = _ReduceType(2)
	_ReduceIdToAtom      = _ReduceType(3)
	_ReduceErrorToAtom   = _ReduceType(4)
	_ReduceBlockToAtom   = _ReduceType(5)
	_ReduceAtomToExpr    = _ReduceType(6)
	_ReduceBinaryToExpr  = _ReduceType(7)
	_ReducePlusToOp      = _ReduceType(8)
	_ReduceMinusToOp     = _ReduceType(9)
	_ReduceToBlock       = _ReduceType(10)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceAddToExprList:
		return "AddToExprList"
	case _ReduceNilToExprList:
		return "NilToExprList"
	case _ReduceIdToAtom:
		return "IdToAtom"
	case _ReduceErrorToAtom:
		return "ErrorToAtom"
	case _ReduceBlockToAtom:
		return "BlockToAtom"
	case _ReduceAtomToExpr:
		return "AtomToExpr"
	case _ReduceBinaryToExpr:
		return "BinaryToExpr"
	case _ReducePlusToOp:
		return "PlusToOp"
	case _ReduceMinusToOp:
		return "MinusToOp"
	case _ReduceToBlock:
		return "ToBlock"
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
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ *GenericSymbol

	Block    *Block
	Err      *Err
	Expr     Expr
	ExprList []Expr
	Ident    *Id
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case ErrorToken:
		val, ok := token.(*Err)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *Err (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Err = val
	case _EndMarker, '+', '-', '{', '}':
		val, ok := token.(*GenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *GenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	case IdToken:
		val, ok := token.(*Id)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *Id (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Ident = val
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
	case BlockType:
		loc, ok := interface{}(s.Block).(locator)
		if ok {
			return loc.Loc()
		}
	case ErrorToken:
		loc, ok := interface{}(s.Err).(locator)
		if ok {
			return loc.Loc()
		}
	case AtomType, ExprType:
		loc, ok := interface{}(s.Expr).(locator)
		if ok {
			return loc.Loc()
		}
	case ExprListType:
		loc, ok := interface{}(s.ExprList).(locator)
		if ok {
			return loc.Loc()
		}
	case IdToken:
		loc, ok := interface{}(s.Ident).(locator)
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
	case _ReduceAddToExprList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExprListType
		symbol.ExprList, err = reducer.AddToExprList(args[0].ExprList, args[1].Expr)
	case _ReduceNilToExprList:
		symbol.SymbolId_ = ExprListType
		symbol.ExprList, err = reducer.NilToExprList()
	case _ReduceIdToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Expr, err = reducer.IdToAtom(args[0].Ident)
	case _ReduceErrorToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Expr, err = reducer.ErrorToAtom(args[0].Err)
	case _ReduceBlockToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		symbol.Expr, err = reducer.BlockToAtom(args[0].Block)
	case _ReduceAtomToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		symbol.Expr, err = reducer.AtomToExpr(args[0].Expr)
	case _ReduceBinaryToExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExprType
		symbol.Expr, err = reducer.BinaryToExpr(args[0].Expr, args[1].Generic_, args[2].Expr)
	case _ReducePlusToOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpType
		symbol.Generic_, err = reducer.PlusToOp(args[0].Generic_)
	case _ReduceMinusToOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpType
		symbol.Generic_, err = reducer.MinusToOp(args[0].Generic_)
	case _ReduceToBlock:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BlockType
		symbol.Block, err = reducer.ToBlock(args[0].Generic_, args[1].ExprList, args[2].Generic_)
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
	_GotoState1Action          = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action          = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action          = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action          = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action          = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action          = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action          = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action          = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action          = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action         = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action         = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action         = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action         = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action         = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action         = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action         = &_Action{_ShiftAction, _State16, 0}
	_ReduceAddToExprListAction = &_Action{_ReduceAction, 0, _ReduceAddToExprList}
	_ReduceNilToExprListAction = &_Action{_ReduceAction, 0, _ReduceNilToExprList}
	_ReduceIdToAtomAction      = &_Action{_ReduceAction, 0, _ReduceIdToAtom}
	_ReduceErrorToAtomAction   = &_Action{_ReduceAction, 0, _ReduceErrorToAtom}
	_ReduceBlockToAtomAction   = &_Action{_ReduceAction, 0, _ReduceBlockToAtom}
	_ReduceAtomToExprAction    = &_Action{_ReduceAction, 0, _ReduceAtomToExpr}
	_ReduceBinaryToExprAction  = &_Action{_ReduceAction, 0, _ReduceBinaryToExpr}
	_ReducePlusToOpAction      = &_Action{_ReduceAction, 0, _ReducePlusToOp}
	_ReduceMinusToOpAction     = &_Action{_ReduceAction, 0, _ReduceMinusToOp}
	_ReduceToBlockAction       = &_Action{_ReduceAction, 0, _ReduceToBlock}
)

var _ActionTable = _ActionTableType{
	{_State3, _EndMarker}:       &_Action{_AcceptAction, 0, 0},
	{_State4, _EndMarker}:       &_Action{_AcceptAction, 0, 0},
	{_State1, ExprListType}:     _GotoState3Action,
	{_State2, '{'}:              _GotoState5Action,
	{_State2, BlockType}:        _GotoState4Action,
	{_State3, '{'}:              _GotoState5Action,
	{_State3, IdToken}:          _GotoState7Action,
	{_State3, ErrorToken}:       _GotoState6Action,
	{_State3, AtomType}:         _GotoState8Action,
	{_State3, ExprType}:         _GotoState10Action,
	{_State3, BlockType}:        _GotoState9Action,
	{_State5, ExprListType}:     _GotoState11Action,
	{_State10, '+'}:             _GotoState12Action,
	{_State10, '-'}:             _GotoState13Action,
	{_State10, OpType}:          _GotoState14Action,
	{_State11, '{'}:             _GotoState5Action,
	{_State11, '}'}:             _GotoState15Action,
	{_State11, IdToken}:         _GotoState7Action,
	{_State11, ErrorToken}:      _GotoState6Action,
	{_State11, AtomType}:        _GotoState8Action,
	{_State11, ExprType}:        _GotoState10Action,
	{_State11, BlockType}:       _GotoState9Action,
	{_State14, '{'}:             _GotoState5Action,
	{_State14, IdToken}:         _GotoState7Action,
	{_State14, ErrorToken}:      _GotoState6Action,
	{_State14, AtomType}:        _GotoState16Action,
	{_State14, BlockType}:       _GotoState9Action,
	{_State1, _WildcardMarker}:  _ReduceNilToExprListAction,
	{_State5, _WildcardMarker}:  _ReduceNilToExprListAction,
	{_State6, _WildcardMarker}:  _ReduceErrorToAtomAction,
	{_State7, _WildcardMarker}:  _ReduceIdToAtomAction,
	{_State8, _WildcardMarker}:  _ReduceAtomToExprAction,
	{_State9, _WildcardMarker}:  _ReduceBlockToAtomAction,
	{_State10, _WildcardMarker}: _ReduceAddToExprListAction,
	{_State12, _WildcardMarker}: _ReducePlusToOpAction,
	{_State13, _WildcardMarker}: _ReduceMinusToOpAction,
	{_State15, _EndMarker}:      _ReduceToBlockAction,
	{_State16, _WildcardMarker}: _ReduceBinaryToExprAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.expr_list
    Reduce:
      * -> [expr_list]
    Goto:
      expr_list -> State 3

  State 2:
    Kernel Items:
      #accept: ^.block
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      block -> State 4

  State 3:
    Kernel Items:
      #accept: ^ expr_list., $
      expr_list: expr_list.expr
    Reduce:
      $ -> [#accept]
    Goto:
      '{' -> State 5
      ID -> State 7
      ERROR -> State 6
      atom -> State 8
      expr -> State 10
      block -> State 9

  State 4:
    Kernel Items:
      #accept: ^ block., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      block: '{'.expr_list '}'
    Reduce:
      * -> [expr_list]
    Goto:
      expr_list -> State 11

  State 6:
    Kernel Items:
      atom: ERROR., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      atom: ID., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      expr: atom., *
    Reduce:
      * -> [expr]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      atom: block., *
    Reduce:
      * -> [atom]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      expr_list: expr_list expr., *
      expr: expr.op atom
    Reduce:
      * -> [expr_list]
    Goto:
      '+' -> State 12
      '-' -> State 13
      op -> State 14

  State 11:
    Kernel Items:
      expr_list: expr_list.expr
      block: '{' expr_list.'}'
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      '}' -> State 15
      ID -> State 7
      ERROR -> State 6
      atom -> State 8
      expr -> State 10
      block -> State 9

  State 12:
    Kernel Items:
      op: '+'., *
    Reduce:
      * -> [op]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      op: '-'., *
    Reduce:
      * -> [op]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      expr: expr op.atom
    Reduce:
      (nil)
    Goto:
      '{' -> State 5
      ID -> State 7
      ERROR -> State 6
      atom -> State 16
      block -> State 9

  State 15:
    Kernel Items:
      block: '{' expr_list '}'., $
    Reduce:
      $ -> [block]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      expr: expr op atom., *
    Reduce:
      * -> [expr]
    Goto:
      (nil)

Number of states: 16
Number of shift actions: 25
Number of reduce actions: 13
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
