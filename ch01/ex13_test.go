package ch01

import (
	"sort"
	"testing"
)

func TestEx13(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want []int
	}{
		{in: []int{0, 1, 2}, want: []int{}},
		{in: []int{1, 1, 2}, want: []int{1}},
		{in: []int{6, 6, 1, 1, 2, 1, 1, 0, 2, 7, 7}, want: []int{1, 2, 6, 7}},
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

		got := FindDuplicated(c.in)

		// make it a bit more testable
		sort.Ints(got)

		if !intComparer(got, c.want) {
			t.Errorf("FindDuplicated(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
