/* definitions */
%{
#include <stdio.h>
%}

%option noyywrap

DIGIT [0-9]

%%
{DIGIT}+ {
    printf("INT: %s\n", yytext);
  }
%%
/* code */

int main(int argc, char** argv) {
  yyin = stdin;

  yylex();
}
