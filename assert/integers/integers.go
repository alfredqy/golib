package integers

import (
	"testing"
)

func AssertThat(reason string, actual int, matcher IntMatcher, t *testing.T) {
  if !matcher.Match(actual) {
    t.Errorf("%s\nActual: %d\nExpected: %d", reason, actual, matcher.expected)
  }
}

func Is(expected int) IntMatcher {
  return IntMatcher{expected}
}

type IntMatcher struct {
  expected int
}

func (m IntMatcher) Match(actual int) bool {
  return m.expected == actual
}
