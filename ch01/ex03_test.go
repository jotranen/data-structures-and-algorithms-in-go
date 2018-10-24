package ch01

import "testing"

func TestEx03(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
	}{
		{in: []int{}, want: 0},
		{in: []int{1, 2, 3, 4, 5}, want: 5},
		{in: []int{99}, want: 99},
		{in: []int{7, 2, 11, 4, 5}, want: 11},
		{in: []int{7, 7, 11, 11, 5}, want: 11},
	} {
		got := Largest(c.in)
		if got != c.want {
			t.Errorf("Average(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
