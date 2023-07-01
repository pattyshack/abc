package parser

import (
	"fmt"
	"strings"
)

type Definition interface {
	Loc() LRLocation
	String() string
}

var _ LRToken = &Token{}

type Token struct {
	LRLocation

	LRSymbolId
	Value string
}

func (t *Token) Id() LRSymbolId {
	return t.LRSymbolId
}

func (t *Token) Loc() LRLocation {
	return t.LRLocation
}

func (t *Token) String() string {
	return fmt.Sprintf("%v: %s (%v)", t.LRSymbolId, t.Value, t.Loc())
}

type StartDeclaration struct {
	LRLocation

	Ids []*Token
}

func NewStartDeclaration(
	start *LRGenericSymbol,
	ids []*Token) *StartDeclaration {

	return &StartDeclaration{
		LRLocation: start.LRLocation,
		Ids:        ids,
	}
}

func (sd *StartDeclaration) Loc() LRLocation {
	return sd.LRLocation
}

func (sd *StartDeclaration) String() string {
	names := []string{}
	for _, id := range sd.Ids {
		names = append(names, id.Value)
	}
	return "%start " + strings.Join(names, " ")
}

type TermDeclaration struct {
	TermType *LRGenericSymbol

	IsTerminal bool

	ValueType *Token

	Terms []*Token
}

func NewTermDeclaration(
	termType *LRGenericSymbol,
	valueType *Token,
	terms []*Token) *TermDeclaration {

	return &TermDeclaration{
		TermType:   termType,
		IsTerminal: termType.Id() == LRTokenToken,
		ValueType:  valueType,
		Terms:      terms,
	}
}

func (td *TermDeclaration) Loc() LRLocation {
	return td.TermType.LRLocation
}

func (td *TermDeclaration) String() string {
	result := TypeKeyword
	if td.IsTerminal {
		result = TokenKeyword
	}

	if td.ValueType != nil {
		result += " <" + td.ValueType.Value + ">"
	}

	for _, term := range td.Terms {
		result += " " + term.Value
	}

	return result
}

type Clause struct {
	Label *Token // optional
	Body  []*Token

	// set by NewRule
	LRLocation
	Parent *Rule

	SortId int
}

func NewClause(label *Token, body []*Token) *Clause {
	return &Clause{
		Label: label,
		Body:  body,
	}
}

func (clause *Clause) String() string {
	result := "(none):"
	if clause.Label != nil {
		result = clause.Label.Value + ":"
	}

	for _, item := range clause.Body {
		result += " " + item.Value
	}

	return result
}

type Rule struct {
	Name    *Token
	Clauses []*Clause
}

func NewRule(name *Token, clauses []*Clause) *Rule {
	rule := &Rule{
		Name:    name,
		Clauses: clauses,
	}

	for _, clause := range clauses {
		loc := name.LRLocation
		if clause.Label != nil {
			loc = clause.Label.LRLocation
		} else if len(clause.Body) > 0 {
			loc = clause.Body[0].LRLocation
		}

		clause.LRLocation = loc
		clause.Parent = rule
	}

	return rule
}

func (r *Rule) Loc() LRLocation {
	return r.Name.LRLocation
}

func (r *Rule) String() string {
	result := r.Name.Value + " ->"
	for i, clause := range r.Clauses {
		result += "\n    " + clause.String()
		if i < len(r.Clauses)-1 {
			result += " |"
		}
	}

	return result
}

type AdditionalSection struct {
	Name    *Token
	Content *Token
}

func NewAdditionalSection(name *Token, content *Token) *AdditionalSection {
	return &AdditionalSection{
		Name:    name,
		Content: content,
	}
}

type Grammar struct {
	Definitions []Definition

	AdditionalSections []*AdditionalSection
}

func NewGrammar(
	defs []Definition,
	additionalSections []*AdditionalSection) *Grammar {

	return &Grammar{defs, additionalSections}
}
