package testvar

import "testing"

const (
	A = "A"
	B
	C
)

func TestConst1(t *testing.T) {
	t.Log(A, B, C)
}
