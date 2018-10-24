package ch01

// Find average of all the elements in the list
func Average(l []float64) float64 {

	if len(l) == 0 {
		return 0
	}

	var sum float64
	for _, value := range l {
		sum += value
	}

	return sum / float64(len(l))
}
