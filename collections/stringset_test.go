package collection

import (
  "testing"
  "github.com/alfredqy/golib/assert/booleans"
  "github.com/alfredqy/golib/assert/integers"
)

func TestAddString(t *testing.T) {
  set := NewStringSet()
  integers.AssertThat("the size of the set is zero initially", set.Size(), integers.Is(0), t)
  exist := set.Add("hello")
  booleans.AssertThat("there is no existing \"hello\"", exist, booleans.Is(true), t)
  integers.AssertThat("there is one element in the set", set.Size(), integers.Is(1), t)
  booleans.AssertThat("set contains \"hello\"", set.Contains("hello"), booleans.Is(true), t)
  exist = set.Add("hello")
  booleans.AssertThat("there is an existing \"hello\"", exist, booleans.Is(false), t)
  integers.AssertThat("there is still one element in the set", set.Size(), integers.Is(1), t)
}