package ch05

// Give a list of 0's, 1's and 2's, write a program to separate 0's, 1's and 2's.
func Separate_0_1_2(in []int) (zeros, ones, twos int) {

	for _, val := range in {
		if val == 0 {
			zeros++
		} else if val == 1 {
			ones++
		} else {
			twos++
		}
	}
	return zeros, ones, twos
}
