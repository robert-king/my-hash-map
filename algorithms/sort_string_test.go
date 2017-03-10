package algorithms

import (
	"testing"
)

func TestSortString(t *testing.T) {
	if SortString("cba") != "abc" {
		t.Error("didnt work")
	}
}