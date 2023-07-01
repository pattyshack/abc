#ifndef BUILDUTIL_TEST_TEST_SCANNER_H_
#define BUILDUTIL_TEST_TEST_SCANNER_H_

#if !defined(yyFlexLexerOnce)
#include <FlexLexer.h>
#endif

#undef YY_DECL
#define YY_DECL int test::TestScanner::yylex()

#include "buildutil/test/test_parser.tab.hh"


namespace test {

class TestScanner : public yyFlexLexer {
 public:
  int yylex(test::TestParser::semantic_type* lval) {
    yylval = lval;
    return yylex();
  }

 private:
  // Scan function created by flex.  Kept in private since we want to
  // initialize yylval_.
  int yylex();

  test::TestParser::semantic_type* yylval;
};

}  // namespace test

#endif  // BUILDUTIL_TEST_TEST_SCANNER_H_
