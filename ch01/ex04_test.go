package ch01

import "testing"

func TestEx04(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
	}{
		{in: []int{}, want: -1},
		{in: []int{1, 2, 3, 4, 5}, want: 1},
		{in: []int{99}, want: 99},
		{in: []int{7, 2, 11, 2, 5}, want: 2},
		{in: []int{7, 7, 0, 1, 5}, want: 0},
	} {
		got := Min(c.in)
		if got != c.want {
			t.Errorf("Min(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
