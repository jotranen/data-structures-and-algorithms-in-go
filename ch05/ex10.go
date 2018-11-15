package ch05

import "sort"

// Find kth smallest element in the union of two sorted list.
func KthSmaller(in1 []int, in2 []int, k int) int {
	cpy := make([]int, len(in1)+len(in2))

	var i int
	for _, v := range in1 {
		cpy[i] = v
		i++
	}
	for _, v := range in2 {
		cpy[i] = v
		i++
	}

	sort.Ints(cpy)

	return cpy[k]
}
