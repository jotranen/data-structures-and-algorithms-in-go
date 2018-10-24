package ch01

// Find the sum of all the elements of a two dimensional list
func Sum(dl [][]int) int {
	var sum int
	for _, y := range dl {
		for _, x := range y {
			sum += x
		}
	}

	return sum
}
