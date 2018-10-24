package ch01

// Find the second largest number in the list
func SecondLargest(in []int) int {

	// -1 if can not find second largest
	if len(in) < 2 || (len(in) == 2 && in[0] == in[1]) {
		return -1
	}

	max, max2 := func(a int, b int) (max int, min int) {
		if a > b {
			return a, b
		} else {
			return b, a
		}
	}(in[0], in[1])

	for _, value := range in {
		if value > max {
			max2 = max
			max = value
		} else if value > max2 && value < max {
			max2 = value
		}
	}

	if max2 == max {
		return -1
	}

	return max2
}
