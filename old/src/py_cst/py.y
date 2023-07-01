%{

package py_cst

import (
)
%}

%union {
    intVal int

    token *Token  // lexer will only use Token
    tokens []*Token

    expression Expression
    expressions *ExprList  // (i) and (i,) have different meaning ...

    statement Statement
    statements []Statement

    argument *Argument
    arguments []*Argument
    argList *ArgumentList

    iterator *Iterator
    iterators []*Iterator

    // either []*Argument (Invocation), []*Subscript, or NAME (Field access)
    trailer interface{}
    trailers []interface{}

    decorator *Decorator
    decorators []*Decorator

    importClause *ImportClause
    importClauses []*ImportClause

    conditionClause *ConditionClause
    conditionClauses []*ConditionClause

    subscript *Subscript
    subscripts []*Subscript

    withClause *WithClause
    withClauses []*WithClause
}


/* Store all tokens as string to preserve the original value */
%token <token> FLOAT INTEGER STRING NAME

%token <token> NEWLINE LINE_CONTINUATION COMMENT_NEWLINE

%token <token> INDENT DEDENT /* python sadness to fake context free */

/* keywords */
%token <token> AND
%token <token> AS
%token <token> ASSERT
%token <token> BREAK
%token <token> CLASS
%token <token> CONTINUE
%token <token> DEF
%token <token> DEL
%token <token> ELIF
%token <token> ELSE
%token <token> EXCEPT
%token <token> EXEC
%token <token> FINALLY
%token <token> FOR
%token <token> FROM
%token <token> GLOBAL
%token <token> IF
%token <token> IMPORT
%token <token> IN
%token <token> IS
%token <token> LAMBDA
%token <token> NONE
%token <token> NOT
%token <token> OR
%token <token> PASS
%token <token> PRINT
%token <token> RAISE
%token <token> RETURN
%token <token> TRY
%token <token> WHILE
%token <token> WITH
%token <token> YIELD

/* compound keywords */
%token <token> NOT_IN
%token <token> IS_NOT

/* symbols */
%token <token> ADD
%token <token> ADD_ASSIGN
%token <token> AND_ASSIGN
%token <token> AND_OP
%token <token> ASSIGN
%token <token> AT
%token <token> BACK_QUOTE
%token <token> COLON
%token <token> COMMA
%token <token> DIV
%token <token> DIV_ASSIGN
%token <token> DOT
%token <token> EQUALS
%token <token> GREATER_THAN
%token <token> GT_EQ
%token <token> IDIV
%token <token> IDIV_ASSIGN
%token <token> LEFT_BRACE
%token <token> LEFT_BRACKET
%token <token> LEFT_PARENTHESIS
%token <token> LEFT_SHIFT
%token <token> LEFT_SHIFT_ASSIGN
%token <token> LESS_THAN
%token <token> LT_EQ
%token <token> MINUS
%token <token> MOD
%token <token> MOD_ASSIGN
%token <token> MULT_ASSIGN
%token <token> NOT_EQUAL
%token <token> NOT_OP
%token <token> OR_ASSIGN
%token <token> OR_OP
%token <token> POWER_ASSIGN
%token <token> RIGHT_BRACE
%token <token> RIGHT_BRACKET
%token <token> RIGHT_PARENTHESIS
%token <token> RIGHT_SHIFT
%token <token> RIGHT_SHIFT_ASSIGN
%token <token> SEMI_COLON
%token <token> STAR
%token <token> STAR_STAR
%token <token> SUB_ASSIGN
%token <token> XOR
%token <token> XOR_ASSIGN



/*
The grammar mostly matches the reference specification:
https://docs.python.org/2/reference/grammar.html

The main difference is that we have to preserve all new lines (not just the
ones that trigger syntactic changes), doc strings, etc.
*/

%type <statements> dotted_as_names
%type <statements> file_input /* code block that's potentially empty */
%type <statements> import_from
%type <statements> import_name
%type <statements> import_stmt
%type <statements> newlines
%type <statements> real_stmts /* code block with at least 1 real statement */
%type <statements> simple_stmt
%type <statements> small_stmt
%type <statements> stmt
%type <statements> stmt_or_newline
%type <statements> stmts /* a code block; may contain only newlines */
%type <statements> suite

%type <statement> assert_stmt
%type <statement> classdef
%type <statement> compound_stmt
%type <statement> decorated
%type <statement> del_stmt
%type <statement> dotted_as_name /* a single import statement */
%type <statement> exec_stmt
%type <statement> expr_stmt
%type <statement> flow_stmt
%type <statement> for_stmt
%type <statement> funcdef
%type <statement> global_stmt
%type <statement> if_stmt
%type <statement> newline  /* an implicit pass statement */
%type <statement> print_stmt
%type <statement> raise_stmt
%type <statement> return_stmt
%type <statement> try_stmt
%type <statement> while_stmt
%type <statement> with_stmt
%type <statement> yield_stmt

%type <token> augassign
%type <token> comp_op
%type <token> factor_sign

%type <tokens> dotted_name
%type <tokens> factor_signs
%type <tokens> namelist
%type <tokens> strings

%type <expression> and_expr
%type <expression> arith_expr
%type <expression> assign_expr /* assign is technically a stmt */
%type <expression> atom
%type <expression> comparison
%type <expression> dict_or_set_maker
%type <expression> expr
%type <expression> factor
%type <expression> fpdef
%type <expression> lambdadef
%type <expression> list_maker
%type <expression> old_lambda
%type <expression> old_test
%type <expression> or_test and_test not_test
%type <expression> power
%type <expression> shift_expr
%type <expression> term
%type <expression> test
%type <expression> testlist_comp
%type <expression> xor_expr
%type <expression> yield_expr
%type <expression> yield_expr_or_testlist

%type <expressions> dictlist
%type <expressions> dictlist_strict
%type <expressions> exprlist
%type <expressions> exprlist_strict
%type <expressions> fplist
%type <expressions> fplist_strict
%type <expressions> old_testlist
%type <expressions> old_testlist_strict
%type <expressions> testlist testlist_strict

%type <argument> argument
%type <argument> fpdef_test

%type <arguments> argumentlist
%type <arguments> argumentlist_strict
%type <arguments> arglist
%type <arguments> vararglist
%type <arguments> fpdef_test_list
%type <arguments> fpdef_test_list_strict

%type <argList> parameters

%type <iterator> list_for
%type <iterator> comp_for

%type <iterators> list_for_list
%type <iterators> comp_for_list


%type <trailer> trailer

%type <trailers> trailers

%type <decorator> decorator

%type <decorators> decorators

%type <intVal> dots

%type <importClause> import_as_name

%type <importClauses> from_clause
%type <importClauses> import_as_names
%type <importClauses> import_as_names_strict

%type <conditionClause> except_clause

%type <conditionClauses> elif_list
%type <conditionClauses> except_list

%type <subscript> sliceop
%type <subscript> subscript
%type <subscripts> subscriptlist
%type <subscripts> subscriptlist_strict

%type <withClause> with_item
%type <withClauses> with_item_list

%%

file_input:
    {
        // empty file is ok
    }
    | stmts {
        // %parser-param is not supported =...(
        yylex.(*Context).Statements = $1
    }
    ;

newline:
    NEWLINE {
        $$ = NewPassStmt($1, true)
    }
    ;

stmt_or_newline:
    stmt {
        $$ = $1
    }
    | newline {
        $$ = []Statement{$1}
    }
    ;

stmts:
    stmt_or_newline {
        $$ = $1
    }
    | stmts stmt_or_newline {
        $$ = append($1, $2...)
    }

newlines:
    newline {
        $$ = []Statement{$1}
    }
    | newlines newline {
        $$ = append($1, $2)
    }
    ;

real_stmts:
    stmt {
        $$ = $1
    }
    | newlines stmt {
        $$ = append($1, $2...)
    }
    | real_stmts stmts {
        $$ = append($1, $2...)
    }
    ;

decorator:
    AT dotted_name NEWLINE {
        $$ = NewDecorator($1, DottedNameToExpr($2), $3)
    }
    | AT dotted_name LEFT_PARENTHESIS RIGHT_PARENTHESIS NEWLINE {
        call := NewCallExpr(DottedNameToExpr($2), NewArgumentList($3, nil, $4))

        $$ = NewDecorator($1, call, $5)
    }
    | AT dotted_name LEFT_PARENTHESIS arglist RIGHT_PARENTHESIS NEWLINE {
        call := NewCallExpr(DottedNameToExpr($2), NewArgumentList($3, $4, $5))

        $$ = NewDecorator($1, call, $6)
    }
    ;

decorators:
    decorator {
        $$ = []*Decorator{$1}
    }
    | decorators decorator {
        $$ = append($1, $2)
    }
    ;

decorated:
    decorators funcdef {
        $2.(*FuncDef).SetDecorators($1)
        $$ = $2
    }
    | decorators classdef {
        $2.(*ClassDef).SetDecorators($1)
        $$ = $2
    }
    ;

funcdef:
    DEF NAME parameters COLON suite {
        $$ = NewFuncDef($1, $2, $3, $4, $5)
    }
    ;

parameters:
    LEFT_PARENTHESIS RIGHT_PARENTHESIS {
        $$ = NewArgumentList($1, nil, $2)
    }
    | LEFT_PARENTHESIS vararglist RIGHT_PARENTHESIS {
        $$ = NewArgumentList($1, $2, $3)
    }
    ;

fpdef_test:
    fpdef {
        $$ = NewArgument($1, nil, nil)
    }
    | fpdef ASSIGN test {
        $$ = NewArgument($1, $2, $3)
    }
    ;

fpdef_test_list_strict:
    fpdef_test {
        $$ = []*Argument{$1}
    }
    | fpdef_test_list_strict COMMA fpdef_test {
        $1[len($1)-1].MergeFrom(&$2.Node)
        $$ = append($1, $3)
    }
    ;

fpdef_test_list:
    fpdef_test_list_strict {
        $$ = $1
    }
    | fpdef_test_list_strict COMMA {
        $$ = $1
        $$[len($$)-1].NodeInfo().MergeFrom(&$2.Node)
    }
    ;

vararglist:
    STAR NAME {
        $$ = []*Argument{NewPositionVarParam($1, NewIdentifier($2))}
    }
    | STAR_STAR NAME {
        $$ = []*Argument{NewKeywordVarParam($1, NewIdentifier($2))}
    }
    | STAR NAME COMMA STAR_STAR NAME {
        arg := NewPositionVarParam($1, NewIdentifier($2))
        arg.MergeFrom(&$3.Node)

        $$ = []*Argument{arg, NewKeywordVarParam($4, NewIdentifier($5))}
    }
    | fpdef_test_list {
        $$ = $1
    }
    | fpdef_test_list_strict COMMA STAR NAME {
        $1[len($1)-1].MergeFrom(&$2.Node)

        $$ = append($1, NewPositionVarParam($3, NewIdentifier($4)))
    }
    | fpdef_test_list_strict COMMA STAR_STAR NAME {
        $1[len($1)-1].MergeFrom(&$2.Node)

        $$ = append($1, NewKeywordVarParam($3, NewIdentifier($4)))
    }
    | fpdef_test_list_strict COMMA STAR NAME COMMA STAR_STAR NAME {
        $1[len($1)-1].MergeFrom(&$2.Node)

        arg := NewPositionVarParam($3, NewIdentifier($4))
        arg.MergeFrom(&$5.Node)

        $$ = append($1, arg)
        $$ = append($1, NewKeywordVarParam($6, NewIdentifier($7)))
    }
    ;

fpdef:
    NAME {
        $$ = NewIdentifier($1)
    }
    | LEFT_PARENTHESIS fplist RIGHT_PARENTHESIS {
        $$ = $2.ConvertToExpr()

        $1.MergeFrom($$.NodeInfo())
        $1.MergeFrom(&$3.Node)

        *$$.NodeInfo() = $1.Node
    }
    ;

fplist_strict:
    fpdef {
        $$ = NewExprList([]Expression{$1})
    }
    | fplist_strict COMMA fpdef {
        $$ = $1
        $$.Expressions[len($$.Expressions)-1].NodeInfo().MergeFrom(&$2.Node)
        $$.Expressions = append($$.Expressions, $3)
    }
    ;

fplist:
    fplist_strict {
        $$ = $1
    }
    | fplist_strict COMMA {
        $$ = $1
        $$.Expressions[len($$.Expressions)-1].NodeInfo().MergeFrom(&$2.Node)
        $$.ExplicitCollection = true
    }
    ;

stmt:
    simple_stmt {
        $$ = $1
    }
    | compound_stmt {
        $$ = []Statement{$1}
    }
    ;

simple_stmt:
    small_stmt NEWLINE {
        $$ = $1
        $$[len($$)-1].NodeInfo().MergeTrailingFrom(&$2.Node)
    }
    | small_stmt SEMI_COLON NEWLINE {
        $$ = $1
        $$[len($$)-1].NodeInfo().MergeTrailingFrom(&$2.Node)
        $$[len($$)-1].NodeInfo().MergeTrailingFrom(&$3.Node)
    }
    | small_stmt SEMI_COLON simple_stmt {
        $1[len($1)-1].NodeInfo().MergeTrailingFrom(&$2.Node)
        $$ = append($1, $3...)
    }

small_stmt:
    expr_stmt {
        $$ = []Statement{$1}
    }
    | print_stmt {
        $$ = []Statement{$1}
    }
    | del_stmt {
        $$ = []Statement{$1}
    }
    | PASS {
        $$ = []Statement{NewPassStmt($1, false)}
    }
    | flow_stmt {
        $$ = []Statement{$1}
    }
    | import_stmt {
        $$ = $1
    }
    | global_stmt {
        $$ = []Statement{$1}
    }
    | exec_stmt {
        $$ = []Statement{$1}
    }
    | assert_stmt {
        $$ = []Statement{$1}
    }
    ;


assign_expr:
    testlist {
        $$ = $1.ConvertToExpr()
    }
    | assign_expr ASSIGN yield_expr_or_testlist {
        // Assign is technically an statement
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

yield_expr_or_testlist:
    yield_expr {
        $$ = $1
    }
    | testlist {
        $$ = $1.ConvertToExpr()
    }
    ;

expr_stmt:
    assign_expr {
        $$ = &ExprStmt{Expression: $1}
    }
    | testlist augassign yield_expr_or_testlist {
        // Augassign is technically an statement
        $$ = &ExprStmt{Expression: NewBinaryExpr($1.ConvertToExpr(), $2, $3)}
    }
    ;

augassign:
    ADD_ASSIGN {
        $$ = $1
    }
    | SUB_ASSIGN {
        $$ = $1
    }
    | MULT_ASSIGN {
        $$ = $1
    }
    | DIV_ASSIGN {
        $$ = $1
    }
    | MOD_ASSIGN {
        $$ = $1
    }
    | AND_ASSIGN {
        $$ = $1
    }
    | OR_ASSIGN {
        $$ = $1
    }
    | XOR_ASSIGN {
        $$ = $1
    }
    | LEFT_SHIFT_ASSIGN {
        $$ = $1
    }
    | RIGHT_SHIFT_ASSIGN {
        $$ = $1
    }
    | POWER_ASSIGN {
        $$ = $1
    }
    | IDIV_ASSIGN {
        $$ = $1
    }
    ;

print_stmt:
    PRINT {
        $$ = NewPrintStmt($1, nil, nil, nil)
    }
    | PRINT testlist {
        $$ = NewPrintStmt($1, nil, nil, $2)
    }
    | PRINT RIGHT_SHIFT test {
        $$ = NewPrintStmt($1, $2, $3, nil)
    }
    | PRINT RIGHT_SHIFT test COMMA testlist {
        $$ = NewPrintStmt($1, $2, $3, $5)
        $$.NodeInfo().MergeFrom(&$4.Node)
    }
    ;

del_stmt:
    DEL exprlist {
        $$ = NewDelStmt($1, $2)
    }
    ;

flow_stmt:
    BREAK {
        $$ = NewBreakStmt($1)
    }
    | CONTINUE {
        $$ = NewContinueStmt($1)
    }
    | return_stmt {
        $$ = $1
    }
    | raise_stmt {
        $$ = $1
    }
    | yield_stmt {
        $$ = $1
    }
    ;

return_stmt:
    RETURN {
        $$ = NewReturnStmt($1, nil)
    }
    | RETURN testlist {
        $$ = NewReturnStmt($1, $2.ConvertToExpr())
    }
    ;

yield_stmt:
    yield_expr {
        $$ = NewExprStmt($1)
    }
    ;

raise_stmt:
    RAISE {
        $$ = NewRaiseStmt($1, nil, nil, nil, nil, nil)
    }
    | RAISE test {
        $$ = NewRaiseStmt($1, $2, nil, nil, nil, nil)
    }
    | RAISE test COMMA test {
        $$ = NewRaiseStmt($1, $2, $3, $4, nil, nil)
    }
    | RAISE test COMMA test COMMA test {
        $$ = NewRaiseStmt($1, $2, $3, $4, $5, $6)
    }
    ;

import_stmt:
    import_name {
        $$ = $1
    }
    | import_from {
        $$ = $1
    }
    ;

import_name:
    IMPORT dotted_as_names {
        $2[0].NodeInfo().MergeLeadingFrom(&$1.Node)
        $$ = $2







//TODO PRESERVE COMMENT FOR BELOW RULES ...

    }
    ;

from_clause:
    STAR {
        $$ = []*ImportClause{}  // Empty list implies all
    }
    | LEFT_PARENTHESIS import_as_names RIGHT_PARENTHESIS {
        $$ = $2
    }
    | import_as_names {
        $$ = $1
    }
    ;

dots:
    DOT {
        $$ = 1
    }
    | dots DOT {
        $$ = $1 + 1
    }
    ;

import_from:
    FROM dotted_name IMPORT from_clause {
        $$ = []Statement{
            &FromStmt{
                DotPrefixCount: 0,
                ModulePath: $2,
                Imports: $4,
            },
        }
    }
    | FROM dots IMPORT from_clause {
        $$ = []Statement{
            &FromStmt{
                DotPrefixCount: $2,
                Imports: $4,
            },
        }
    }
    | FROM dots dotted_name IMPORT from_clause {
        $$ = []Statement{
            &FromStmt{
                DotPrefixCount: $2,
                ModulePath: $3,
                Imports: $5,
            },
        }
    }
    ;

import_as_name:
    NAME {
        $$ = &ImportClause{
            Name: $1,
        }
    }
    | NAME AS NAME {
        $$ = &ImportClause{
            Name: $1,
            Alias: $2,
        }
    }
    ;

dotted_as_name:
    dotted_name {
        $$ = &ImportStmt{ModulePath: $1}
    }
    | dotted_name AS NAME {
        $$ = &ImportStmt{ModulePath: $1, Alias: $2}
    }
    ;

import_as_names_strict:
    import_as_name {
        $$ = []*ImportClause{$1}
    }
    | import_as_names_strict COMMA import_as_name {
        $$ = append($1, $3)
    }
    ;

import_as_names:
    import_as_names_strict {
        $$ = $1
    }
    | import_as_names_strict COMMA {
        $$ = $1
    }
    ;

dotted_as_names:
    dotted_as_name {
        $$ = []Statement{$1}
    }
    | dotted_as_names COMMA dotted_as_name {
        $$ = append($1, $3)
    }
    ;

dotted_name:
    NAME {
        $$ = []*Token{$1}
    }
    | dotted_name DOT NAME {
        $$ = append($1, $3)
    }
    ;

namelist:
    NAME {
        $$ = []*Token{$1}
    }
    | namelist COMMA NAME {
        $$ = append($1, $3)
    }
    ;
global_stmt:
    GLOBAL namelist {
        $$ = &GlobalStmt{Names: $2}
    }
    ;

exec_stmt:
    EXEC expr {
        $$ = &ExecStmt{Expr: $2}
    }
    | EXEC expr IN test {
        $$ = &ExecStmt{Expr: $2, Global: $4}
    }
    | EXEC expr IN test COMMA test {
        $$ = &ExecStmt{Expr: $2, Global: $4, Local: $6}
    }
    ;

assert_stmt:
    ASSERT test {
        $$ = &AssertStmt{
            Expr: $2,
        }
    }
    | ASSERT test COMMA test {
        $$ = &AssertStmt{
            Expr: $2,
            Debug: $4,
        }
    }
    ;

compound_stmt:
    if_stmt {
        $$ = $1
    }
    | while_stmt {
        $$ = $1
    }
    | for_stmt {
        $$ = $1
    }
    | try_stmt {
        $$ = $1
    }
    | with_stmt {
        $$ = $1
    }
    | funcdef {
        $$ = $1
    }
    | classdef {
        $$ = $1
    }
    | decorated {
        $$ = $1
    }
    ;

elif_list:
    ELIF test COLON suite {
        $$ = []*ConditionClause{
            &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
        }
    }
    | elif_list ELIF test COLON suite {
        $$ = append(
            $1,
            &ConditionClause{
                Matches: $3,
                Branch: $5,
            })
    }
    ;

if_stmt:
    IF test COLON suite {
        $$ = &ConditionStmt{
            If: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
        }
    }
    | IF test COLON suite elif_list {
        $$ = &ConditionStmt{
            If: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
            Elif: $5,
        }
    }
    | IF test COLON suite elif_list ELSE COLON suite {
        $$ = &ConditionStmt{
            If: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
            Elif: $5,
            Else: $8,
        }
    }
    |IF test COLON suite ELSE COLON suite {
        $$ = &ConditionStmt{
            If: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
            Else: $7,
        }
    }
    ;

while_stmt:
    WHILE test COLON suite {
        $$ = &WhileStmt{
            Loop: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
        }
    }
    | WHILE test COLON suite ELSE COLON suite {
        $$ = &WhileStmt{
            Loop: &ConditionClause{
                Matches: $2,
                Branch: $4,
            },
            Else: $7,
        }
    }
    ;

for_stmt:
    FOR exprlist IN testlist COLON suite {
        $$ = &ForStmt{
            Iterator: &Iterator{
                BoundVariables: $2.ConvertToExpr(),
                Source: $4.ConvertToExpr(),
            },
            Loop: $6,
        }
    }
    | FOR exprlist IN testlist COLON suite ELSE COLON suite {
        $$ = &ForStmt{
            Iterator: &Iterator{
                BoundVariables: $2.ConvertToExpr(),
                Source: $4.ConvertToExpr(),
            },
            Loop: $6,
            Else: $9,
        }
    }
    ;

except_list:
    except_clause COLON suite {
        $1.Branch = $3
        $$ = []*ConditionClause{$1}
    }
    | except_list except_clause COLON suite {
        $2.Branch = $4
        $$ = append($1, $2)
    }

try_stmt:
    TRY COLON suite except_list {
        $$ = &TryStmt{
            Try: $3,
            Except: $4,
        }
    }
    | TRY COLON suite except_list ELSE COLON suite {
        $$ = &TryStmt{
            Try: $3,
            Except: $4,
            Else: $7,
        }
    }
    | TRY COLON suite except_list FINALLY COLON suite {
        $$ = &TryStmt{
            Try: $3,
            Except: $4,
            Finally: $7,
        }
    }
    | TRY COLON suite except_list ELSE COLON suite FINALLY COLON suite {
        $$ = &TryStmt{
            Try: $3,
            Except: $4,
            Else: $7,
            Finally: $10,
        }
    }
    | TRY COLON suite FINALLY COLON suite {
        $$ = &TryStmt{
            Try: $3,
            Finally: $6,
        }
    }
    ;

with_stmt:
    WITH with_item_list COLON suite {
        $$ = &WithStmt{
            WithClauses: $2,
            Statements: $4,
        }
    }
    ;

with_item_list:
    with_item {
        $$ = []*WithClause{$1}
    }
    | with_item_list COMMA with_item {
        $$ = append($1, $3)
    }
    ;

with_item:
    test {
        $$ = &WithClause{
            Value: $1,
        }
    }
    | test AS expr {
        $$ = &WithClause{
            Value: $1,
            BoundVariable: $3,
        }
    }
    ;

except_clause:
    EXCEPT {
        $$ = &ConditionClause{}
    }
    | EXCEPT test {
        $$ = &ConditionClause{
            Matches: $2,
        }
    }
    | EXCEPT test AS test {
        $$ = &ConditionClause{
            Matches: $2,
            Alias: $4,
        }
    }
    | EXCEPT test COMMA test {
        $$ = &ConditionClause{
            Matches: $2,
            Alias: $4,
        }
    }
    ;

suite:
    simple_stmt {
        $$ = $1
    }
    | NEWLINE INDENT real_stmts DEDENT {
        $$ = $3
    }
    ;

old_test:
    or_test {
        $$ = $1
    }
    | old_lambda {
        $$ = $1
    }
    ;

old_lambda:
    LAMBDA COLON old_test {
        $$ = &LambdaExpr{
            Value: $3,
        }
    }
    | LAMBDA vararglist COLON old_test {
        $$ = &LambdaExpr{
            Arguments: $2,
            Value: $4,
        }
    }
    ;

test:
    or_test {
        $$ = $1
    }
    | or_test IF or_test ELSE test {
        $$ = &ConditionExpr{
            True: $1,
            Predicate: $3,
            False: $5,
        }
    }
    | lambdadef {
        $$ = $1
    }
    ;

or_test:
    and_test {
        $$ = $1
    }
    | or_test OR and_test {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

and_test:
    not_test {
        $$ = $1
    }
    | and_test AND not_test {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

not_test:
    comparison {
        $$ = $1
    }
    | NOT not_test {
        $$ = &UnaryExpr{
            Op: $1,
            Value: $2,
        }
    }
    ;

comparison:
    expr {
        $$ = $1
    }
    | comparison comp_op expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

comp_op:
    LESS_THAN {
        $$ = $1
    }
    | GREATER_THAN {
        $$ = $1
    }
    | EQUALS {
        $$ = $1
    }
    | GT_EQ {
        $$ = $1
    }
    | LT_EQ {
        $$ = $1
    }
    | NOT_EQUAL {
        $$ = $1
    }
    | IN {
        $$ = $1
    }
    | IS {
        $$ = $1
    }
    | NOT IN {
        $$= $1
        $$.TokenType = NOT_IN
        // TODO pull IN's comment into NOT_IN
    }
    | IS NOT {
        $$= $1
        $$.TokenType = IS_NOT
        // TODO pull NOT's comment into NOT_IN
    }
    ;

expr:
    xor_expr {
        $$ = $1
    }
    | expr OR_OP xor_expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

xor_expr:
    and_expr {
        $$ = $1
    }
    | xor_expr XOR and_expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

and_expr:
    shift_expr {
        $$ = $1
    }
    | and_expr AND_OP shift_expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

shift_expr:
    arith_expr {
        $$ = $1
    }
    | shift_expr LEFT_SHIFT arith_expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | shift_expr RIGHT_SHIFT arith_expr {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

arith_expr:
    term {
        $$ = $1
    }
    | arith_expr ADD term {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | arith_expr MINUS term {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

term:
    factor {
        $$ = $1
    }
    | term STAR factor {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | term DIV factor {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | term MOD factor {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | term IDIV factor {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    ;

factor_sign:
    ADD {
        $$ = $1
    }
    | MINUS {
        $$ = $1
    }
    | NOT_OP {
        $$ = $1
    }
    ;

factor_signs:
    factor_sign {
        $$ = []*Token{$1}
    }
    | factor_signs factor_sign {
        $$ = append($1, $2)
    }
    ;

factor:
    power {
        $$ = $1
    }
    | factor_signs power {
        $$ = NewFactorExpr($1, $2)
    }
    ;

power:
    atom {
        $$ = $1
    }
    | atom STAR_STAR factor {
        $$ = NewBinaryExpr($1, $2, $3)
    }
    | atom trailers {
        $$ = NewExpressionFromTrailers($1, $2)
    }
    | atom trailers STAR_STAR factor {
        $$ = NewBinaryExpr(NewExpressionFromTrailers($1, $2), $3, $4)
    }
    ;

trailers:
    trailer {
        $$ = []interface{}{$1}
    }
    | trailers trailer {
        $$ = append($1, $2)
    }
    ;

atom:
    LEFT_PARENTHESIS RIGHT_PARENTHESIS {
        $$ = &CollectionExpr{
            Type: TupleCollection,
        }
    }
    | LEFT_PARENTHESIS yield_expr RIGHT_PARENTHESIS {
        $$ = $2
    }
    | LEFT_PARENTHESIS testlist_comp RIGHT_PARENTHESIS {
        $$ = $2
    }
    | LEFT_BRACKET RIGHT_BRACKET {
        $$ = &CollectionExpr{
            Type: ListCollection,
        }
    }
    | LEFT_BRACKET list_maker RIGHT_BRACKET {
        $$ = $2
    }
    | LEFT_BRACE RIGHT_BRACE {
        $$ = &CollectionExpr{
            Type: DictCollection,
        }
    }
    | LEFT_BRACE dict_or_set_maker RIGHT_BRACE {
        $$ = $2
    }
    | BACK_QUOTE testlist_strict BACK_QUOTE {
        $$ = &EvalExpr{
            Expression: $2.ConvertToExpr(),
        }
    }
    | NONE {
        $$ = NewNone($1)
    }
    | NAME {
        // NOTE: True / False are classified as NAME because they are
        // reassignable variables, not constants ...
        //
        // >>> True = False
        // >>> True
        // False
        $$ = NewIdentifier($1)
    }
    | INTEGER {
        $$ = NewNumber($1)
    }
    | FLOAT {
        $$ = NewNumber($1)
    }
    | strings {
        $$ = &String{
            Pieces: $1,
        }
    }
    ;

strings:
    STRING {
        $$ = []*Token{$1}
    }
    | strings STRING {
        $$ = append($1, $2)
    }
    ;

list_maker:
    testlist {
        $1.ExplicitCollection = true
        $$ = $1.ConvertToExpr()
        $$.(*CollectionExpr).Type = ListCollection
    }
    | test list_for_list {
        $$ = &ComprehensionExpr{
            Type: ListComprehension,
            Value: $1,
            Iterators: $2,
        }
    }
    ;

testlist_comp:
    testlist {
        $$ = $1.ConvertToExpr()
    }
    | test comp_for_list {
        $$ = &ComprehensionExpr{
            Type: GeneratorComprehension,
            Value: $1,
            Iterators: $2,
        }
    }
    ;

lambdadef:
    LAMBDA COLON test {
        $$ = &LambdaExpr{
            Value: $3,
        }
    }
    | LAMBDA vararglist COLON test {
        $$ = &LambdaExpr{
            Arguments: $2,
            Value: $4,
        }
    }
    ;

trailer:
    LEFT_PARENTHESIS RIGHT_PARENTHESIS {
        $$ = []*Argument{}
    }
    | LEFT_PARENTHESIS arglist RIGHT_PARENTHESIS {
        $$ = $2
    }
    | LEFT_BRACKET subscriptlist RIGHT_BRACKET {
        $$ = $2
    }
    | DOT NAME {
        $$ = $2
    }
    ;

subscriptlist_strict:
    subscript {
        $$ = []*Subscript{$1}
    }
    | subscriptlist_strict COMMA subscript {
        $$ = append($1, $3)
    }
    ;

subscriptlist:
    subscriptlist_strict {
        $$ = $1
    }
    | subscriptlist_strict COMMA {
        $$ = $1
    }
    ;

subscript:
    DOT DOT DOT {
        $$ = &Subscript{Ellipsis: true}
    }
    | test {
        $$ = &Subscript{Index: $1}
    }
    | COLON {
        $$ = &Subscript{}
    }
    | test COLON {
        $$ = &Subscript{Left: $1}
    }
    | COLON test {
        $$ = &Subscript{Middle: $2}
    }
    | test COLON test {
        $$ = &Subscript{Left: $1, Middle: $3}
    }
    | COLON sliceop {
        $$ = $2
    }
    | test COLON sliceop {
        $$ = $3
        $$.Left = $1
    }
    | COLON test sliceop {
        $$ = $3
        $$.Middle = $2
    }
    | test COLON test sliceop {
        $$ = $4
        $$.Left = $1
        $$.Middle = $3
    }
    ;

sliceop:
    COLON {
        $$ = &Subscript{}
    }
    | COLON test {
        $$ = &Subscript{Right: $2}
    }
    ;

old_testlist_strict:
    old_test {
        $$ = &ExprList{
            Expressions: []Expression{$1},
        }
    }
    | old_testlist_strict COMMA old_test {
        $$ = $1
        $$.Expressions = append($$.Expressions, $3)
        $$.ExplicitCollection = true
    }
    ;

old_testlist:
    old_testlist_strict {
        $$ = $1
    }
    | old_testlist_strict COMMA {
        $$ = $1
        $$.ExplicitCollection = true
    }
    ;
exprlist_strict:
    expr {
        $$ = &ExprList{
            Expressions: []Expression{$1},
        }
    }
    | exprlist_strict COMMA expr {
        $$ = $1
        $$.Expressions = append($$.Expressions, $3)
        $$.ExplicitCollection = true
    }
    ;

exprlist:
    exprlist_strict {
        $$ = $1
    }
    | exprlist_strict COMMA {
        $$ = $1
        $$.ExplicitCollection = true
    }
    ;


testlist_strict:
    test {
        $$ = &ExprList{
            Expressions: []Expression{$1},
        }
    }
    | testlist_strict COMMA test {
        $$ = $1
        $$.Expressions = append($$.Expressions, $3)
    }
    ;

testlist:
    testlist_strict {
        $$ = $1
    }
    | testlist_strict COMMA {
        $$ = $1
        $$.ExplicitCollection = true
    }
    ;

dictlist_strict:
    test COLON test {
        $$ = &ExprList{
            Expressions: []Expression{
                &CollectionExpr{
                    Type: TupleCollection,
                    Items: []Expression{$1, $3},
                },
            },
            ExplicitCollection: true,
        }
    }
    | dictlist_strict COMMA test COLON test {
        $$ = $1
        $$.Expressions = append(
            $$.Expressions,
            &CollectionExpr{
                Type: TupleCollection,
                Items: []Expression{$3, $5},
            })
    }

dictlist:
    dictlist_strict {
        $$ = $1
    }
    | dictlist_strict COMMA {
        $$ = $1
    }
    ;

dict_or_set_maker:
    dictlist {
        $$ = $1.ConvertToExpr()
        $$.(*CollectionExpr).Type = DictCollection
    }
    | test COLON test comp_for_list {
        $$ = &ComprehensionExpr{
            Type: DictComprehension,
            Key: $1,
            Value: $3,
            Iterators: $4,
        }
    }
    | testlist {
        $1.ExplicitCollection = true
        $$ = $1.ConvertToExpr()
        $$.(*CollectionExpr).Type = SetCollection

    }
    | test comp_for_list {
        $$ = &ComprehensionExpr{
            Type: SetComprehension,
            Key: $1,
            Iterators: $2,
        }
    }
    ;

classdef:
    CLASS NAME COLON suite {
        $$ = &ClassDef {
            Name: $2,
            Statements: $4,
        }
    }
    | CLASS NAME LEFT_PARENTHESIS RIGHT_PARENTHESIS COLON suite {
        $$ = &ClassDef {
            Name: $2,
            Statements: $6,
        }
    }
    | CLASS NAME LEFT_PARENTHESIS testlist RIGHT_PARENTHESIS COLON suite {
        $$ = &ClassDef {
            Name: $2,
            ParentClasses: $4.Expressions,
            Statements: $7,
        }
    }
    ;

argumentlist_strict:
    argument {
        $$ = []*Argument{$1}
    }
    | argumentlist_strict COMMA argument {
        $$ = append($1, $3)
    }
    ;

argumentlist:
    argumentlist_strict {
        $$ = $1
    }
    | argumentlist_strict COMMA {
        $$ = $1
    }
    ;

arglist:
    argumentlist {
        $$ = $1
    }
    | argumentlist_strict COMMA STAR test {
        $1[len($1)-1].MergeFrom(&$2.Node)

        $$ = append($1, NewPositionVarArg($3, $4))
    }
    | argumentlist_strict COMMA STAR test COMMA argumentlist_strict {
        $1[len($1)-1].MergeFrom(&$2.Node)

        arg := NewPositionVarArg($3, $4)
        arg.MergeFrom(&$5.Node)

        $$ = append($1, arg)
        $$ = append($$, $6...)
    }
    | argumentlist_strict COMMA STAR test COMMA STAR_STAR test {
        $1[len($1)-1].MergeFrom(&$2.Node)

        arg := NewPositionVarArg($3, $4)
        arg.MergeFrom(&$5.Node)

        $$ = append($1, arg)
        $$ = append($$, NewKeywordVarArg($6, $7))
    }
    | argumentlist_strict COMMA STAR test COMMA argumentlist_strict COMMA STAR_STAR test {
        $1[len($1)-1].MergeFrom(&$2.Node)

        arg := NewPositionVarArg($3, $4)
        arg.MergeFrom(&$5.Node)

        $$ = append($1, arg)
        $$ = append($$, $6...)

        $$[len($$)-1].MergeFrom(&$7.Node)

        $$ = append($$, NewKeywordVarArg($8, $9))
    }
    | argumentlist_strict COMMA STAR_STAR test {
        $1[len($1)-1].MergeFrom(&$2.Node)
        $$ = append($1, NewKeywordVarArg($3, $4))
    }
    | STAR test {
        $$ = []*Argument{NewPositionVarArg($1, $2)}
    }
    | STAR test COMMA argumentlist_strict {
        arg := NewPositionVarArg($1, $2)
        arg.MergeFrom(&$3.Node)

        $$ = append([]*Argument{arg}, $4...)
    }
    | STAR test COMMA STAR_STAR test {
        arg := NewPositionVarArg($1, $2)
        arg.MergeFrom(&$3.Node)

        $$ = []*Argument{arg, NewKeywordVarArg($4, $5)}
    }
    | STAR test COMMA argumentlist_strict COMMA STAR_STAR test {
        arg := NewPositionVarArg($1, $2)
        arg.MergeFrom(&$3.Node)

        $$ = append([]*Argument{arg}, $4...)
        $$[len($$)-1].MergeFrom(&$5.Node)

        $$ = append($$, NewKeywordVarArg($6, $7))
    }
    | STAR_STAR test {
        $$ = []*Argument{NewKeywordVarArg($1, $2)}
    }
    ;


argument:
    test {
        $$ = NewArgument(nil, nil, $1)
    }
    | test comp_for_list {
        $$ = NewArgument(
            nil,
            nil,
            &ComprehensionExpr{
                Type: GeneratorComprehension,
                Value: $1,
                Iterators: $2,
            })
    }
    | test ASSIGN test {
        $$ = NewArgument($1, $2, $3)
    }
    ;

list_for:
    FOR exprlist IN old_testlist {
        $$ = &Iterator {
            BoundVariables: $2.ConvertToExpr(),
            Source: $4.ConvertToExpr(),
        }
    }
    | FOR exprlist IN old_testlist IF old_test {
        $$ = &Iterator {
            BoundVariables: $2.ConvertToExpr(),
            Source: $4.ConvertToExpr(),
            Filters: []Expression{$6},
        }
    }
    ;

list_for_list:
    list_for {
        $$ = []*Iterator{$1}
    }
    | list_for_list list_for {
        $$ = append($1, $2)
    }
    | list_for_list IF old_test {
        $$ = $1
        last := $$[len($$)-1]
        last.Filters = append(last.Filters, $3)
    }
    ;

comp_for:
    FOR exprlist IN or_test {
        $$ = &Iterator {
            BoundVariables: $2.ConvertToExpr(),
            Source: $4,
        }
    }
    | FOR exprlist IN or_test IF old_test {
        $$ = &Iterator {
            BoundVariables: $2.ConvertToExpr(),
            Source: $4,
            Filters: []Expression{$6},
        }
    }
    ;

comp_for_list:
    comp_for {
        $$ = []*Iterator{$1}
    }
    | comp_for_list comp_for {
        $$ = append($1, $2)
    }
    | comp_for_list IF old_test {
        $$ = $1
        last := $$[len($$)-1]
        last.Filters = append(last.Filters, $3)
    }
    ;

yield_expr:
    YIELD {
        $$ = &YieldExpr{}
    }
    | YIELD testlist {
        $$ = &YieldExpr{Expression: $2.ConvertToExpr()}
    }
    ;

%%
