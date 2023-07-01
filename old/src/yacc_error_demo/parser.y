%{
package main

%}

%union {
    Node
    NodeList []Node
}

%token <Node> ID PLUS MINUS LBRACE RBRACE

%token <Node> BLOCK ERROR EOF // pseudo token

%type <Node> block stmt
%type <NodeList> stmt_list


%%

start:
    block {
        demolex.(*parseContext).result = $1
    }
    ;

block:
    LBRACE stmt_list RBRACE {
        $$ = &Block{$1, $2, $3}
    }
    ;

stmt_list:
    stmt_list stmt {
        $$ = append($1, $2)
    }
    | { // empty
        $$ = nil
    }
    ;

stmt:
    ID {
        $$ = $1
    }
    | ID PLUS ID {
        $$ = &Binary{$1, $2, $3}
    }
    | ID MINUS ID {
        $$ = &Binary{$1, $2, $3}
    }
    | BLOCK {
        $$ = $1
    }
    | ERROR {
        $$ = $1
    }
    ;

%%

func init() {
    demoErrorVerbose = true
}
