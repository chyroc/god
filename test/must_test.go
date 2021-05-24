package test

import (
	"fmt"
	"testing"

	"github.com/chyroc/god/god_assert"
)

func Test_MustNil(t *testing.T) {
	as := NewAssert(t)

	god_assert.MustNil(nil)

	defer func() {
		e := recover()
		as.Equal("err", fmt.Sprintf("%s", e))
	}()

	god_assert.MustNil(fmt.Errorf("err"))
}
