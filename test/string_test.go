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

func Test_FindLastSubstr(t *testing.T) {
	as := internal.NewAssert(t)

	as.Equal("", god.Str.FindLastSubstr("", ""))
	as.Equal("", god.Str.FindLastSubstr("1", ""))
	as.Equal("", god.Str.FindLastSubstr("12", ""))
	as.Equal("", god.Str.FindLastSubstr("123", ""))

	as.Equal("", god.Str.FindLastSubstr("", ""))
	as.Equal("", god.Str.FindLastSubstr("", "1"))
	as.Equal("", god.Str.FindLastSubstr("", "12"))
	as.Equal("", god.Str.FindLastSubstr("", "123"))

	as.Equal("", god.Str.FindLastSubstr("", "/"))
	as.Equal("1", god.Str.FindLastSubstr("1", "/"))
	as.Equal("12", god.Str.FindLastSubstr("12", "/"))
	as.Equal("123", god.Str.FindLastSubstr("123", "/"))

	as.Equal("", god.Str.FindLastSubstr("", "/"))
	as.Equal("1", god.Str.FindLastSubstr("/1", "/"))
	as.Equal("2", god.Str.FindLastSubstr("1/2", "/"))
	as.Equal("3", god.Str.FindLastSubstr("12/3", "/"))

	as.Equal("", god.Str.FindLastSubstr("", "/"))
	as.Equal("1", god.Str.FindLastSubstr("x/x/1", "/"))
	as.Equal("2", god.Str.FindLastSubstr("x/x//1/2", "/"))
	as.Equal("3", god.Str.FindLastSubstr("x/x/x/x///12/3", "/"))

	as.Equal("", god.Str.FindLastSubstr("", "/"))
	as.Equal("", god.Str.FindLastSubstr("/1/", "/"))
	as.Equal("", god.Str.FindLastSubstr("1/2/", "/"))
	as.Equal("", god.Str.FindLastSubstr("12/3/", "/"))
}
