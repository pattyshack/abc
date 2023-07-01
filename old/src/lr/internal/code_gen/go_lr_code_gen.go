package code_gen

import (
	"fmt"
	"io"
	"sort"

	"github.com/pattyshack/abc/src/lr/codegenutil"
	lr "github.com/pattyshack/abc/src/lr/internal"
	"github.com/pattyshack/abc/src/lr/internal/code_gen/go_template"
	"github.com/pattyshack/abc/src/lr/internal/parser"
)

var (
	escapedChar = map[string]byte{
		"'\\t'":  '\t',
		"'\\n'":  '\n',
		"'\\''":  '\'',
		"'\\\\'": '\\',
	}
)

type NameGenerator struct {
    prefix string

    nameCount map[string]int
}

func NewNameGenerator(prefix string) *NameGenerator {
    return &NameGenerator{
        prefix: prefix,
        nameCount: map[string]int{},
    }
}

func (ng *NameGenerator) Add(name string) string {
    ng.nameCount[name] += 1
    cnt := ng.nameCount[name]
    if cnt > 1 {
        name = fmt.Sprintf("%s_%d", name, cnt)
    }
    return name
}

func (ng *NameGenerator) Public(name string) string {
    name = ng.prefix + name
    return ng.Add(name)
}

func (ng *NameGenerator) Internal(name string) string {
    name = "_" + ng.prefix + name
    return ng.Add(name)
}

func populateCodeGenVariables(
    prefix string,
    Terms map[string]*lr.Term,
    OrderedStates []*lr.ItemSet,
    valueTypes map[string]*lr.Param,
    nameGen *NameGenerator) error {

	for _, term := range Terms {
		valueType := valueTypes[term.ValueType]
		if valueType == nil {
            return fmt.Errorf(
                "Undefined value type for <%s> %s",
                term.ValueType,
                term.LRLocation)
		}
		term.CodeGenType = valueType.ParamType

		if term.SymbolId == parser.LRCharacterToken {
			term.CodeGenSymbolConst = term.Name
		} else {
            suffix := "Type"
            if term.IsTerminal {
                suffix = "Token"
            }
		    term.CodeGenSymbolConst = nameGen.Public(
                go_template.SnakeToCamel(term.Name) + suffix)
		}

		for _, clause := range term.Clauses {
			reducerName := nameGen.Add(
                go_template.SnakeToCamel(clause.Label) +
                "To" +
				go_template.SnakeToCamel(term.Name))
			clause.CodeGenReducerName = reducerName
			clause.CodeGenReducerNameConst = nameGen.Internal(
                "Reduce" + reducerName)
			clause.CodeGenReduceAction = nameGen.Internal(
                "Reduce" + reducerName + "Action")
		}
	}

	for _, state := range OrderedStates {
		state.CodeGenConst = nameGen.Internal(
            fmt.Sprintf("State%d", state.StateNum))
		state.CodeGenAction = nameGen.Internal(
            fmt.Sprintf("GotoState%dAction", state.StateNum))
	}

	return nil
}

func GenerateGoLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	io.WriterTo,
	error) {

	cfg := grammar.LangSpecs.Go
	if cfg == nil {
		return nil, fmt.Errorf("go configuration not specified in lang_specs")
	}

	if cfg.Package == "" {
		return nil, fmt.Errorf("package name not specified")
	}

    imports := codegenutil.NewGoImports()
    nameGen := NewNameGenerator(cfg.Prefix)

	endSymbol := nameGen.Internal("EndMarker")
	wildcardSymbol := nameGen.Internal("WildcardMarker")
	genericSymbol := nameGen.Public("GenericSymbol")
	genericSymbolPtr := "*" + genericSymbol

	orderedValueTypes := lr.ParamList{
		&lr.Param{lr.Generic, imports.Obj(genericSymbolPtr)},
	}
	for name, valueType := range cfg.ValueTypes {
		orderedValueTypes = append(
			orderedValueTypes,
			&lr.Param{name, imports.Obj(valueType)})
	}
	sort.Sort(orderedValueTypes)

    valueTypes := make(map[string]*lr.Param, len(orderedValueTypes))
    for _, vt := range orderedValueTypes {
        valueTypes[vt.Name] = vt
    }

	orderedSymbols := []*lr.Term{
		&lr.Term{
			Name:               lr.StartMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: lr.StartMarker,
			CodeGenType:        genericSymbolPtr,
		},
		&lr.Term{
			Name:               lr.Wildcard,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: wildcardSymbol,
			CodeGenType:        genericSymbolPtr,
		},
		&lr.Term{
			Name:               lr.EndMarker,
			IsTerminal:         true,
			ValueType:          lr.Generic,
			CodeGenSymbolConst: endSymbol,
			CodeGenType:        genericSymbolPtr,
		},
	}
	orderedSymbols = append(orderedSymbols, grammar.Terminals...)
	orderedSymbols = append(orderedSymbols, grammar.NonTerminals...)

	symbols := make(map[string]*lr.Term, len(orderedSymbols))
	for _, symbol := range orderedSymbols {
		symbols[symbol.Name] = symbol
	}

	file := &go_template.File{
		Source:                    grammar.Source,
		Package:                   cfg.Package,
		Imports:                   imports,
		ActionType:                nameGen.Internal("Action"),
		ActionIdType:              nameGen.Internal("ActionType"),
		ShiftAction:               nameGen.Internal("ShiftAction"),
		ReduceAction:              nameGen.Internal("ReduceAction"),
		AcceptAction:              nameGen.Internal("AcceptAction"),
		StateIdType:               nameGen.Internal("StateId"),
		ReduceType:                nameGen.Internal("ReduceType"),
		SymbolType:                nameGen.Public("Symbol"),
		GenericSymbolType:         genericSymbol,
		StackItemType:             nameGen.Internal("StackItem"),
		StackType:                 nameGen.Internal("Stack"),
		SymbolStackType:           nameGen.Internal("PseudoSymbolStack"),
		SymbolIdType:              nameGen.Public("SymbolId"),
		EndSymbolId:               endSymbol,
		WildcardSymbolId:          wildcardSymbol,
		LocationType:              nameGen.Public("Location"),
		TokenType:                 nameGen.Public("Token"),
		LexerType:                 nameGen.Public("Lexer"),
		ReducerType:               nameGen.Public("Reducer"),
		ErrHandlerType:            nameGen.Public("ParseErrorHandler"),
		DefaultErrHandlerType:     nameGen.Public("DefaultParseErrorHandler"),
		ExpectedTerminalsFunc:     nameGen.Public("ExpectedTerminals"),
		ParseFuncPrefix:           nameGen.Public("Parse"),
		InternalParseFunc:         nameGen.Internal("Parse"),
		TableKeyType:              nameGen.Internal("ActionTableKey"),
		ActionTableType:           nameGen.Internal("ActionTableType"),
		ActionTable:               nameGen.Internal("ActionTable"),
		SortSlice:                 imports.Obj("sort.Slice"),
		Sprintf:                   imports.Obj("fmt.Sprintf"),
		Errorf:                    imports.Obj("fmt.Errorf"),
		EOF:                       imports.Obj("io.EOF"),
		OrderedSymbols:            orderedSymbols,
		Symbols:                   symbols,
		StartSymbols:              grammar.Starts,
		OrderedStates:             states.OrderedStates,
		OrderedValueTypes:         orderedValueTypes,
		OutputDebugNonKernelItems: cfg.OutputDebugNonKernelItems,
	}

	err := populateCodeGenVariables(
        cfg.Prefix,
        grammar.Terms,
        states.OrderedStates,
        valueTypes,
        nameGen)
	if err != nil {
		return nil, err
	}

	return codegenutil.NewFormattedGoSource(file), nil
}
