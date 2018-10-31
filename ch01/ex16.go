package ch01

// Write a method to find the sum of every number in an int number. Example input = 1984 output should be 22 (1+9+8+4)
func SumNumbers(in string) int {

	var sum int
	for _, val:= range in {
		sum = sum + int(val) - 48 	// ascii 0 value is 48
	}

	return sum
}