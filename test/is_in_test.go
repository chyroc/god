package test

import (
	"testing"

	"github.com/chyroc/god/god_list"
)

func Test_IsIn(t *testing.T) {
	as := NewAssert(t)

	as.True(god_list.IsStringIn("1", []string{"1"}))
	as.True(god_list.IsBoolIn(false, []bool{false}))
	as.True(god_list.IsIntIn(1, []int{1}))
	as.True(god_list.IsInt8In(1, []int8{1}))
	as.True(god_list.IsInt16In(1, []int16{1}))
	as.True(god_list.IsInt32In(1, []int32{1}))
	as.True(god_list.IsInt64In(1, []int64{1}))
	as.True(god_list.IsUintIn(1, []uint{1}))
	as.True(god_list.IsUint8In(1, []uint8{1}))
	as.True(god_list.IsUint16In(1, []uint16{1}))
	as.True(god_list.IsUint32In(1, []uint32{1}))
	as.True(god_list.IsUint64In(1, []uint64{1}))

	as.False(god_list.IsStringIn("1", []string{"2"}))
	as.False(god_list.IsBoolIn(false, []bool{true}))
	as.False(god_list.IsIntIn(1, []int{2}))
	as.False(god_list.IsInt8In(1, []int8{2}))
	as.False(god_list.IsInt16In(1, []int16{2}))
	as.False(god_list.IsInt32In(1, []int32{2}))
	as.False(god_list.IsInt64In(1, []int64{2}))
	as.False(god_list.IsUintIn(1, []uint{2}))
	as.False(god_list.IsUint8In(1, []uint8{2}))
	as.False(god_list.IsUint16In(1, []uint16{2}))
	as.False(god_list.IsUint32In(1, []uint32{2}))
	as.False(god_list.IsUint64In(1, []uint64{2}))
}
