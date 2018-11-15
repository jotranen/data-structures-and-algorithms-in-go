package ch05

import "sort"

// Find min in the sorted rotated list.

func FindMin(in []int) int {
	cpy := make([]int, len(in))
	copy(cpy, in)

	sort.Ints(cpy)

	return cpy[0]

}
