package ch01

// Find the smallest element in the list
func Min(in []int) int {

	if len(in) == 0 {
		return 0
	}

	min := in[0]

	for _, value := range in {
		if value < min {
			min = value
		}
	}

	return min
}
