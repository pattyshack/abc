#include "buildutil/test/lib1.h"
#include "buildutil/test/tmpl_inl.h"

int h() {
  return max(f(), g());
}
