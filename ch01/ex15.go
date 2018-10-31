package ch01

import (
	"sort"
)

// Given a slice with 'n' elemenents & a value 'x', find two elements in the list that sums to 'x'
// Hint:
//	Approach 1: sort the list
//	Approach 2: Using a hash table
func Find2Sum(in []int, x int) []int {

	var res []int
	cpy := make([]int, len(in))
	copy(cpy, in)

	sort.Slice(cpy, func(i, j int) bool {
		return in[i] < in[j]
	})

	var lsIndex int
	// find largest smaller than x value
	for i := len(cpy) - 1; i >= 0; i-- {
		if cpy[i] < x {
			lsIndex = i
			break
		}
	}

	// not found
	if lsIndex == 0 {
		return res
	}

	// find first part that sums up to x
	iter := func(arr []int, start int, end int) int {
		for i := start; i < end; i++ {
			if arr[i] + arr[lsIndex] == x {
				return i
			}
		}
		return -1
	}

	sIndex := -1
	for lsIndex > 0 {
		sIndex = iter(cpy, 0, lsIndex)
		if sIndex != -1 {
			break
		}
		lsIndex--
	}

	if sIndex == -1 {
		return res
	}

	res = []int{cpy[sIndex], cpy[lsIndex]}

	return res
}