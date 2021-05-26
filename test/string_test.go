package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/chyroc/god/god_string"
)

func Test_Prefix(t *testing.T) {
	as := NewAssert(t)

	as.True(god_string.HasPrefixList("abc1234", "a", "b", "c"))
	as.False(god_string.HasPrefixList("abc1234", "x", "b", "c"))
}

func Test_CountPre(t *testing.T) {
	as := NewAssert(t)

	as.Equal(int(0), god_string.CountPre("", ""))
	as.Equal(int(0), god_string.CountPre("1", ""))
	as.Equal(int(0), god_string.CountPre("12", ""))

	as.Equal(int(0), god_string.CountPre("123", "2"))
	as.Equal(int(0), god_string.CountPre("123", "13"))
	as.Equal(int(0), god_string.CountPre("123", "23"))

	as.Equal(int(0), god_string.CountPre("123", ""))
	as.Equal(int(0), god_string.CountPre("123", ""))
	as.Equal(int(0), god_string.CountPre("123", ""))

	as.Equal(int(0), god_string.CountPre("", "2"))
	as.Equal(int(0), god_string.CountPre("", "12"))
	as.Equal(int(0), god_string.CountPre("", "23"))

	as.Equal(int(1), god_string.CountPre(".", "."))
	as.Equal(int(2), god_string.CountPre("..", "."))

	as.Equal(int(0), god_string.CountPre(".", ".."))
	as.Equal(int(1), god_string.CountPre("..", ".."))
}

func Test_FindLastSubstr(t *testing.T) {
	as := NewAssert(t)

	as.Equal("", god_string.FindLastSubstr("", ""))
	as.Equal("", god_string.FindLastSubstr("1", ""))
	as.Equal("", god_string.FindLastSubstr("12", ""))
	as.Equal("", god_string.FindLastSubstr("123", ""))

	as.Equal("", god_string.FindLastSubstr("", ""))
	as.Equal("", god_string.FindLastSubstr("", "1"))
	as.Equal("", god_string.FindLastSubstr("", "12"))
	as.Equal("", god_string.FindLastSubstr("", "123"))

	as.Equal("", god_string.FindLastSubstr("", "/"))
	as.Equal("1", god_string.FindLastSubstr("1", "/"))
	as.Equal("12", god_string.FindLastSubstr("12", "/"))
	as.Equal("123", god_string.FindLastSubstr("123", "/"))

	as.Equal("", god_string.FindLastSubstr("", "/"))
	as.Equal("1", god_string.FindLastSubstr("/1", "/"))
	as.Equal("2", god_string.FindLastSubstr("1/2", "/"))
	as.Equal("3", god_string.FindLastSubstr("12/3", "/"))

	as.Equal("", god_string.FindLastSubstr("", "/"))
	as.Equal("1", god_string.FindLastSubstr("x/x/1", "/"))
	as.Equal("2", god_string.FindLastSubstr("x/x//1/2", "/"))
	as.Equal("3", god_string.FindLastSubstr("x/x/x/x///12/3", "/"))

	as.Equal("", god_string.FindLastSubstr("", "/"))
	as.Equal("", god_string.FindLastSubstr("/1/", "/"))
	as.Equal("", god_string.FindLastSubstr("1/2/", "/"))
	as.Equal("", god_string.FindLastSubstr("12/3/", "/"))

	fmt.Println(strings.Split("", "1"), len(strings.Split("", "1")), strings.Split("", "1")[0])
}

func Test_TrimLeftRightWrap(t *testing.T) {
	as := NewAssert(t)

	as.Equal("", god_string.TrimLeftRightWrap("", "1"))
	as.Equal("", god_string.TrimLeftRightWrap("", "12"))
	as.Equal("", god_string.TrimLeftRightWrap("", ""))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("1", god_string.TrimLeftRightWrap("1", ""))
	as.Equal("12", god_string.TrimLeftRightWrap("12", ""))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("x1", god_string.TrimLeftRightWrap("x1", "x"))
	as.Equal("x12", god_string.TrimLeftRightWrap("x12", "x"))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("1x", god_string.TrimLeftRightWrap("1x", "x"))
	as.Equal("12x", god_string.TrimLeftRightWrap("12x", "x"))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("1", god_string.TrimLeftRightWrap("x1x", "x"))
	as.Equal("12", god_string.TrimLeftRightWrap("x12x", "x"))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("1", god_string.TrimLeftRightWrap("xx1xx", "xx"))
	as.Equal("12", god_string.TrimLeftRightWrap("xx12xx", "xx"))

	as.Equal("", god_string.TrimLeftRightWrap("", ""))
	as.Equal("x1x", god_string.TrimLeftRightWrap("xx1xx", "x"))
	as.Equal("x12x", god_string.TrimLeftRightWrap("xx12xx", "x"))
}
