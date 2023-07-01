#include <stdio.h>

#include "buildutil/test/hello_world.h"
#include "buildutil/test/lib2.h"

int main(int argc, char** argv) {
  printf("hello %s! %d\n", world().c_str(), h());

  return 0;
}
