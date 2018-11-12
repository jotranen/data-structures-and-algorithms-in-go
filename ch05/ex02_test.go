package ch05

import (
	"testing"
)

func TestFindFirstRepeated(t *testing.T) {
	for _, c := range []struct {
		in   []int
		res  int
		want []int
		err  error
	}{
		{in: []int{1, 3, 4}, res: 8, want: []int{1, 3, 4}, err: nil},
		{in: []int{1, 3, 3}, res: 8, want: nil, err: errFindThreeFirstToSum},
	} {
		got, err := FindThreeFirstToSum(c.in, c.res)

		arrEquals := true

		if len(c.want) == len(got) {
			for i, val := range c.want {
				if val != got[i] {
					arrEquals = false
					break
				}
			}
		} else {
			arrEquals = false
		}

		if err != c.err || !arrEquals {
			t.Errorf("FindThreeFirstToSum(%v, %d) = (%v, %v), expected (%v, %v)", c.in, c.res, got, err, c.want, c.err)
		}
	}
}
