package ch05

// Give a list of 1's and 0's, write a program to separate 0's and 1's.
// Hint: QuickSelect, counting
func Separate_0_1(in []int) (zeros, ones int) {

	for _, val := range in {
		if val == 0 {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
