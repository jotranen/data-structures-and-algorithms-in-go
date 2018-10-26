package ch01

// Given and array of 0s, 1s, and 2s. When need to stort it so that all the 0s are before all the 1s and all the 1s are before 2s.
// Hint: same ass above first think of 0s and 1s as one groupd and move all the 2s on the right side. Then do a second pass over the array to sort 0s and 1s.
func Sort_0_1_2(in []int) []int {
	cpy := make([]int, len(in))
	copy(cpy, in)

	for start, end := 0, len(cpy)-1; start <= end; end-- {
		if (cpy[start] == 1 || cpy[start] == 2) && cpy[end] == 0 {
			cpy[start], cpy[end] = cpy[end], cpy[start]
		}
		if start < end && cpy[start+1] != 0 {
			start++
		}
	}
	for start, end := 0, len(cpy)-1; start <= end; start++ {
		if cpy[start] == 2 && cpy[end] != 2 {
			cpy[start], cpy[end] = cpy[end], cpy[start]
			end--
		}
	}
	return cpy
}
