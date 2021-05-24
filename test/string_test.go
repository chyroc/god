package test

import (
	"testing"

	"github.com/chyroc/god"
	"github.com/chyroc/god/internal"
)

func Test_Prefix(t *testing.T) {
	as := internal.NewAssert(t)

	as.True(god.Str.HasPrefixList("abc1234", "a", "b", "c"))
	as.False(god.Str.HasPrefixList("abc1234", "x", "b", "c"))
}
