package ch05

// Given a list of -ve and +ve numbers, write a program to separate -ve numbers from the +ve numbers.
// (Segregate positive and negative numbers)
func SegregatePositiveAndNegativeNumbers(in []int) []int {

	cpy := make([]int, len(in))
	copy(cpy, in)

	for i, j := 0, len(cpy)-1; i <= j; i = i {
		if cpy[i] < 0 {
			i++
		} else if cpy[j] > 0 {
			j--
		} else {
			cpy[i], cpy[j] = cpy[j], cpy[i]
		}
	}

	return cpy
}
