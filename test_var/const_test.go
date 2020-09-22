package testvar

import (
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	var pi = math.Pi
	t.Log(pi)
	t.Logf("%f\n", pi)
	t.Logf("%.2f\n", pi)
}
