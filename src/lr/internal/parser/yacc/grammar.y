%{
package yacc

import (
    "github.com/pattyshack/abc/src/lr/internal/parser"
)
%}

%union {
    Generic_ *parser.LRGenericSymbol

    *parser.Token
    Tokens []*parser.Token

    parser.Definition  // interface
    Definitions []parser.Definition

    *parser.Rule

    Clause *parser.Clause
    Clauses []*parser.Clause

    *parser.AdditionalSection
    AdditionalSections []*parser.AdditionalSection

    *parser.Grammar
}

// yacc input syntax:
// https://docs.oracle.com/cd/E19504-01/802-5880/yacc-19/index.html

// XXX: add LEFT / RIGHT / NONASSOC?
%token <Generic_> TOKEN TYPE START // %<identifier>

// Intermediate token that should not reach the parser
%token <Token> ARROW COLON

// <identifier> followed by -> (ignoring whitespace and comment), tokenized as a
// single token by the lexer.  Equivalent to C_IDENTIFIER in yacc
%token <Token> RULE_DEF

// <identifier> followed by : (ignoring whitespace and comment), tokenized as
// a single token by the lexer.
%token <Token> LABEL

%token <Generic_> LT GT OR SEMICOLON SECTION_MARKER
%token <Token> IDENTIFIER CHARACTER

%token <Token> SECTION_CONTENT

%type <Generic_> rword
%type <Tokens> nonempty_ident_list rule_body

%type <Definition> def
%type <Definitions> defs

%type <Rule> rule

%type <Clause> labeled_clause
%type <Clauses> labeled_clauses

%type <AdditionalSection> additional_section
%type <AdditionalSections> additional_sections

%type <Grammar>  grammar

%start grammar

%%

// NOTE: there's no tail and section separator, line terminator ';' is optional

grammar:
    defs additional_sections {
        Lrlex.(*ParseContext).Grammar, _ = Lrlex.(*ParseContext).ToGrammar($1, $2)
    }
    ;

additional_sections:
    additional_sections additional_section {
        $$, _ = Lrlex.(*ParseContext).AddToAdditionalSections($1, $2)
    }
    |
    {
        $$, _ = Lrlex.(*ParseContext).NilToAdditionalSections()
    }
    ;

additional_section:
    SECTION_MARKER IDENTIFIER SECTION_CONTENT {
        $$, _ = Lrlex.(*ParseContext).ToAdditionalSection($1, $2, $3)
    }
    ;

defs:
    defs def {
        $$, _ =  Lrlex.(*ParseContext).AddToDefs($1, $2)
    }
    |
    defs def SEMICOLON {
        $$, _ =  Lrlex.(*ParseContext).AddExplicitToDefs($1, $2, $3)
    }
    |
    def {
        $$, _ =  Lrlex.(*ParseContext).DefToDefs($1)
    }
    |
    def SEMICOLON {
        $$, _ =  Lrlex.(*ParseContext).ExplicitDefToDefs($1, $2)
    }
    ;

// TODO: handle language specific boiler plate, union/struct
def:
    // type / token declaration
    rword LT IDENTIFIER GT nonempty_ident_list {
        $$, _ =  Lrlex.(*ParseContext).TermDeclToDef($1, $2, $3, $4, $5)
    }
    |
    rword nonempty_ident_list {
        $$, _ =  Lrlex.(*ParseContext).UntypedTermDeclToDef($1, $2)
    }
    |
    // start declaration
    START nonempty_ident_list {
        $$, _ =  Lrlex.(*ParseContext).StartDeclToDef($1, $2)
    }
    | rule {
        $$, _ =  Lrlex.(*ParseContext).RuleToDef($1)
    }
    ;

rword:
    TOKEN {
        $$, _ =  Lrlex.(*ParseContext).TokenToRword($1)
    }
    |
    TYPE {
        $$, _ =  Lrlex.(*ParseContext).TypeToRword($1)
    }
    ;

nonempty_ident_list:
    nonempty_ident_list IDENTIFIER {
        $$, _ = Lrlex.(*ParseContext).AddToNonemptyIdentList($1, $2)
    }
    |
    IDENTIFIER {
        $$, _ = Lrlex.(*ParseContext).IdentToNonemptyIdentList($1)
    }
    ;

rule:
    RULE_DEF rule_body {
        $$, _ = Lrlex.(*ParseContext).UnlabeledClauseToRule($1, $2)
    }
    |
    RULE_DEF labeled_clauses {
        $$, _ = Lrlex.(*ParseContext).ClausesToRule($1, $2)
    }
    ;

rule_body:
    rule_body IDENTIFIER {
        $$, _ = Lrlex.(*ParseContext).AddIdToRuleBody($1, $2)
    }
    |
    rule_body CHARACTER {
        $$, _ = Lrlex.(*ParseContext).AddCharToRuleBody($1, $2)
    }
    | {
        $$, _ = Lrlex.(*ParseContext).NilToRuleBody()
    }
    ;

labeled_clauses:
    labeled_clauses OR labeled_clause {
        $$, _ = Lrlex.(*ParseContext).AddToLabeledClauses($1, $2, $3)
    }
    |
    labeled_clause {
        $$, _ = Lrlex.(*ParseContext).ClauseToLabeledClauses($1)
    }
    ;

labeled_clause:
    LABEL rule_body {
        $$, _ = Lrlex.(*ParseContext).ToLabeledClause($1, $2)
    }
    ;

%%

func init() {
    LrErrorVerbose = true
}

