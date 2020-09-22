package map_test

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestMapMake(t *testing.T) {
	t.Log("this is a test log")
	var n1 float64 = 3.1
	var n2 float64 = 4.2
	d1 := decimal.NewFromFloat(n1).Add(decimal.NewFromFloat(n2))
	fmt.Println(d1)
	t.Log(d1)
}

func TestSwap(t *testing.T) {
	a := 2
	b := 3
	b = a ^ b
	t.Log("b is: ", b)
	a = b ^ a
	t.Log("a is: ", a)
	b = a ^ b
	t.Log("b is: ", b)
}
