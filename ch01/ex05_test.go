package ch01

import "testing"

func TestEx05(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
	}{
		{in: []int{}, want: -1},
		{in: []int{1, 2, 3, 4, 5}, want: 4},
		{in: []int{99}, want: -1},
		{in: []int{7, 2, 11, 4, 5}, want: 7},
		{in: []int{7, 7, 11, 11, 5}, want: 7},
	} {
		got := SecondLargest(c.in)
		if got != c.want {
			t.Errorf("SecondLargest(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
