package ch01

import (
	"testing"
)

func TestEx09(t *testing.T) {
	for _, c := range []struct {
		in   [][]int
		want [][]int
	}{
		{in: [][]int{{1, 4}, {3, 6}, {8, 10}}, want: [][]int{{1, 6}, {8, 10}}},
		{in: [][]int{{1, 4}, {3, 6}}, want: [][]int{{1, 6}}},
		{in: [][]int{{1, 4}}, want: [][]int{{1, 4}}},
	} {
		got := MergeIntervals(c.in)

		for i, value := range got {
			for j, inner := range value {
				if inner != c.want[i][j] {
					t.Errorf("MergeIntervals(%v) == %v, want %v", c.in, got, c.want)
				}
			}
		}
	}
}
