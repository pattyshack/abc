package lr

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pattyshack/abc/src/lr/internal/parser"

	"gopkg.in/yaml.v2"
)

const (
	Generic = "Generic_"
)

type Clause struct {
	SortId int // 0 (and negatives) are reserved for the start rules.
	parser.LRLocation

	Label string

	Bindings []*Term

	// Temp variable populated by the code generator.
	CodeGenReducerName      string
	CodeGenReducerNameConst string
	CodeGenReduceAction     string
}

type Term struct {
	Name string
	parser.LRLocation

	SymbolId parser.LRSymbolId

	IsTerminal bool

	ValueType string

	RuleLocation parser.LRLocation
	Clauses      []*Clause

	Reachable bool

	CodeGenSymbolConst string

	// Temp variable populated by code generator.  The generated type is
	// language specific.
	CodeGenType interface{}
}

type Grammar struct {
	Source string

	Terms map[string]*Term

	Terminals    []*Term // sorted by declaration location
	NonTerminals []*Term // sorted by rule location

	Starts []*Term

	*LangSpecs
}

func classifyDefinitions(
	parsed []parser.Definition) (
	map[string]*Term,
	map[string]*parser.Rule,
	[]string, // start rule(s)
	[]string) { // error strings

	terms := map[string]*Term{}

	rules := map[string]*parser.Rule{}

	firstRuleName := ""
	var start *parser.StartDeclaration

	errStrs := []string{}

	sortId := 1
	for _, d := range parsed {
		switch def := d.(type) {
		case *parser.StartDeclaration:
			if start != nil {
				errStrs = append(
					errStrs,
					fmt.Sprintf(
						"Duplicate start declaration: %s %s",
						start.Loc().ShortString(),
						def.Loc().ShortString()))
			}

			start = def

		case *parser.TermDeclaration:
			for _, term := range def.Terms {
				prev, ok := terms[term.Value]
				if ok {
					errStrs = append(
						errStrs,
						fmt.Sprintf(
							"Duplicate token/type declaration: %s %s %s",
							term.Value,
							prev.LRLocation.ShortString(),
							term.LRLocation.ShortString()))
				}

				valueType := Generic
				if def.ValueType != nil {
					valueType = def.ValueType.Value
				}

				terms[term.Value] = &Term{
					Name:       term.Value,
					LRLocation: term.LRLocation,
					SymbolId:   term.Id(),
					IsTerminal: def.IsTerminal,
					ValueType:  valueType,
					Reachable:  false,
				}
			}

		case *parser.Rule:
			if firstRuleName == "" {
				firstRuleName = def.Name.Value
			}

			prev, ok := rules[def.Name.Value]
			if ok {
				errStrs = append(
					errStrs,
					fmt.Sprintf(
						"Duplicate rule: %s %s %s",
						def.Name.Value,
						prev.Loc().ShortString(),
						def.Loc().ShortString()))
			}
			rules[def.Name.Value] = def

			for _, clause := range def.Clauses {
				clause.SortId = sortId
				sortId += 1
			}
		}
	}

	startRules := []string{}
	if start != nil {
		ids := map[string]parser.LRLocation{}
		for _, id := range start.Ids {
			prev, ok := ids[id.Value]
			if ok {
				errStrs = append(
					errStrs,
					fmt.Sprintf(
						"Duplicate start entry: %s %s %s",
						id.Value,
						prev.ShortString(),
						id.Loc().ShortString()))
			} else {
				ids[id.Value] = id.Loc()
				startRules = append(startRules, id.Value)
			}
		}
	} else {
		startRules = append(startRules, firstRuleName)
	}

	return terms, rules, startRules, errStrs
}

func bindTerms(
	terms map[string]*Term,
	rules map[string]*parser.Rule,
	startRuleNames []string) (
	[]*Term,
	[]string) {

	errStrs := []string{}

	for name, rule := range rules {
		term, ok := terms[name]
		if !ok {
			errStrs = append(
				errStrs,
				fmt.Sprintf("Undefined type: %s %v", name, rule.Loc()))
			continue
		}

		term.RuleLocation = rule.Loc()

		clauses := []*Clause{}
		for _, parsedClause := range rule.Clauses {
			label := ""
			if parsedClause.Label != nil {
				label = parsedClause.Label.Value
			}
			clause := &Clause{
				SortId:     parsedClause.SortId,
				LRLocation: parsedClause.LRLocation,
				Label:      label,
				Bindings:   []*Term{},
			}

			for _, id_or_char := range parsedClause.Body {
				t, ok := terms[id_or_char.Value]
				if ok {
					clause.Bindings = append(clause.Bindings, t)
				} else {
					errStrs = append(
						errStrs,
						fmt.Sprintf(
							"Undefined token/type: %s %v",
							id_or_char.Value,
							id_or_char.Loc()))
				}
			}

			clauses = append(clauses, clause)
		}

		term.Clauses = clauses
	}

	for name, term := range terms {
		rule, ok := rules[name]
		if !term.IsTerminal && !ok {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"No rule specified for type: %s %v",
					name,
					term.LRLocation))
		} else if term.IsTerminal && ok {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"token cannot have associated rule: %s %v",
					name,
					rule.Loc()))
		}
	}

	startTerms := []*Term{}
	for _, name := range startRuleNames {
		startTerm, ok := terms[name]
		if !ok || startTerm.IsTerminal {
			errStrs = append(
				errStrs,
				fmt.Sprintf("Invalid start rule: %s", name))
		} else {
			startTerms = append(startTerms, startTerm)
		}
	}

	return startTerms, errStrs
}

func checkReachability(starts []*Term, terms map[string]*Term) []string {
	if len(starts) == 0 {
		return nil
	}

	exploreSet := map[string]*Term{}
	for _, start := range starts {
		exploreSet[start.Name] = start
	}

	for len(exploreSet) > 0 {
		nextExploreSet := map[string]*Term{}

		for _, term := range exploreSet {
			if term.Reachable {
				continue
			}
			term.Reachable = true

			for _, clause := range term.Clauses {
				for _, item := range clause.Bindings {
					if !item.Reachable {
						nextExploreSet[item.Name] = item
					}
				}
			}
		}

		exploreSet = nextExploreSet
	}

	errStrs := []string{}
	for _, term := range terms {
		if !term.Reachable {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Unused token/type. Not reachable from start rule: %s %v",
					term.Name,
					term.LRLocation))
		}
	}

	return errStrs
}

func extractLangSpecs(
	sections []*parser.AdditionalSection) (
	*LangSpecs,
	[]string) {

	errStrs := []string{}

	var langSpecsSection *parser.AdditionalSection
	for _, section := range sections {
		if section.Name.Value != "lang_specs" {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Unexpected additional section: %s %v",
					section.Name.Value,
					section.Name.Loc()))
			continue
		}

		if langSpecsSection != nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf(
					"Duplicated lang_specs section specified: %v %v",
					langSpecsSection.Name.Loc(),
					section.Name.Loc()))
		}

		langSpecsSection = section
	}

	langSpecs := &LangSpecs{}
	if langSpecsSection != nil {
		err := yaml.Unmarshal([]byte(langSpecsSection.Content.Value), langSpecs)
		if err != nil {
			errStrs = append(
				errStrs,
				fmt.Sprintf("Failed to unmarshal lang_specs: %s", err))
		}
	}

	return langSpecs, errStrs
}

func NewGrammar(
	sourceFile string,
	parsed *parser.Grammar) (
	*Grammar,
	error) {

	terms, rules, startRuleNames, errStrs := classifyDefinitions(
		parsed.Definitions)

	if len(rules) == 0 {
		errStrs = append(errStrs, "No rules specified in grammar.")
	}

	startTerms, bindErrStrs := bindTerms(terms, rules, startRuleNames)
	errStrs = append(errStrs, bindErrStrs...)

	errStrs = append(errStrs, checkReachability(startTerms, terms)...)

	langSpecs, asErrStrs := extractLangSpecs(parsed.AdditionalSections)
	errStrs = append(errStrs, asErrStrs...)

	if len(errStrs) > 0 {
		return nil, fmt.Errorf(strings.Join(errStrs, "\n"))
	}

	terminals := []*Term{}
	nonTerminals := []*Term{}

	for _, term := range terms {
		if term.IsTerminal {
			terminals = append(terminals, term)
		} else {
			nonTerminals = append(nonTerminals, term)
		}
	}

	sort.Sort(ByDeclLoc(terminals))
	sort.Sort(ByRuleLoc(nonTerminals))

	return &Grammar{
		Source:       sourceFile,
		Terms:        terms,
		Terminals:    terminals,
		NonTerminals: nonTerminals,
		Starts:       startTerms,
		LangSpecs:    langSpecs,
	}, nil
}
