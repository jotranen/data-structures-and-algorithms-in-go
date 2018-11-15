package ch05

import "testing"

func TestFindNumberOrSlot(t *testing.T) {
	for _, c := range []struct {
		in     []int
		number int
		want   int
		ok     bool
	}{
		{in: []int{1, 2, 3, 5}, number: 4, want: 3, ok: false},
		{in: []int{1, 2, 3, 5, 6}, number: 4, want: 3, ok: false},
		{in: []int{1, 2, 3, 5, 6, 7}, number: 4, want: 3, ok: false},
	} {
		got, ok := FindNumberOrSlot(c.in, c.number)

		if got != c.want || ok != c.ok {
			t.Errorf("FindNumberOrSlot(%v, %d) = (%d, %t), want: (%d, %t)", c.in, c.number, got, ok, c.want, c.ok)
		}
	}
}
