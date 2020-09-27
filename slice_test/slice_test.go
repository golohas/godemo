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
	var s = make([]int, 4, 8)
	t.Log(s)
	for i := 0; i < len(s); i++ {
		s[i] = i + 2
	}
	t.Log(s)
	s1 := s[2:3]
	t.Log("s1 is: ", s1)
	t.Logf("s1 len: %d, s1 cap: %d", len(s1), cap(s1))
}

func TestAppend(t *testing.T) {
	s := []int{1, 2, 3, 4}
	t.Logf("s'address :%p", &s)
	s = append(s, 5)
	t.Log(s)
	t.Logf("s'address :%p", &s)
}
