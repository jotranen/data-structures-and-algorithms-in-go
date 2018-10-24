package ch01

import "testing"

func TestEx07(t *testing.T) {
	for _, c := range []struct {
		in   int
		want int
	}{
		{in: 1, want: 1},
		{in: 2, want: 3},
		{in: 3, want: 6},
		{in: 10, want: 55},
	} {
		got := SumN(c.in)
		if got != c.want {
			t.Errorf("SumN(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
