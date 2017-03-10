package algorithms

import (
	"testing"
)

func TestPrimeHash(t *testing.T) {
	if PrimeProduct("bad") != 2*3*7 {
		t.Error("didnt work")
	}
	if PrimeProduct("datamine") != PrimeProduct("animated") {
		t.Error("didnt work")
	}
}