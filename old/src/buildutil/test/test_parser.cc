#include "buildutil/test/test_parser.tab.hh"

#include "buildutil/test/test_scanner.h"

int main(int argc, char** argv) {
  test::TestScanner scanner;
  test::TestParser parser(&scanner);

  parser.parse();
}
