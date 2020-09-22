package testvar

import (
	"testing"
	"unsafe"
)

const (
	n1, n2 = iota + 1, iota + 2
	n3, n4
	n5, n6
)

func TestVar(t *testing.T) {
	var name = "John"
	t.Log(name)
}

func TestConst(t *testing.T) {
	t.Log(n1, n2, n3, n4, n5, n6)
}

func TestBian(t *testing.T) {
	var a int8 = 15
	t.Logf("num=%v type:%T\n", a, a)
	t.Log(unsafe.Sizeof(a))
}
