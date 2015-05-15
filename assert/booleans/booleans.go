package booleans

import (
	"testing"
)

func AssertThat(reason string, actual bool, matcher BooleanMatcher, t *testing.T) {
	if !matcher.Match(actual) {
		t.Errorf("%s\nActual: %t\nExpected: %t", reason, actual, matcher.expected)
	}
}

func Is(expected bool) BooleanMatcher {
	return BooleanMatcher{expected}
}

type BooleanMatcher struct {
	expected bool
}

func (m BooleanMatcher) Match(actual bool) bool {
	return m.expected == actual
}