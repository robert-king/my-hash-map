package algorithms

import (
	"testing"
)

func TestDifferentBits(t *testing.T) {
	if diffBit(1, 5, 0) {
		t.Error("wrong")
	}
	if !diffBit(1, 2, 1) {
		t.Error("wrong")
	}
	dbs := differentBits(8+2, 3)
	if len(dbs) != 2 || dbs[0] != 0 || dbs[1] != 3 {
		t.Error("wrong")
	}
}

func TestResolveCollisions(t *testing.T) {
	nums := []uint64{1, 2, 4}
	bits := minimumDistinguishingBits(nums)

	if bitScore(4, bits) != 0 {
		t.Error("wrong")
	}

	if maxBitScore(nums, bits) != 2 {
		t.Error("wrong")
	}

}

func TestNewLocator(t *testing.T) {
	a := uint64(34534)
	b := a + 1<<20
	c := b + 1<<21
	d := a + 1<<22
	locator := NewLocator([]uint64{a, b, c, d, 30, 31})
	if len(locator.collisions[a]) != 4 {
		t.Error("these should have collided")
	}
	if locator.Index(30) != 0 {
		t.Error("error")
	}
	if locator.Index(31) != 1 {
		t.Error("error")
	}
}
