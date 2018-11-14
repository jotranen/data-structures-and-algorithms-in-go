package ch05

import "testing"

func TestFindMonotonicalList0Point(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
		err  error
	}{
		{in: []int{-10, -9, 0, 1, 2}, want: 3},
	} {
		got, err := FindMonotonicalList0Point(c.in)

		if got != c.want || err != c.err {
			t.Errorf("TestFindMonotonicalList0Point(%v) = (%d, %v) want: (%d,%v) ", c.in, got, err, c.want, c.err)
		}
	}
}
