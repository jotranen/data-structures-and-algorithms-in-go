package ch01

// Find the largest element in the list
func Largest(in []int) int {
	var max int
	for _, value := range in {
		if value > max {
			max = value
		}
	}
	return max
}
