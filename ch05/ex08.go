package ch05

import "sort"

// Find max in sorted rotated list.

func FindMax(in []int) int {
	cpy := make([]int, len(in))
	copy(cpy, in)

	sort.Ints(cpy)

	return cpy[len(cpy)-1]

}
