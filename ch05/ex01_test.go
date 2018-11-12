package ch05

import (
	"testing"
)

func TestEx01(t *testing.T) {
	for _, c := range []struct {
		in   []int
		want int
		err  error
	}{
		{in: []int{1, 3, 4, 5}, want: 0, err: errFindFirstRepeated},
		{in: []int{1, 3, 1, 5}, want: 1, err: nil},
		{in: []int{1, 3, 1, 5, 5}, want: 1, err: nil},
	} {

		got, err := FindFirstRepeated(c.in)
		if err != c.err || got != c.want {
			t.Errorf("Func(%v) == (%d, %v), want: (%d, %v)", c.in, got, err, c.want, c.err)
		}
	}
}
