package ch01

// Reverse a list in-place (You cannot use any additional list or other variables Space Complexity should be O(1))
// Hint: Use two variable, start and end. Start set to 0 and end set to (n-1). Increment start and decrement end. Swap the values stored at arr[start] and arr[end]. Stop when
// start is equal to end or start is greater than end.
func Reverse(in []int) []int {

	start := 0
	end := len(in) - 1

	for start <= end {
		in[start], in[end] = in[end], in[start]
		start++
		end--
	}

	return in
}
