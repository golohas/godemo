package slice_test

import (
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	var sliceInt []int
	var b bool
	var s string
	t.Log("int slice: ", sliceInt)
	t.Log("bool b: ", b)
	t.Log("string s: ", s)
}

func TestMake(t *testing.T) {
<<<<<<< HEAD
	var s = make([]int, 4)
	t.Log(s)
	t.Logf("s len: %d, cap: %d", len(s), cap(s))
=======
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

func TestAppend1(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 8, 9}
	t.Logf("s1 addr: %p", &s1)
	s1 = append(s1, s2...)
	t.Log("s1: ", s1)
	t.Logf("s1 len:%d, s1 cap:%d,s1 address:%p ", len(s1), cap(s1), &s1)
}
func TestCopySlice(t *testing.T) {
	var s4 = []int{1, 2, 3, 4}
	var s5 = make([]int, len(s4), cap(s4))
	copy(s5, s4)
	t.Log(s5)
	s5[0] = 4
	t.Log(s5)
}

func TestDelSlice(t *testing.T) {
	var s6 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s6 = append(s6[:1], s6[2:]...)
	t.Log(s6)
}

func TestSliceSorted(t *testing.T) {
	var nSlice = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(nSlice); i++ {
		flag := false
		for j := 0; j < len(nSlice)-i-1; j++ {
			nSlice[j], nSlice[j+1] = nSlice[j+1], nSlice[j]
			flag = true
		}
		if !flag {
			break
		}
	}
	t.Log(nSlice)
}

func TestSorted(t *testing.T) {
	var s = []int{9, 8, 7, 6, 5, 4}
	sort.Ints(s)
	t.Log(s)
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	t.Log(s)
}

func TestSortedSlice(t *testing.T) {
	s := []int{1, 4, 5, 3, 2}
	t.Log(s)
	s = sort.IntSlice(s)
	t.Log(s)
	sort.Ints(s)
	t.Log(s)

>>>>>>> 86969309b12574a13edb3825cd75246696139d4f
}
