package internal

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type Assertion struct {
	t    TestingT
	Fail func(t TestingT, failureMessage string)
}

func NewAssert(t TestingT) *Assertion {
	return &Assertion{
		t: t,
		Fail: func(t TestingT, failureMessage string) {
			t.Errorf(failureMessage)
		},
	}
}

func (r *Assertion) True(value bool, ) {
	if !value {
		r.Fail(r.t, "Should be true")
	}
}

func (r *Assertion) False(value bool, ) {
	if value {
		r.Fail(r.t, "Should be false")
	}
}
