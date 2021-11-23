package code_gen

import (
	"fmt"
	"sort"
	"strings"

	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/go_template"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

var escapedChar = map[string]byte{
	"'\\t'":  '\t',
	"'\\n'":  '\n',
	"'\\''":  '\'',
	"'\\\\'": '\\',
}

// TODO handle this more gracefully
func snakeToCamel(str string) string {
	chunks := strings.Split(str, "_")

	result := ""
	for _, chunk := range chunks {
		result += strings.Title(strings.ToLower(chunk))
	}

	return result
}

type goCodeGen struct {
	*lr.Grammar
	*lr.GoSpec

	*lr.LRStates

	*GoCodeBuilder

	nameLocs map[string]parser.LRLocation

	location string

	symbolId string
	stateId  string

	symbol string

	endSymbol      string
	wildcardSymbol string

	token   string
	lexer   string
	reducer string

	errHandler        string
	defaultErrHandler string
	expectedTerminals string

	genericSymbol string

	symbolStack string

	stackItem string
	stack     string

	reduceType string
	actionType string

	shiftAction  string
	reduceAction string
	acceptAction string

	action string

	tableKey        string
	actionTableType string
	actionTable     string

	parse string
}

func newGoCodeGen(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	*goCodeGen,
	error) {

	cfg := grammar.LangSpecs.Go
	if cfg == nil {
		return nil, fmt.Errorf("go configuration not specified in lang_specs")
	}

	if cfg.Package == "" {
		return nil, fmt.Errorf("package name not specified")
	}

	builder := NewGoCodeBuilder(cfg.Package)

	builder.HeaderBoilerplate.Line(
		"// Auto-generated from source: %s",
		grammar.Source)
	builder.HeaderBoilerplate.Line("")

	return &goCodeGen{
		Grammar:       grammar,
		GoSpec:        cfg,
		LRStates:      states,
		GoCodeBuilder: builder,
		nameLocs:      map[string]parser.LRLocation{},
	}, nil
}

func (gen *goCodeGen) check(name string, loc parser.LRLocation) error {
	prev, ok := gen.nameLocs[name]
	if ok {
		return fmt.Errorf(
			"Generated conflicting name: %s (from %s and %s)",
			name,
			prev,
			loc)
	}
	gen.nameLocs[name] = loc

	return nil
}

func (gen *goCodeGen) populateCodeGenVariables() error {
	gen.location = gen.Prefix + "Location"
	gen.symbolId = gen.Prefix + "SymbolId"
	gen.symbol = gen.Prefix + "Symbol"
	gen.stateId = "_" + gen.Prefix + "StateId"
	gen.endSymbol = "_" + gen.Prefix + "EndMarker"
	gen.wildcardSymbol = "_" + gen.Prefix + "WildcardMarker"
	gen.token = gen.Prefix + "Token"
	gen.lexer = gen.Prefix + "Lexer"
	gen.reducer = gen.Prefix + "Reducer"
	gen.errHandler = gen.Prefix + "ParseErrorHandler"
	gen.defaultErrHandler = gen.Prefix + "DefaultParseErrorHandler"
	gen.expectedTerminals = "_" + gen.Prefix + "ExpectedTerminals"
	gen.genericSymbol = gen.Prefix + "GenericSymbol"
	gen.symbolStack = "_" + gen.Prefix + "PseudoSymbolStack"
	gen.stackItem = "_" + gen.Prefix + "StackItem"
	gen.stack = "_" + gen.Prefix + "Stack"
	gen.reduceType = "_" + gen.Prefix + "ReduceType"
	gen.actionType = "_" + gen.Prefix + "ActionType"
	gen.shiftAction = "_" + gen.Prefix + "ShiftAction"
	gen.reduceAction = "_" + gen.Prefix + "ReduceAction"
	gen.acceptAction = "_" + gen.Prefix + "AcceptAction"
	gen.action = "_" + gen.Prefix + "Action"
	gen.tableKey = "_" + gen.Prefix + "ActionTableKey"
	gen.actionTableType = "_" + gen.Prefix + "ActionTableType"
	gen.actionTable = "_" + gen.Prefix + "ActionTable"
	gen.parse = "_" + gen.Prefix + "Parse"

	for _, term := range gen.Terms {
		valueType := gen.ValueTypes[term.ValueType]
		if valueType == "" {
			if term.ValueType != lr.Generic {
				return fmt.Errorf(
					"Undefined value type for <%s> %s",
					term.ValueType,
					term.LRLocation)
			}
			valueType = "*" + gen.genericSymbol
		}

		term.CodeGenType = gen.Obj(valueType)

		symbolConst := ""
		if term.SymbolId == parser.LRCharacterToken {
			symbolConst = term.Name
		} else {
			symbolConst = gen.Prefix + snakeToCamel(term.Name)
			if term.IsTerminal {
				symbolConst += "Token"
			} else {
				symbolConst += "Type"
			}
		}

		err := gen.check(symbolConst, term.LRLocation)
		if err != nil {
			return err
		}

		term.CodeGenSymbolConst = symbolConst

		for _, clause := range term.Clauses {
			reducerName := snakeToCamel(clause.Label) + "To" +
				snakeToCamel(term.Name)

			err := gen.check(reducerName, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReducerName = reducerName

			reducerConst := "_" + gen.Prefix + "Reduce" + reducerName
			err = gen.check(reducerConst, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReducerNameConst = reducerConst

			actionConst := reducerConst + "Action"
			err = gen.check(actionConst, clause.LRLocation)
			if err != nil {
				return err
			}

			clause.CodeGenReduceAction = actionConst
		}
	}

	for _, state := range gen.OrderedStates {
		state.CodeGenConst = fmt.Sprintf(
			"_%sState%d",
			gen.Prefix,
			state.StateNum)

		actionConst := fmt.Sprintf(
			"_%sGotoState%dAction",
			gen.Prefix,
			state.StateNum)
		err := gen.check(actionConst, parser.LRLocation{})
		if err != nil {
			return err
		}

		state.CodeGenAction = actionConst
	}

	return nil
}

func (gen *goCodeGen) generateExpectedTerminals() {
	l := gen.Line

	idToConst := map[string]string{
		lr.EndMarker: gen.endSymbol,
	}
	symbols := []string{}

	for _, term := range gen.Terminals {
		idToConst[term.Name] = term.CodeGenSymbolConst
		symbols = append(symbols, term.Name)
	}

	l("var %s = map[%s][]%s{", gen.expectedTerminals, gen.stateId, gen.symbolId)
	gen.PushIndent()

	for _, state := range gen.OrderedStates {
		consts := []string{}

		for _, symbol := range symbols {
			_, ok := state.Goto[symbol]
			if !ok {
				continue
			}
			consts = append(consts, idToConst[symbol])
		}

		for _, item := range state.Items {
			if item.IsReduce && item.LookAhead != lr.Wildcard {
				consts = append(consts, idToConst[item.LookAhead])
			}
		}

		if len(consts) > 0 {
			l("%s: []%s{%s},",
				state.CodeGenConst,
				gen.symbolId,
				strings.Join(consts, ", "))
		}
	}

	gen.PopIndent()
	l("}")
	l("")
}

func GenerateGoLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	*GoCodeBuilder,
	error) {

	gen, err := newGoCodeGen(grammar, states)
	if err != nil {
		return nil, err
	}

	err = gen.populateCodeGenVariables()
	if err != nil {
		return nil, err
	}

	orderedValueTypes := lr.ParamList{
		&lr.Param{lr.Generic, gen.Obj("*" + gen.genericSymbol)},
	}
	for name, valueType := range gen.ValueTypes {
		orderedValueTypes = append(
			orderedValueTypes,
			&lr.Param{name, gen.Obj(valueType)})
	}
	sort.Sort(orderedValueTypes)

	genericPtr := "*" + gen.genericSymbol
	orderedSymbols := []*lr.Term{
		&lr.Term{
			Name:               lr.StartMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: lr.StartMarker,
			CodeGenType:        genericPtr,
		},
		&lr.Term{
			Name:               lr.EndMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: gen.endSymbol,
			CodeGenType:        genericPtr,
		},
		&lr.Term{
			Name:               lr.Wildcard,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: gen.wildcardSymbol,
			CodeGenType:        genericPtr,
		},
	}
	orderedSymbols = append(orderedSymbols, gen.Terminals...)
	orderedSymbols = append(orderedSymbols, gen.NonTerminals...)

	symbols := make(map[string]*lr.Term, len(orderedSymbols))
	for _, symbol := range orderedSymbols {
		symbols[symbol.Name] = symbol
	}

	gen.Embed(
		&go_template.PublicDefinitions{
			LocationType:      gen.location,
			SymbolIdType:      gen.symbolId,
			SymbolType:        gen.token,
			GenericSymbolType: gen.genericSymbol,
			LexerType:         gen.lexer,
			ReducerType:       gen.reducer,
			ErrHandler:        gen.errHandler,
			DefaultErrHandler: gen.defaultErrHandler,
			StackType:         gen.stack,
			ExpectedTerminals: gen.expectedTerminals,
			ParsePrefix:       gen.Prefix + "Parse",
			InternalParse:     gen.parse,
			Sprintf:           gen.Obj("fmt.Sprintf"),
			Errorf:            gen.Obj("fmt.Errorf"),
			Terminals:         gen.Terminals,
			NonTerminals:      gen.NonTerminals,
			Starts:            gen.Starts,
			OrderedStates:     gen.OrderedStates,
		})

	l := gen.Line
	l("// ================================================================")
	l("// Parser internal implementation")
	l("// User should normally avoid directly accessing the following code")
	l("// ================================================================")
	l("")

	gen.Embed(
		&go_template.ParseFunc{
			ParseFuncName:   gen.parse,
			LexerType:       gen.lexer,
			ReducerType:     gen.reducer,
			ErrHandlerType:  gen.errHandler,
			StateIdType:     gen.stateId,
			SymbolType:      gen.symbol,
			StackItemType:   gen.stackItem,
			StackType:       gen.stack,
			SymbolStackType: gen.symbolStack,
			ActionTable:     gen.actionTable,
			AcceptAction:    gen.acceptAction,
			ShiftAction:     gen.shiftAction,
			ReduceAction:    gen.reduceAction,
		})

	gen.Embed(
		&go_template.InternalDefinitions{
			ActionType:        gen.action,
			ActionIdType:      gen.actionType,
			ShiftAction:      gen.shiftAction,
			ReduceAction:     gen.reduceAction,
			AcceptAction:     gen.acceptAction,
			StateIdType:       gen.stateId,
			ReduceType:        gen.reduceType,
			SymbolType:        gen.symbol,
			GenericSymbolType: gen.genericSymbol,
			StackItemType:     gen.stackItem,
			StackType:         gen.stack,
			LexerType:         gen.lexer,
			ReducerType:       gen.reducer,
			SymbolStackType:   gen.symbolStack,
			SymbolIdType:      gen.symbolId,
			EndSymbolId:       gen.endSymbol,
			WildcardSymbolId:       gen.wildcardSymbol,
			LocationType:      gen.location,
			TokenType:         gen.token,
			Sprintf:           gen.Obj("fmt.Sprintf"),
			Errorf:            gen.Obj("fmt.Errorf"),
			EOF:               gen.Obj("io.EOF"),
			OrderedSymbols:    orderedSymbols,
			OrderedStates:     gen.OrderedStates,
			OrderedValueTypes: orderedValueTypes,
		})

	// Maybe make the action table map[StateId]map[Symbol]Action and
	// extract the expected terminal symbols from the action table
	gen.generateExpectedTerminals()

	gen.Embed(
		&go_template.ActionTable{
			TableKeyType:     gen.tableKey,
			StateIdType:      gen.stateId,
			SymbolIdType:     gen.symbolId,
			EndSymbolId:      gen.endSymbol,
			WildcardSymbolId: gen.wildcardSymbol,
			ActionTableType:  gen.actionTableType,
			ActionType:       gen.action,
			AcceptAction:     gen.acceptAction,
			ShiftAction:      gen.shiftAction,
			ReduceAction:     gen.reduceAction,
			OrderedStates:    gen.OrderedStates,
			ActionTable:      gen.actionTable,
			OrderedSymbols:   orderedSymbols,
			Symbols:          symbols,
		})

	l("")

	gen.Embed(
		&go_template.DebugStates{
			OutputDebugNonKernelItems: gen.OutputDebugNonKernelItems,
			OrderedSymbols:            orderedSymbols,
			OrderedStates:             gen.OrderedStates,
		})

	return gen.GoCodeBuilder, nil
}
