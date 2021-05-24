package test

import (
	"fmt"
	"reflect"
	"testing"
)

// TestingT .
type TestingT interface {
	Errorf(format string, args ...interface{})
}

// Assertion .
type Assertion struct {
	t    TestingT
	Fail func(t TestingT, failureMessage string)
}

// NewAssert .
func NewAssert(t TestingT) *Assertion {
	return &Assertion{
		t: t,
		Fail: func(t TestingT, failureMessage string) {
			t.Errorf(failureMessage)
		},
	}
}

// True .
func (r *Assertion) True(value bool) {
	if !value {
		r.Fail(r.t, "Should be true")
	}
}

// Equal .
func (r *Assertion) Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		r.Fail(r.t, fmt.Sprintf("%#v != %#v", a, b))
	}
}

// False .
func (r *Assertion) False(value bool) {
	if value {
		r.Fail(r.t, "Should be false")
	}
}

// Nil .
func (r *Assertion) Nil(err error) {
	if err != nil {
		r.Fail(r.t, fmt.Sprintf("%v should be nil", err))
	}
}

// NotNil .
func (r *Assertion) NotNil(err error) {
	if err == nil {
		r.Fail(r.t, fmt.Sprintf("%v should not be nil", err))
	}
}

func Test_Assert(t *testing.T) {
	as := NewAssert(t)
	as.Fail = func(t TestingT, failureMessage string) {
	}

	as.True(true)
	as.False(true)
	as.True(false)
	as.False(false)
	as.Fail(t, "fail")
	as.Equal(1, 1)
	as.Equal(1, 2)
}
