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

func Test_CountPre(t *testing.T) {
	as := internal.NewAssert(t)

	as.Equal(int(0), god.Str.CountPre("", ""))
	as.Equal(int(0), god.Str.CountPre("1", ""))
	as.Equal(int(0), god.Str.CountPre("12", ""))

	as.Equal(int(0), god.Str.CountPre("123", "2"))
	as.Equal(int(0), god.Str.CountPre("123", "13"))
	as.Equal(int(0), god.Str.CountPre("123", "23"))

	as.Equal(int(0), god.Str.CountPre("123", ""))
	as.Equal(int(0), god.Str.CountPre("123", ""))
	as.Equal(int(0), god.Str.CountPre("123", ""))

	as.Equal(int(0), god.Str.CountPre("", "2"))
	as.Equal(int(0), god.Str.CountPre("", "12"))
	as.Equal(int(0), god.Str.CountPre("", "23"))

	as.Equal(int(1), god.Str.CountPre(".", "."))
	as.Equal(int(2), god.Str.CountPre("..", "."))

	as.Equal(int(0), god.Str.CountPre(".", ".."))
	as.Equal(int(1), god.Str.CountPre("..", ".."))
}
