package ch01

import (
	"testing"
)

func TestEx08(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want []int
	}{
		{in: []int{1, 2, 1}, want: []int{2}},
		{in: []int{2, 2, 1}, want: []int{}},
		{in: []int{3, 2, 1}, want: []int{3}},
		{in: []int{3, 2, 1, 8}, want: []int{8, 3}},
		{in: []int{3, 2, 1, 8, 7}, want: []int{8, 3}},
		{in: []int{}, want: []int{}},
		{in: []int{99}, want: []int{99}},
		{in: []int{1, 1, 1, 1, 1, 1, 1, 1}, want: []int{}},
	} {
		got := Maximums(c.in)

		for i := range got {
			if c.want[i] != got[i] {
				t.Errorf("Maximums(%v) == %v, want %v", c.in, got[i], c.want[i])
			}
		}
	}
}
