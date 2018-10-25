package ch01

import "fmt"

// Given a list of 0s and 1s. We need to sort it so that all the 0s are before all 1s.
//	Hint: Use two variables, start and end.
//  Start set to 0 and end set to (n-1). Increment start and decrement end.
//	Swap the values stored at arr[start] and arr[end] only when arr[start] == 1 and arr[end] == 0.
//  Stop when start is equal to end or start is greater than end.

func Sort_0_1(in []int) []int {

	cpy := make([]int, len(in))
	copy(cpy, in)

	for start, end := 0, len(cpy)-1; start <= end; end-- {
		if cpy[start] == 1 && cpy[end] == 0 {
			cpy[start], cpy[end] = cpy[end], cpy[start]
			fmt.Println(cpy)
		}
		if start < end && cpy[start+1] != 0 {
			start++
		}
	}
	return cpy
}
