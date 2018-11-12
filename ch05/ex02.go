package ch05

import "errors"

// Given a list of n elements, write an algorithm to find three elements in a list whose sum is a given value.
// Hist: Try to do this problem using a brute fore approach. Then try to apply the sorting approach along with brute force approach.
// The time complexity will be O(n2)

var errFindThreeFirstToSum = errors.New("not found")

func FindThreeFirstToSum(in []int, res int) ([]int, error) {
	hm := make(map[int][]int)
	for i := 0; i < len(in)-2; i++ {
		for j := i + 1; j < len(in)-1; j++ {
			for k := i + 2; k < len(in); k++ {
				sum := in[i] + in[j] + in[k]
				hm[sum] = []int{in[i], in[j], in[k]}
			}
		}
	}

	if val, ok := hm[res]; ok {
		return val, nil
	}

	return nil, errFindThreeFirstToSum
}
