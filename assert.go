package assert

import (
	"regexp"
	"testing"
)

type Assert struct {
	t *testing.T
}

func NewAssert(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) Equal(actual interface{}, expected interface{}, msg string) {
	if actual != expected {
		a.t.Errorf("%s, actual: '%v', expected: '%v'", msg, actual, expected)
	}
}

func (a *Assert) Match(pattern string, target string, msg string) {
	match, err := regexp.MatchString(pattern, target)
	if !match || err != nil {
		a.t.Errorf("%s, '%v' does not match '%v'", msg, target, pattern)
	}
}
