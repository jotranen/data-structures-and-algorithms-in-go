package ch01

import "testing"

func TestEx16(t *testing.T) {
	for _, c := range []struct {
		in   string
		want int
	}{
		{in: "1984", want: 22},
	} {
		got := SumNumbers(c.in)
		if got != c.want {
			t.Errorf("SumNumbers(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}