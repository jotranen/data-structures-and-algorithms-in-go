package ch01

// Print all the maximum's in a list
// (A value is maximum if the value before and after its index are smaller than it is or does not exist)
// Hint: start traversing list form the end and keep track of the max element.
// If we encounter an element whose value is greater then max, print the element and update max
func Maximums(in []int) []int {

	if len(in) == 0 {
		return []int{}
	}
	if len(in) == 1 {
		return []int{in[0]}
	}

	var maximums []int
	var max int
	var prev bool

	max = in[len(in)-1]
	prev = false

	for i := len(in) - 1; i >= 0; i-- {
		if i == 0 && !prev && in[0] > in[1] {
			maximums = append(maximums, in[0])
		} else if in[i] < max && !prev {
			maximums = append(maximums, max)
			prev = true
		} else if in[i] > max {
			max = in[i]
			prev = false
		} else {
			prev = false
		}
	}

	return maximums
}
