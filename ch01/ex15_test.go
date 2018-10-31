package ch01

import (
	"sort"
	"testing"
)

func TestEx15(t *testing.T) {
	for _, c := range []struct {
		in   []int
		x 	int
		want []int
	}{
		{in: []int{0, 1, 2}, x: 3, want: []int{1,2}},
		{in: []int{2, 1, 0}, x: 3, want: []int{1,2}},
		{in: []int{2, 0, 1}, x: 3, want: []int{1,2}},
		{in: []int{0, 1, 3, 3, 4, 5, 6, 7, 8, 2}, x: 3, want: []int{1,2}},
	} {

		intComparer := func(a []int, b []int) bool {
			if len(a) != len(b) {
				return false
			}
			for i, v := range a {
				if v != b[i] {
					return false
				}
			}
			return true
		}

		got := Find2Sum(c.in, c.x)

		// make it a bit more testable
		sort.Ints(got)

		if !intComparer(got, c.want) {
			t.Errorf("Find2Sum(%v, %d) == %v want %v", c.in, c.x, got, c.want)
		}
	}
}