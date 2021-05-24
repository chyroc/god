package internal

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type Assertion struct {
	t TestingT
}

func NewAssert(t TestingT) *Assertion {
	return &Assertion{
		t: t,
	}
}

func (r *Assertion) True(value bool, ) {
	if !value {
		Fail(r.t, "Should be true")
	}
}

func (r *Assertion) False(value bool, ) {
	if value {
		Fail(r.t, "Should be false")
	}
}

func Fail(t TestingT, failureMessage string) {
	t.Errorf(failureMessage)
}
