package ch01

import "testing"

func TestEx02(t *testing.T) {
	for _, c := range []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{}}, want: 0},
		{in: [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}, want: 30},
		{in: [][]int{{99}}, want: 99},
	} {
		got := Sum(c.in)
		if got != c.want {
			t.Errorf("Sum(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
