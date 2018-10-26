package ch01

import "testing"

func TestEx12(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want []int
	}{
		{in: []int{0, 1, 2}, want: []int{0, 1, 2}},
		{in: []int{1, 0, 2}, want: []int{0, 1, 2}},
		{in: []int{1, 1, 2, 1, 1, 0}, want: []int{0, 1, 1, 1, 1, 2}},
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

		got := Sort_0_1_2(c.in)

		if !intComparer(got, c.want) {
			t.Errorf("Sort_0_1_2(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
