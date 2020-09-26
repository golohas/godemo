package arr_test

import "testing"

func TestArr(t *testing.T) {
	a := [...]int{1: 1, 3: 5}
	t.Log(a)
	b := [...]int{1, 2}
	t.Log(b, len(b))
}

func TestArrFor(t *testing.T) {
	a := [...]int{1: 1, 3: 5}
	t.Log("method one:")
	for i := 0; i < len(a); i++ {
		t.Log(a[i], " ")
	}
	t.Log("method two:")
	for _, value := range a {
		t.Log(value, " ")
	}
}
