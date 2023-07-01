#include <iostream>
#include <string>
#include <vector>

#include "absl/strings/str_join.h"

int main() {
  std::vector<std::string> list = {"hello", "world"};

  std::cout << absl::StrJoin(list, " ") << "\n";

  return 0;
}
