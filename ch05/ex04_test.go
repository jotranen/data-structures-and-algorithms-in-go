package ch05

import "testing"

func TestSeparate_0_1(t *testing.T) {
	for _, c := range []struct {
		in    []int
		zeros int
		ones  int
	}{
		{in: []int{0, 1, 0, 1, 0, 1}, zeros: 3, ones: 3},
	} {
		zeros, ones := Separate_0_1(c.in)

		if zeros != c.zeros || ones != c.ones {
			t.Errorf("Separate_0_1(%v) = (%d, %d), want: (%d, %d)", c.in, zeros, ones, c.zeros, c.ones)
		}
	}
}
