#ifndef BUILDUTIL_TEST_TMPL_INL_H_
#define BUILDUTIL_TEST_TMPL_INL_H_

template<typename T>
T max(T a, T b) {
  if (a > b) {
    return a;
  }

  return b;
}

#endif  // BUILDTUIL_TEST_TMPL_INL_H_

