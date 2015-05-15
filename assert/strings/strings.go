package strings

import (
	"testing"
)

func AssertThat(reason string, actual string, matcher StringMatcher, t *testing.T) {
	if !matcher.Match(actual) {
		t.Errorf("%s\nActual: %s\nExpected: %s", reason, actual, matcher.expected)
	}
}

func Is(expected string) StringMatcher {
	return StringMatcher{expected}
}

type StringMatcher struct {
	expected string
}

func (m StringMatcher) Match(actual string) bool {
	return m.expected == actual
}