// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type LRSymbolId int

const (
	LRTokenToken          = LRSymbolId(256)
	LRTypeToken           = LRSymbolId(257)
	LRStartToken          = LRSymbolId(258)
	LRRuleDefToken        = LRSymbolId(259)
	LRLabelToken          = LRSymbolId(260)
	LRSectionMarkerToken  = LRSymbolId(261)
	LRCharacterToken      = LRSymbolId(262)
	LRIdentifierToken     = LRSymbolId(263)
	LRSectionContentToken = LRSymbolId(264)
)

type LRLocation struct {
	FileName string
	Line     int
	Column   int
}

func (l LRLocation) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l LRLocation) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type LRToken interface {
	Id() LRSymbolId
	Loc() LRLocation
}

type LRGenericSymbol struct {
	LRSymbolId
	LRLocation
}

func (t *LRGenericSymbol) Id() LRSymbolId { return t.LRSymbolId }

func (t *LRGenericSymbol) Loc() LRLocation { return t.LRLocation }

type LRLexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *LRGenericSymbol
	Next() (LRToken, error)

	CurrentLocation() LRLocation
}

type LRReducer interface {
	// 42:11: grammar -> ...
	ToGrammar(Defs_ []Definition, AdditionalSections_ []*AdditionalSection) (*Grammar, error)

	// 45:4: additional_sections -> add: ...
	AddToAdditionalSections(AdditionalSections_ []*AdditionalSection, AdditionalSection_ *AdditionalSection) ([]*AdditionalSection, error)

	// 46:4: additional_sections -> nil: ...
	NilToAdditionalSections() ([]*AdditionalSection, error)

	// 48:22: additional_section -> ...
	ToAdditionalSection(SectionMarker_ *LRGenericSymbol, Identifier_ *Token, SectionContent_ *Token) (*AdditionalSection, error)

	// 51:4: defs -> add: ...
	AddToDefs(Defs_ []Definition, Def_ Definition) ([]Definition, error)

	// 52:4: defs -> add_explicit: ...
	AddExplicitToDefs(Defs_ []Definition, Def_ Definition, char *LRGenericSymbol) ([]Definition, error)

	// 53:4: defs -> def: ...
	DefToDefs(Def_ Definition) ([]Definition, error)

	// 54:4: defs -> explicit_def: ...
	ExplicitDefToDefs(Def_ Definition, char *LRGenericSymbol) ([]Definition, error)

	// 59:4: def -> term_decl: ...
	TermDeclToDef(Rword_ *LRGenericSymbol, char *LRGenericSymbol, Identifier_ *Token, char2 *LRGenericSymbol, NonemptyIdOrCharList_ []*Token) (Definition, error)

	// 60:4: def -> untyped_term_decl: ...
	UntypedTermDeclToDef(Rword_ *LRGenericSymbol, NonemptyIdOrCharList_ []*Token) (Definition, error)

	// 62:4: def -> start_decl: ...
	StartDeclToDef(Start_ *LRGenericSymbol, NonemptyIdentList_ []*Token) (Definition, error)

	// 63:4: def -> rule: ...
	RuleToDef(Rule_ *Rule) (Definition, error)

	// 66:4: rword -> token: ...
	TokenToRword(Token_ *LRGenericSymbol) (*LRGenericSymbol, error)

	// 67:4: rword -> type: ...
	TypeToRword(Type_ *LRGenericSymbol) (*LRGenericSymbol, error)

	// 70:4: nonempty_ident_list -> add: ...
	AddToNonemptyIdentList(NonemptyIdentList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 71:4: nonempty_ident_list -> ident: ...
	IdentToNonemptyIdentList(Identifier_ *Token) ([]*Token, error)

	// 74:4: nonempty_id_or_char_list -> add_id: ...
	AddIdToNonemptyIdOrCharList(NonemptyIdOrCharList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 75:4: nonempty_id_or_char_list -> add_char: ...
	AddCharToNonemptyIdOrCharList(NonemptyIdOrCharList_ []*Token, Character_ *Token) ([]*Token, error)

	// 76:4: nonempty_id_or_char_list -> id: ...
	IdToNonemptyIdOrCharList(Identifier_ *Token) ([]*Token, error)

	// 77:4: nonempty_id_or_char_list -> char: ...
	CharToNonemptyIdOrCharList(Character_ *Token) ([]*Token, error)

	// 80:4: id_or_char_list -> list: ...
	ListToIdOrCharList(NonemptyIdOrCharList_ []*Token) ([]*Token, error)

	// 81:4: id_or_char_list -> nil: ...
	NilToIdOrCharList() ([]*Token, error)

	// 84:4: rule -> unlabeled_clause: ...
	UnlabeledClauseToRule(RuleDef_ *Token, IdOrCharList_ []*Token) (*Rule, error)

	// 85:4: rule -> clauses: ...
	ClausesToRule(RuleDef_ *Token, LabeledClauses_ []*Clause) (*Rule, error)

	// 88:4: labeled_clauses -> add: ...
	AddToLabeledClauses(LabeledClauses_ []*Clause, char *LRGenericSymbol, LabeledClause_ *Clause) ([]*Clause, error)

	// 89:4: labeled_clauses -> clause: ...
	ClauseToLabeledClauses(LabeledClause_ *Clause) ([]*Clause, error)

	// 91:18: labeled_clause -> ...
	ToLabeledClause(Label_ *Token, IdOrCharList_ []*Token) (*Clause, error)
}

type LRParseErrorHandler interface {
	Error(nextToken LRToken, parseStack _LRStack) error
}

type LRDefaultParseErrorHandler struct{}

func (LRDefaultParseErrorHandler) Error(nextToken LRToken, stack _LRStack) error {
	return fmt.Errorf(
		"Syntax error: unexpected symbol %v. Expecting %v (%v)",
		nextToken.Id(),
		LRExpectedTerminals(stack[len(stack)-1].StateId),
		nextToken.Loc())
}

func LRExpectedTerminals(id _LRStateId) []LRSymbolId {
	result := []LRSymbolId{}
	for key, _ := range _LRActionTable {
		if key._LRStateId != id {
			continue
		}
		result = append(result, key.LRSymbolId)
	}

	sort.Slice(result, func(i int, j int) bool { return result[i] < result[j] })
	return result
}

func LRParse(lexer LRLexer, reducer LRReducer) (*Grammar, error) {

	return LRParseWithCustomErrorHandler(
		lexer,
		reducer,
		LRDefaultParseErrorHandler{})
}

func LRParseWithCustomErrorHandler(
	lexer LRLexer,
	reducer LRReducer,
	errHandler LRParseErrorHandler) (
	*Grammar,
	error) {

	item, err := _LRParse(lexer, reducer, errHandler, _LRState1)
	if err != nil {
		var errRetVal *Grammar
		return errRetVal, err
	}
	return item.Grammar, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _LRParse(
	lexer LRLexer,
	reducer LRReducer,
	errHandler LRParseErrorHandler,
	startState _LRStateId) (
	*_LRStackItem,
	error) {

	stateStack := _LRStack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_LRStackItem{startState, nil},
	}

	symbolStack := &_LRPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _LRActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _LRShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _LRReduceAction {
			var reduceSymbol *LRSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _LRAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i LRSymbolId) String() string {
	switch i {
	case _LREndMarker:
		return "$"
	case _LRWildcardMarker:
		return "*"
	case LRTokenToken:
		return "TOKEN"
	case LRTypeToken:
		return "TYPE"
	case LRStartToken:
		return "START"
	case LRRuleDefToken:
		return "RULE_DEF"
	case LRLabelToken:
		return "LABEL"
	case '<':
		return "'<'"
	case '>':
		return "'>'"
	case '|':
		return "'|'"
	case ';':
		return "';'"
	case LRSectionMarkerToken:
		return "SECTION_MARKER"
	case LRCharacterToken:
		return "CHARACTER"
	case LRIdentifierToken:
		return "IDENTIFIER"
	case LRSectionContentToken:
		return "SECTION_CONTENT"
	case LRGrammarType:
		return "grammar"
	case LRAdditionalSectionsType:
		return "additional_sections"
	case LRAdditionalSectionType:
		return "additional_section"
	case LRDefsType:
		return "defs"
	case LRDefType:
		return "def"
	case LRRwordType:
		return "rword"
	case LRNonemptyIdentListType:
		return "nonempty_ident_list"
	case LRNonemptyIdOrCharListType:
		return "nonempty_id_or_char_list"
	case LRIdOrCharListType:
		return "id_or_char_list"
	case LRRuleType:
		return "rule"
	case LRLabeledClausesType:
		return "labeled_clauses"
	case LRLabeledClauseType:
		return "labeled_clause"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_LREndMarker      = LRSymbolId(0)
	_LRWildcardMarker = LRSymbolId(-1)

	LRGrammarType              = LRSymbolId(269)
	LRAdditionalSectionsType   = LRSymbolId(270)
	LRAdditionalSectionType    = LRSymbolId(271)
	LRDefsType                 = LRSymbolId(272)
	LRDefType                  = LRSymbolId(273)
	LRRwordType                = LRSymbolId(274)
	LRNonemptyIdentListType    = LRSymbolId(275)
	LRNonemptyIdOrCharListType = LRSymbolId(276)
	LRIdOrCharListType         = LRSymbolId(277)
	LRRuleType                 = LRSymbolId(278)
	LRLabeledClausesType       = LRSymbolId(279)
	LRLabeledClauseType        = LRSymbolId(280)
)

type _LRActionType int

const (
	// NOTE: error action is implicit
	_LRShiftAction  = _LRActionType(0)
	_LRReduceAction = _LRActionType(1)
	_LRAcceptAction = _LRActionType(2)
)

func (i _LRActionType) String() string {
	switch i {
	case _LRShiftAction:
		return "shift"
	case _LRReduceAction:
		return "reduce"
	case _LRAcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _LRReduceType int

const (
	_LRReduceToGrammar                     = _LRReduceType(1)
	_LRReduceAddToAdditionalSections       = _LRReduceType(2)
	_LRReduceNilToAdditionalSections       = _LRReduceType(3)
	_LRReduceToAdditionalSection           = _LRReduceType(4)
	_LRReduceAddToDefs                     = _LRReduceType(5)
	_LRReduceAddExplicitToDefs             = _LRReduceType(6)
	_LRReduceDefToDefs                     = _LRReduceType(7)
	_LRReduceExplicitDefToDefs             = _LRReduceType(8)
	_LRReduceTermDeclToDef                 = _LRReduceType(9)
	_LRReduceUntypedTermDeclToDef          = _LRReduceType(10)
	_LRReduceStartDeclToDef                = _LRReduceType(11)
	_LRReduceRuleToDef                     = _LRReduceType(12)
	_LRReduceTokenToRword                  = _LRReduceType(13)
	_LRReduceTypeToRword                   = _LRReduceType(14)
	_LRReduceAddToNonemptyIdentList        = _LRReduceType(15)
	_LRReduceIdentToNonemptyIdentList      = _LRReduceType(16)
	_LRReduceAddIdToNonemptyIdOrCharList   = _LRReduceType(17)
	_LRReduceAddCharToNonemptyIdOrCharList = _LRReduceType(18)
	_LRReduceIdToNonemptyIdOrCharList      = _LRReduceType(19)
	_LRReduceCharToNonemptyIdOrCharList    = _LRReduceType(20)
	_LRReduceListToIdOrCharList            = _LRReduceType(21)
	_LRReduceNilToIdOrCharList             = _LRReduceType(22)
	_LRReduceUnlabeledClauseToRule         = _LRReduceType(23)
	_LRReduceClausesToRule                 = _LRReduceType(24)
	_LRReduceAddToLabeledClauses           = _LRReduceType(25)
	_LRReduceClauseToLabeledClauses        = _LRReduceType(26)
	_LRReduceToLabeledClause               = _LRReduceType(27)
)

func (i _LRReduceType) String() string {
	switch i {
	case _LRReduceToGrammar:
		return "ToGrammar"
	case _LRReduceAddToAdditionalSections:
		return "AddToAdditionalSections"
	case _LRReduceNilToAdditionalSections:
		return "NilToAdditionalSections"
	case _LRReduceToAdditionalSection:
		return "ToAdditionalSection"
	case _LRReduceAddToDefs:
		return "AddToDefs"
	case _LRReduceAddExplicitToDefs:
		return "AddExplicitToDefs"
	case _LRReduceDefToDefs:
		return "DefToDefs"
	case _LRReduceExplicitDefToDefs:
		return "ExplicitDefToDefs"
	case _LRReduceTermDeclToDef:
		return "TermDeclToDef"
	case _LRReduceUntypedTermDeclToDef:
		return "UntypedTermDeclToDef"
	case _LRReduceStartDeclToDef:
		return "StartDeclToDef"
	case _LRReduceRuleToDef:
		return "RuleToDef"
	case _LRReduceTokenToRword:
		return "TokenToRword"
	case _LRReduceTypeToRword:
		return "TypeToRword"
	case _LRReduceAddToNonemptyIdentList:
		return "AddToNonemptyIdentList"
	case _LRReduceIdentToNonemptyIdentList:
		return "IdentToNonemptyIdentList"
	case _LRReduceAddIdToNonemptyIdOrCharList:
		return "AddIdToNonemptyIdOrCharList"
	case _LRReduceAddCharToNonemptyIdOrCharList:
		return "AddCharToNonemptyIdOrCharList"
	case _LRReduceIdToNonemptyIdOrCharList:
		return "IdToNonemptyIdOrCharList"
	case _LRReduceCharToNonemptyIdOrCharList:
		return "CharToNonemptyIdOrCharList"
	case _LRReduceListToIdOrCharList:
		return "ListToIdOrCharList"
	case _LRReduceNilToIdOrCharList:
		return "NilToIdOrCharList"
	case _LRReduceUnlabeledClauseToRule:
		return "UnlabeledClauseToRule"
	case _LRReduceClausesToRule:
		return "ClausesToRule"
	case _LRReduceAddToLabeledClauses:
		return "AddToLabeledClauses"
	case _LRReduceClauseToLabeledClauses:
		return "ClauseToLabeledClauses"
	case _LRReduceToLabeledClause:
		return "ToLabeledClause"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _LRStateId int

func (id _LRStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_LRState1  = _LRStateId(1)
	_LRState2  = _LRStateId(2)
	_LRState3  = _LRStateId(3)
	_LRState4  = _LRStateId(4)
	_LRState5  = _LRStateId(5)
	_LRState6  = _LRStateId(6)
	_LRState7  = _LRStateId(7)
	_LRState8  = _LRStateId(8)
	_LRState9  = _LRStateId(9)
	_LRState10 = _LRStateId(10)
	_LRState11 = _LRStateId(11)
	_LRState12 = _LRStateId(12)
	_LRState13 = _LRStateId(13)
	_LRState14 = _LRStateId(14)
	_LRState15 = _LRStateId(15)
	_LRState16 = _LRStateId(16)
	_LRState17 = _LRStateId(17)
	_LRState18 = _LRStateId(18)
	_LRState19 = _LRStateId(19)
	_LRState20 = _LRStateId(20)
	_LRState21 = _LRStateId(21)
	_LRState22 = _LRStateId(22)
	_LRState23 = _LRStateId(23)
	_LRState24 = _LRStateId(24)
	_LRState25 = _LRStateId(25)
	_LRState26 = _LRStateId(26)
	_LRState27 = _LRStateId(27)
	_LRState28 = _LRStateId(28)
	_LRState29 = _LRStateId(29)
	_LRState30 = _LRStateId(30)
	_LRState31 = _LRStateId(31)
	_LRState32 = _LRStateId(32)
	_LRState33 = _LRStateId(33)
	_LRState34 = _LRStateId(34)
	_LRState35 = _LRStateId(35)
	_LRState36 = _LRStateId(36)
	_LRState37 = _LRStateId(37)
	_LRState38 = _LRStateId(38)
)

type LRSymbol struct {
	SymbolId_ LRSymbolId

	Generic_ *LRGenericSymbol

	AdditionalSection  *AdditionalSection
	AdditionalSections []*AdditionalSection
	Clause             *Clause
	Clauses            []*Clause
	Definition         Definition
	Definitions        []Definition
	Grammar            *Grammar
	Rule               *Rule
	Token              *Token
	Tokens             []*Token
}

func NewSymbol(token LRToken) (*LRSymbol, error) {
	symbol, ok := token.(*LRSymbol)
	if ok {
		return symbol, nil
	}

	symbol = &LRSymbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _LREndMarker, LRTokenToken, LRTypeToken, LRStartToken, '<', '>', '|', ';', LRSectionMarkerToken:
		val, ok := token.(*LRGenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *LRGenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	case LRRuleDefToken, LRLabelToken, LRCharacterToken, LRIdentifierToken, LRSectionContentToken:
		val, ok := token.(*Token)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *Token (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Token = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *LRSymbol) Id() LRSymbolId {
	return s.SymbolId_
}

func (s *LRSymbol) Loc() LRLocation {
	type locator interface{ Loc() LRLocation }
	switch s.SymbolId_ {
	case LRAdditionalSectionType:
		loc, ok := interface{}(s.AdditionalSection).(locator)
		if ok {
			return loc.Loc()
		}
	case LRAdditionalSectionsType:
		loc, ok := interface{}(s.AdditionalSections).(locator)
		if ok {
			return loc.Loc()
		}
	case LRLabeledClauseType:
		loc, ok := interface{}(s.Clause).(locator)
		if ok {
			return loc.Loc()
		}
	case LRLabeledClausesType:
		loc, ok := interface{}(s.Clauses).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefType:
		loc, ok := interface{}(s.Definition).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefsType:
		loc, ok := interface{}(s.Definitions).(locator)
		if ok {
			return loc.Loc()
		}
	case LRGrammarType:
		loc, ok := interface{}(s.Grammar).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleType:
		loc, ok := interface{}(s.Rule).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleDefToken, LRLabelToken, LRCharacterToken, LRIdentifierToken, LRSectionContentToken:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Loc()
		}
	case LRNonemptyIdentListType, LRNonemptyIdOrCharListType, LRIdOrCharListType:
		loc, ok := interface{}(s.Tokens).(locator)
		if ok {
			return loc.Loc()
		}
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return LRLocation{}
}

type _LRPseudoSymbolStack struct {
	lexer LRLexer
	top   []*LRSymbol
}

func (stack *_LRPseudoSymbolStack) Top() (*LRSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &LRGenericSymbol{_LREndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_LRPseudoSymbolStack) Push(symbol *LRSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_LRPseudoSymbolStack) Pop() (*LRSymbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _LRStackItem struct {
	StateId _LRStateId

	*LRSymbol
}

type _LRStack []*_LRStackItem

type _LRAction struct {
	ActionType _LRActionType

	ShiftStateId _LRStateId
	ReduceType   _LRReduceType
}

func (act *_LRAction) ShiftItem(symbol *LRSymbol) *_LRStackItem {
	return &_LRStackItem{StateId: act.ShiftStateId, LRSymbol: symbol}
}

func (act *_LRAction) ReduceSymbol(
	reducer LRReducer,
	stack _LRStack) (
	_LRStack,
	*LRSymbol,
	error) {

	var err error
	symbol := &LRSymbol{}
	switch act.ReduceType {
	case _LRReduceToGrammar:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRGrammarType
		symbol.Grammar, err = reducer.ToGrammar(args[0].Definitions, args[1].AdditionalSections)
	case _LRReduceAddToAdditionalSections:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.AddToAdditionalSections(args[0].AdditionalSections, args[1].AdditionalSection)
	case _LRReduceNilToAdditionalSections:
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.NilToAdditionalSections()
	case _LRReduceToAdditionalSection:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRAdditionalSectionType
		symbol.AdditionalSection, err = reducer.ToAdditionalSection(args[0].Generic_, args[1].Token, args[2].Token)
	case _LRReduceAddToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddToDefs(args[0].Definitions, args[1].Definition)
	case _LRReduceAddExplicitToDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddExplicitToDefs(args[0].Definitions, args[1].Definition, args[2].Generic_)
	case _LRReduceDefToDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.DefToDefs(args[0].Definition)
	case _LRReduceExplicitDefToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.ExplicitDefToDefs(args[0].Definition, args[1].Generic_)
	case _LRReduceTermDeclToDef:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.TermDeclToDef(args[0].Generic_, args[1].Generic_, args[2].Token, args[3].Generic_, args[4].Tokens)
	case _LRReduceUntypedTermDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.UntypedTermDeclToDef(args[0].Generic_, args[1].Tokens)
	case _LRReduceStartDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.StartDeclToDef(args[0].Generic_, args[1].Tokens)
	case _LRReduceRuleToDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.RuleToDef(args[0].Rule)
	case _LRReduceTokenToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		symbol.Generic_, err = reducer.TokenToRword(args[0].Generic_)
	case _LRReduceTypeToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		symbol.Generic_, err = reducer.TypeToRword(args[0].Generic_)
	case _LRReduceAddToNonemptyIdentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.AddToNonemptyIdentList(args[0].Tokens, args[1].Token)
	case _LRReduceIdentToNonemptyIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.IdentToNonemptyIdentList(args[0].Token)
	case _LRReduceAddIdToNonemptyIdOrCharList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.AddIdToNonemptyIdOrCharList(args[0].Tokens, args[1].Token)
	case _LRReduceAddCharToNonemptyIdOrCharList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.AddCharToNonemptyIdOrCharList(args[0].Tokens, args[1].Token)
	case _LRReduceIdToNonemptyIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.IdToNonemptyIdOrCharList(args[0].Token)
	case _LRReduceCharToNonemptyIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.CharToNonemptyIdOrCharList(args[0].Token)
	case _LRReduceListToIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRIdOrCharListType
		symbol.Tokens, err = reducer.ListToIdOrCharList(args[0].Tokens)
	case _LRReduceNilToIdOrCharList:
		symbol.SymbolId_ = LRIdOrCharListType
		symbol.Tokens, err = reducer.NilToIdOrCharList()
	case _LRReduceUnlabeledClauseToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRRuleType
		symbol.Rule, err = reducer.UnlabeledClauseToRule(args[0].Token, args[1].Tokens)
	case _LRReduceClausesToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRRuleType
		symbol.Rule, err = reducer.ClausesToRule(args[0].Token, args[1].Clauses)
	case _LRReduceAddToLabeledClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRLabeledClausesType
		symbol.Clauses, err = reducer.AddToLabeledClauses(args[0].Clauses, args[1].Generic_, args[2].Clause)
	case _LRReduceClauseToLabeledClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRLabeledClausesType
		symbol.Clauses, err = reducer.ClauseToLabeledClauses(args[0].Clause)
	case _LRReduceToLabeledClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRLabeledClauseType
		symbol.Clause, err = reducer.ToLabeledClause(args[0].Token, args[1].Tokens)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _LRActionTableKey struct {
	_LRStateId
	LRSymbolId
}

type _LRActionTableType map[_LRActionTableKey]*_LRAction

func (table _LRActionTableType) Get(
	stateId _LRStateId,
	symbolId LRSymbolId) (
	*_LRAction,
	bool) {

	action, ok := table[_LRActionTableKey{stateId, symbolId}]
	if ok {
		return action, ok
	}

	action, ok = table[_LRActionTableKey{stateId, _LRWildcardMarker}]
	return action, ok
}

var (
	_LRGotoState1Action                          = &_LRAction{_LRShiftAction, _LRState1, 0}
	_LRGotoState2Action                          = &_LRAction{_LRShiftAction, _LRState2, 0}
	_LRGotoState3Action                          = &_LRAction{_LRShiftAction, _LRState3, 0}
	_LRGotoState4Action                          = &_LRAction{_LRShiftAction, _LRState4, 0}
	_LRGotoState5Action                          = &_LRAction{_LRShiftAction, _LRState5, 0}
	_LRGotoState6Action                          = &_LRAction{_LRShiftAction, _LRState6, 0}
	_LRGotoState7Action                          = &_LRAction{_LRShiftAction, _LRState7, 0}
	_LRGotoState8Action                          = &_LRAction{_LRShiftAction, _LRState8, 0}
	_LRGotoState9Action                          = &_LRAction{_LRShiftAction, _LRState9, 0}
	_LRGotoState10Action                         = &_LRAction{_LRShiftAction, _LRState10, 0}
	_LRGotoState11Action                         = &_LRAction{_LRShiftAction, _LRState11, 0}
	_LRGotoState12Action                         = &_LRAction{_LRShiftAction, _LRState12, 0}
	_LRGotoState13Action                         = &_LRAction{_LRShiftAction, _LRState13, 0}
	_LRGotoState14Action                         = &_LRAction{_LRShiftAction, _LRState14, 0}
	_LRGotoState15Action                         = &_LRAction{_LRShiftAction, _LRState15, 0}
	_LRGotoState16Action                         = &_LRAction{_LRShiftAction, _LRState16, 0}
	_LRGotoState17Action                         = &_LRAction{_LRShiftAction, _LRState17, 0}
	_LRGotoState18Action                         = &_LRAction{_LRShiftAction, _LRState18, 0}
	_LRGotoState19Action                         = &_LRAction{_LRShiftAction, _LRState19, 0}
	_LRGotoState20Action                         = &_LRAction{_LRShiftAction, _LRState20, 0}
	_LRGotoState21Action                         = &_LRAction{_LRShiftAction, _LRState21, 0}
	_LRGotoState22Action                         = &_LRAction{_LRShiftAction, _LRState22, 0}
	_LRGotoState23Action                         = &_LRAction{_LRShiftAction, _LRState23, 0}
	_LRGotoState24Action                         = &_LRAction{_LRShiftAction, _LRState24, 0}
	_LRGotoState25Action                         = &_LRAction{_LRShiftAction, _LRState25, 0}
	_LRGotoState26Action                         = &_LRAction{_LRShiftAction, _LRState26, 0}
	_LRGotoState27Action                         = &_LRAction{_LRShiftAction, _LRState27, 0}
	_LRGotoState28Action                         = &_LRAction{_LRShiftAction, _LRState28, 0}
	_LRGotoState29Action                         = &_LRAction{_LRShiftAction, _LRState29, 0}
	_LRGotoState30Action                         = &_LRAction{_LRShiftAction, _LRState30, 0}
	_LRGotoState31Action                         = &_LRAction{_LRShiftAction, _LRState31, 0}
	_LRGotoState32Action                         = &_LRAction{_LRShiftAction, _LRState32, 0}
	_LRGotoState33Action                         = &_LRAction{_LRShiftAction, _LRState33, 0}
	_LRGotoState34Action                         = &_LRAction{_LRShiftAction, _LRState34, 0}
	_LRGotoState35Action                         = &_LRAction{_LRShiftAction, _LRState35, 0}
	_LRGotoState36Action                         = &_LRAction{_LRShiftAction, _LRState36, 0}
	_LRGotoState37Action                         = &_LRAction{_LRShiftAction, _LRState37, 0}
	_LRGotoState38Action                         = &_LRAction{_LRShiftAction, _LRState38, 0}
	_LRReduceToGrammarAction                     = &_LRAction{_LRReduceAction, 0, _LRReduceToGrammar}
	_LRReduceAddToAdditionalSectionsAction       = &_LRAction{_LRReduceAction, 0, _LRReduceAddToAdditionalSections}
	_LRReduceNilToAdditionalSectionsAction       = &_LRAction{_LRReduceAction, 0, _LRReduceNilToAdditionalSections}
	_LRReduceToAdditionalSectionAction           = &_LRAction{_LRReduceAction, 0, _LRReduceToAdditionalSection}
	_LRReduceAddToDefsAction                     = &_LRAction{_LRReduceAction, 0, _LRReduceAddToDefs}
	_LRReduceAddExplicitToDefsAction             = &_LRAction{_LRReduceAction, 0, _LRReduceAddExplicitToDefs}
	_LRReduceDefToDefsAction                     = &_LRAction{_LRReduceAction, 0, _LRReduceDefToDefs}
	_LRReduceExplicitDefToDefsAction             = &_LRAction{_LRReduceAction, 0, _LRReduceExplicitDefToDefs}
	_LRReduceTermDeclToDefAction                 = &_LRAction{_LRReduceAction, 0, _LRReduceTermDeclToDef}
	_LRReduceUntypedTermDeclToDefAction          = &_LRAction{_LRReduceAction, 0, _LRReduceUntypedTermDeclToDef}
	_LRReduceStartDeclToDefAction                = &_LRAction{_LRReduceAction, 0, _LRReduceStartDeclToDef}
	_LRReduceRuleToDefAction                     = &_LRAction{_LRReduceAction, 0, _LRReduceRuleToDef}
	_LRReduceTokenToRwordAction                  = &_LRAction{_LRReduceAction, 0, _LRReduceTokenToRword}
	_LRReduceTypeToRwordAction                   = &_LRAction{_LRReduceAction, 0, _LRReduceTypeToRword}
	_LRReduceAddToNonemptyIdentListAction        = &_LRAction{_LRReduceAction, 0, _LRReduceAddToNonemptyIdentList}
	_LRReduceIdentToNonemptyIdentListAction      = &_LRAction{_LRReduceAction, 0, _LRReduceIdentToNonemptyIdentList}
	_LRReduceAddIdToNonemptyIdOrCharListAction   = &_LRAction{_LRReduceAction, 0, _LRReduceAddIdToNonemptyIdOrCharList}
	_LRReduceAddCharToNonemptyIdOrCharListAction = &_LRAction{_LRReduceAction, 0, _LRReduceAddCharToNonemptyIdOrCharList}
	_LRReduceIdToNonemptyIdOrCharListAction      = &_LRAction{_LRReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}
	_LRReduceCharToNonemptyIdOrCharListAction    = &_LRAction{_LRReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}
	_LRReduceListToIdOrCharListAction            = &_LRAction{_LRReduceAction, 0, _LRReduceListToIdOrCharList}
	_LRReduceNilToIdOrCharListAction             = &_LRAction{_LRReduceAction, 0, _LRReduceNilToIdOrCharList}
	_LRReduceUnlabeledClauseToRuleAction         = &_LRAction{_LRReduceAction, 0, _LRReduceUnlabeledClauseToRule}
	_LRReduceClausesToRuleAction                 = &_LRAction{_LRReduceAction, 0, _LRReduceClausesToRule}
	_LRReduceAddToLabeledClausesAction           = &_LRAction{_LRReduceAction, 0, _LRReduceAddToLabeledClauses}
	_LRReduceClauseToLabeledClausesAction        = &_LRAction{_LRReduceAction, 0, _LRReduceClauseToLabeledClauses}
	_LRReduceToLabeledClauseAction               = &_LRAction{_LRReduceAction, 0, _LRReduceToLabeledClause}
)

var _LRActionTable = _LRActionTableType{
	{_LRState2, _LREndMarker}:                &_LRAction{_LRAcceptAction, 0, 0},
	{_LRState1, LRTokenToken}:                _LRGotoState5Action,
	{_LRState1, LRTypeToken}:                 _LRGotoState6Action,
	{_LRState1, LRStartToken}:                _LRGotoState4Action,
	{_LRState1, LRRuleDefToken}:              _LRGotoState3Action,
	{_LRState1, LRGrammarType}:               _LRGotoState2Action,
	{_LRState1, LRDefsType}:                  _LRGotoState8Action,
	{_LRState1, LRDefType}:                   _LRGotoState7Action,
	{_LRState1, LRRwordType}:                 _LRGotoState10Action,
	{_LRState1, LRRuleType}:                  _LRGotoState9Action,
	{_LRState3, LRLabelToken}:                _LRGotoState13Action,
	{_LRState3, LRCharacterToken}:            _LRGotoState11Action,
	{_LRState3, LRIdentifierToken}:           _LRGotoState12Action,
	{_LRState3, LRNonemptyIdOrCharListType}:  _LRGotoState17Action,
	{_LRState3, LRIdOrCharListType}:          _LRGotoState14Action,
	{_LRState3, LRLabeledClausesType}:        _LRGotoState16Action,
	{_LRState3, LRLabeledClauseType}:         _LRGotoState15Action,
	{_LRState4, LRIdentifierToken}:           _LRGotoState18Action,
	{_LRState4, LRNonemptyIdentListType}:     _LRGotoState19Action,
	{_LRState7, ';'}:                         _LRGotoState20Action,
	{_LRState8, LRTokenToken}:                _LRGotoState5Action,
	{_LRState8, LRTypeToken}:                 _LRGotoState6Action,
	{_LRState8, LRStartToken}:                _LRGotoState4Action,
	{_LRState8, LRRuleDefToken}:              _LRGotoState3Action,
	{_LRState8, LRAdditionalSectionsType}:    _LRGotoState21Action,
	{_LRState8, LRDefType}:                   _LRGotoState22Action,
	{_LRState8, LRRwordType}:                 _LRGotoState10Action,
	{_LRState8, LRRuleType}:                  _LRGotoState9Action,
	{_LRState10, '<'}:                        _LRGotoState23Action,
	{_LRState10, LRCharacterToken}:           _LRGotoState11Action,
	{_LRState10, LRIdentifierToken}:          _LRGotoState12Action,
	{_LRState10, LRNonemptyIdOrCharListType}: _LRGotoState24Action,
	{_LRState13, LRCharacterToken}:           _LRGotoState11Action,
	{_LRState13, LRIdentifierToken}:          _LRGotoState12Action,
	{_LRState13, LRNonemptyIdOrCharListType}: _LRGotoState17Action,
	{_LRState13, LRIdOrCharListType}:         _LRGotoState25Action,
	{_LRState16, '|'}:                        _LRGotoState26Action,
	{_LRState17, LRCharacterToken}:           _LRGotoState27Action,
	{_LRState17, LRIdentifierToken}:          _LRGotoState28Action,
	{_LRState19, LRIdentifierToken}:          _LRGotoState29Action,
	{_LRState21, LRSectionMarkerToken}:       _LRGotoState30Action,
	{_LRState21, LRAdditionalSectionType}:    _LRGotoState31Action,
	{_LRState22, ';'}:                        _LRGotoState32Action,
	{_LRState23, LRIdentifierToken}:          _LRGotoState33Action,
	{_LRState24, LRCharacterToken}:           _LRGotoState27Action,
	{_LRState24, LRIdentifierToken}:          _LRGotoState28Action,
	{_LRState26, LRLabelToken}:               _LRGotoState13Action,
	{_LRState26, LRLabeledClauseType}:        _LRGotoState34Action,
	{_LRState30, LRIdentifierToken}:          _LRGotoState35Action,
	{_LRState33, '>'}:                        _LRGotoState36Action,
	{_LRState35, LRSectionContentToken}:      _LRGotoState37Action,
	{_LRState36, LRCharacterToken}:           _LRGotoState11Action,
	{_LRState36, LRIdentifierToken}:          _LRGotoState12Action,
	{_LRState36, LRNonemptyIdOrCharListType}: _LRGotoState38Action,
	{_LRState38, LRCharacterToken}:           _LRGotoState27Action,
	{_LRState38, LRIdentifierToken}:          _LRGotoState28Action,
	{_LRState3, _LRWildcardMarker}:           _LRReduceNilToIdOrCharListAction,
	{_LRState5, _LRWildcardMarker}:           _LRReduceTokenToRwordAction,
	{_LRState6, _LRWildcardMarker}:           _LRReduceTypeToRwordAction,
	{_LRState7, _LRWildcardMarker}:           _LRReduceDefToDefsAction,
	{_LRState8, _LRWildcardMarker}:           _LRReduceNilToAdditionalSectionsAction,
	{_LRState9, _LRWildcardMarker}:           _LRReduceRuleToDefAction,
	{_LRState11, _LRWildcardMarker}:          _LRReduceCharToNonemptyIdOrCharListAction,
	{_LRState12, _LRWildcardMarker}:          _LRReduceIdToNonemptyIdOrCharListAction,
	{_LRState13, _LRWildcardMarker}:          _LRReduceNilToIdOrCharListAction,
	{_LRState14, _LRWildcardMarker}:          _LRReduceUnlabeledClauseToRuleAction,
	{_LRState15, _LRWildcardMarker}:          _LRReduceClauseToLabeledClausesAction,
	{_LRState16, _LRWildcardMarker}:          _LRReduceClausesToRuleAction,
	{_LRState17, _LRWildcardMarker}:          _LRReduceListToIdOrCharListAction,
	{_LRState18, _LRWildcardMarker}:          _LRReduceIdentToNonemptyIdentListAction,
	{_LRState19, _LRWildcardMarker}:          _LRReduceStartDeclToDefAction,
	{_LRState20, _LRWildcardMarker}:          _LRReduceExplicitDefToDefsAction,
	{_LRState21, _LREndMarker}:               _LRReduceToGrammarAction,
	{_LRState22, _LRWildcardMarker}:          _LRReduceAddToDefsAction,
	{_LRState24, _LRWildcardMarker}:          _LRReduceUntypedTermDeclToDefAction,
	{_LRState25, _LRWildcardMarker}:          _LRReduceToLabeledClauseAction,
	{_LRState27, _LRWildcardMarker}:          _LRReduceAddCharToNonemptyIdOrCharListAction,
	{_LRState28, _LRWildcardMarker}:          _LRReduceAddIdToNonemptyIdOrCharListAction,
	{_LRState29, _LRWildcardMarker}:          _LRReduceAddToNonemptyIdentListAction,
	{_LRState31, _LRWildcardMarker}:          _LRReduceAddToAdditionalSectionsAction,
	{_LRState32, _LRWildcardMarker}:          _LRReduceAddExplicitToDefsAction,
	{_LRState34, _LRWildcardMarker}:          _LRReduceAddToLabeledClausesAction,
	{_LRState37, _LRWildcardMarker}:          _LRReduceToAdditionalSectionAction,
	{_LRState38, _LRWildcardMarker}:          _LRReduceTermDeclToDefAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.grammar
    Non-kernel Items:
      grammar:.defs additional_sections
      defs:.defs def
      defs:.defs def ';'
      defs:.def
      defs:.def ';'
      def:.rword '<' IDENTIFIER '>' nonempty_id_or_char_list
      def:.rword nonempty_id_or_char_list
      def:.START nonempty_ident_list
      def:.rule
      rword:.TOKEN
      rword:.TYPE
      rule:.RULE_DEF id_or_char_list
      rule:.RULE_DEF labeled_clauses
    Reduce:
      (nil)
    Goto:
      TOKEN -> State 5
      TYPE -> State 6
      START -> State 4
      RULE_DEF -> State 3
      grammar -> State 2
      defs -> State 8
      def -> State 7
      rword -> State 10
      rule -> State 9

  State 2:
    Kernel Items:
      #accept: ^ grammar., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 3:
    Kernel Items:
      rule: RULE_DEF.id_or_char_list
      rule: RULE_DEF.labeled_clauses
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
      id_or_char_list:.nonempty_id_or_char_list
      id_or_char_list:., *
      labeled_clauses:.labeled_clauses '|' labeled_clause
      labeled_clauses:.labeled_clause
      labeled_clause:.LABEL id_or_char_list
    Reduce:
      * -> [id_or_char_list]
    Goto:
      LABEL -> State 13
      CHARACTER -> State 11
      IDENTIFIER -> State 12
      nonempty_id_or_char_list -> State 17
      id_or_char_list -> State 14
      labeled_clauses -> State 16
      labeled_clause -> State 15

  State 4:
    Kernel Items:
      def: START.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
      nonempty_ident_list:.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 18
      nonempty_ident_list -> State 19

  State 5:
    Kernel Items:
      rword: TOKEN., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      rword: TYPE., *
    Reduce:
      * -> [rword]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      defs: def., *
      defs: def.';'
    Reduce:
      * -> [defs]
    Goto:
      ';' -> State 20

  State 8:
    Kernel Items:
      grammar: defs.additional_sections
      defs: defs.def
      defs: defs.def ';'
    Non-kernel Items:
      additional_sections:.additional_sections additional_section
      additional_sections:., *
      def:.rword '<' IDENTIFIER '>' nonempty_id_or_char_list
      def:.rword nonempty_id_or_char_list
      def:.START nonempty_ident_list
      def:.rule
      rword:.TOKEN
      rword:.TYPE
      rule:.RULE_DEF id_or_char_list
      rule:.RULE_DEF labeled_clauses
    Reduce:
      * -> [additional_sections]
    Goto:
      TOKEN -> State 5
      TYPE -> State 6
      START -> State 4
      RULE_DEF -> State 3
      additional_sections -> State 21
      def -> State 22
      rword -> State 10
      rule -> State 9

  State 9:
    Kernel Items:
      def: rule., *
    Reduce:
      * -> [def]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      def: rword.'<' IDENTIFIER '>' nonempty_id_or_char_list
      def: rword.nonempty_id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
    Reduce:
      (nil)
    Goto:
      '<' -> State 23
      CHARACTER -> State 11
      IDENTIFIER -> State 12
      nonempty_id_or_char_list -> State 24

  State 11:
    Kernel Items:
      nonempty_id_or_char_list: CHARACTER., *
    Reduce:
      * -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      nonempty_id_or_char_list: IDENTIFIER., *
    Reduce:
      * -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      labeled_clause: LABEL.id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
      id_or_char_list:.nonempty_id_or_char_list
      id_or_char_list:., *
    Reduce:
      * -> [id_or_char_list]
    Goto:
      CHARACTER -> State 11
      IDENTIFIER -> State 12
      nonempty_id_or_char_list -> State 17
      id_or_char_list -> State 25

  State 14:
    Kernel Items:
      rule: RULE_DEF id_or_char_list., *
    Reduce:
      * -> [rule]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      labeled_clauses: labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      rule: RULE_DEF labeled_clauses., *
      labeled_clauses: labeled_clauses.'|' labeled_clause
    Reduce:
      * -> [rule]
    Goto:
      '|' -> State 26

  State 17:
    Kernel Items:
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
      id_or_char_list: nonempty_id_or_char_list., *
    Reduce:
      * -> [id_or_char_list]
    Goto:
      CHARACTER -> State 27
      IDENTIFIER -> State 28

  State 18:
    Kernel Items:
      nonempty_ident_list: IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      def: START nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    Goto:
      IDENTIFIER -> State 29

  State 20:
    Kernel Items:
      defs: def ';'., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      grammar: defs additional_sections., $
      additional_sections: additional_sections.additional_section
    Non-kernel Items:
      additional_section:.SECTION_MARKER IDENTIFIER SECTION_CONTENT
    Reduce:
      $ -> [grammar]
    Goto:
      SECTION_MARKER -> State 30
      additional_section -> State 31

  State 22:
    Kernel Items:
      defs: defs def., *
      defs: defs def.';'
    Reduce:
      * -> [defs]
    Goto:
      ';' -> State 32

  State 23:
    Kernel Items:
      def: rword '<'.IDENTIFIER '>' nonempty_id_or_char_list
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 33

  State 24:
    Kernel Items:
      def: rword nonempty_id_or_char_list., *
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
    Reduce:
      * -> [def]
    Goto:
      CHARACTER -> State 27
      IDENTIFIER -> State 28

  State 25:
    Kernel Items:
      labeled_clause: LABEL id_or_char_list., *
    Reduce:
      * -> [labeled_clause]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      labeled_clauses: labeled_clauses '|'.labeled_clause
    Non-kernel Items:
      labeled_clause:.LABEL id_or_char_list
    Reduce:
      (nil)
    Goto:
      LABEL -> State 13
      labeled_clause -> State 34

  State 27:
    Kernel Items:
      nonempty_id_or_char_list: nonempty_id_or_char_list CHARACTER., *
    Reduce:
      * -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      nonempty_id_or_char_list: nonempty_id_or_char_list IDENTIFIER., *
    Reduce:
      * -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      nonempty_ident_list: nonempty_ident_list IDENTIFIER., *
    Reduce:
      * -> [nonempty_ident_list]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      additional_section: SECTION_MARKER.IDENTIFIER SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 35

  State 31:
    Kernel Items:
      additional_sections: additional_sections additional_section., *
    Reduce:
      * -> [additional_sections]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      defs: defs def ';'., *
    Reduce:
      * -> [defs]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      def: rword '<' IDENTIFIER.'>' nonempty_id_or_char_list
    Reduce:
      (nil)
    Goto:
      '>' -> State 36

  State 34:
    Kernel Items:
      labeled_clauses: labeled_clauses '|' labeled_clause., *
    Reduce:
      * -> [labeled_clauses]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER.SECTION_CONTENT
    Reduce:
      (nil)
    Goto:
      SECTION_CONTENT -> State 37

  State 36:
    Kernel Items:
      def: rword '<' IDENTIFIER '>'.nonempty_id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
    Reduce:
      (nil)
    Goto:
      CHARACTER -> State 11
      IDENTIFIER -> State 12
      nonempty_id_or_char_list -> State 38

  State 37:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER SECTION_CONTENT., *
    Reduce:
      * -> [additional_section]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      def: rword '<' IDENTIFIER '>' nonempty_id_or_char_list., *
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
    Reduce:
      * -> [def]
    Goto:
      CHARACTER -> State 27
      IDENTIFIER -> State 28

Number of states: 38
Number of shift actions: 55
Number of reduce actions: 29
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
