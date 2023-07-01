%require "2.5.37"
%skeleton "lalr1.cc"  /* aka c++ template */
%locations
%defines
%define api.namespace { test }
%define parser_class_name { TestParser }
%parse-param { test::TestScanner* scanner }
%lex-param { test::TestScanner* scanner }

%code requires {
#include <string>

namespace test {
  class TestScanner;
}
}

%code {
static int yylex(
    test::TestParser::semantic_type* yylval,
    test::TestParser::location_type* loc,
    test::TestScanner* scanner);
}

%union {
  int int_val;
  std::string* str_val;
}

%token <int_val> INTEGER
%token <str_val> STRING

%%

program
  : program val
  | val
  ;

val
  : STRING {
      std::cout << "STR: " << *$1 << std::endl;
      delete $1;
    }
  | INTEGER {
      std::cout << "INT: " << $1 << std::endl;
    }
  ;

%%

#include "buildutil/test/test_scanner.h"

void test::TestParser::error(
    const test::TestParser::location_type& loc,
    const std::string& msg) {
  std::cerr << "Error: " << msg << std::endl;
}

static int yylex(
    test::TestParser::semantic_type* yylval,
    test::TestParser::location_type* loc,
    test::TestScanner* scanner) {
  return scanner->yylex(yylval);
}

