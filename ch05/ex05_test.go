package ch05

import "testing"

func TestSeparate_0_1_2(t *testing.T) {
	for _, c := range []struct {
		in    []int
		zeros int
		ones  int
		twos  int
	}{
		{in: []int{0, 1, 0, 1, 2, 0, 1}, zeros: 3, ones: 3, twos: 1},
	} {
		zeros, ones, twos := Separate_0_1_2(c.in)

		if zeros != c.zeros || ones != c.ones {
			t.Errorf("Separate_0_1(%v) = (%d, %d, %d), want: (%d, %d, %d)", c.in, zeros, ones, twos, c.zeros, c.ones, c.twos)
		}
	}
}
