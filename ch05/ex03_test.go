package ch05

import "testing"

func TestSegregatePositiveAndNegativeNumbers(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want []int
	}{
		{in: []int{-1, 1, -2, 2}, want: []int{-1, -2, 1, 2}},
	} {
		got := SegregatePositiveAndNegativeNumbers(c.in)

		if len(got) != len(c.want) {
			t.Errorf("SegregatePositiveAndNegativeNumbers(%v) = %v, want: %v", c.in, got, c.want)
		} else {
			for i, val := range got {
				if val != c.want[i] {
					t.Errorf("SegregatePositiveAndNegativeNumbers(%v) = %v, want: %v", c.in, got, c.want)
					break
				}
			}
		}
	}
}
