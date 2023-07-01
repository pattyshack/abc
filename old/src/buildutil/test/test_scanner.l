%{
#include <stdlib.h>

#include "buildutil/test/test_scanner.h"
%}

%option nodefault yyclass="FlexScanner" noyywrap c++

%%

[0-9][0-9]* {
  yylval->int_val = atoi(yytext);
  return test::TestParser::token::INTEGER;
}
.+|\n       {
  yylval->str_val = new std::string(yytext);
  return test::TestParser::token::STRING;
}
