package internal

import (
	"testing"
)

func Test_Assert(t *testing.T) {
	as := NewAssert(t)
	as.Fail = func(t TestingT, failureMessage string) {
	}

	as.True(true)
	as.False(true)
	as.True(false)
	as.False(false)
	as.Fail(t, "fail")
}
