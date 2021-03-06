package ch01

import "testing"

func TestEx01(t *testing.T) {
	for _, c := range []struct {
		in   []float64
		want float64
	}{
		{in: []float64{}, want: 0},
		{in: []float64{1, 2, 3, 4, 5}, want: 3},
		{in: []float64{99}, want: 99},
	} {
		got := Average(c.in)
		if got != c.want {
			t.Errorf("Average(%v) == %f, want %f", c.in, got, c.want)
		}
	}
}
