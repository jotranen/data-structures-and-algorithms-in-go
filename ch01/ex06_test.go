package ch01

import (
	"testing"
)

func TestEx06(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
	}{
		{in: []int{1, 1}, want: 1},
		{in: []int{1, 2, 3}, want: 6},
		{in: []int{1, 1, 3}, want: 3},
	} {
		got := DistinctPermutations(c.in)
		if len(got) != c.want {
			t.Errorf("DistinctPermutations(%v) == %v, want %v", c.in, len(got), c.want)
		}
	}
}
