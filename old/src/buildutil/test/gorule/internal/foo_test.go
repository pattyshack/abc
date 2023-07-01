package internal

import "testing"

func TestFoo(t *testing.T) {
  if Foo() != 42 {
    t.FailNow()
  }
}
