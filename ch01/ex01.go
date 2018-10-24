package ch01

// Find average of all the elements in the list
func average(l []float64) float64 {
	var sum float64
	for _, value := range l {
		sum += value
	}

	return sum/float64(len(l))
}

