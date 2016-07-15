package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var bitPattern uint64 = 1<<0 | 1<<1 | 1<<4
	if PopCount(bitPattern) != 3 {
		t.Errorf("PopCount(%064b) == 3 failed", bitPattern)
	}
}
