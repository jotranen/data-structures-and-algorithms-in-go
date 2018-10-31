package ch01

import (
	"testing"
)

func TestEx14(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
	}{
		{in: []int{5, 6, 1, 2, 3, 4}, want: 6},
		{in: []int{1, 2, 3, 4}, want: 4},
		{in: []int{2, 1}, want: 2},
	} {

		got := FindMaxElement(c.in)

		// make it a bit more testable

		if got != c.want {
			t.Errorf("FindMaxElement(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
