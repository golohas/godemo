package slice_test

import "testing"

func TestSlice(t *testing.T) {
	var sliceInt []int
	var b bool
	var s string
	t.Log("int slice: ", sliceInt)
	t.Log("bool b: ", b)
	t.Log("string s: ", s)
}

func TestMake(t *testing.T) {
	var s = make([]int, 4)
	t.Log(s)
	t.Logf("s len: %d, cap: %d", len(s), cap(s))
}
