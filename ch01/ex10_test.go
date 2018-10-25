package ch01

import (
	"testing"
)

func TestEx10(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want []int
	}{
		{in: []int{1}, want: []int{1}},
		{in: []int{2, 2, 1}, want: []int{1, 2, 2}},
		{in: []int{2, 2, 1}, want: []int{1, 2, 2}},
		{in: []int{2, 2, 1, 3}, want: []int{3, 1, 2, 2}},
	} {
		got := Reverse(c.in)

		for i := range got {
			if c.want[i] != got[i] {
				t.Errorf("Reverse(%v) == %v, want %v", c.in, got, c.want)
				break
			}
		}
	}
}
