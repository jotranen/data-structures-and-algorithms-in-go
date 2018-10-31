package ch01

// Find the maximum element in a sorted and rotated list. Complexity O(log(n))
// Hint: user binary search algorithm
// Comments: 	Must be rotated at some unknown point, otherwise this problem doesn't make much sense
// 				(Largest value is the first element...)
// 				If duplicates are allowed then binary search won't work, and complexity is O(n) worst case.
//				// e.g. {3, 3, 3, 3, 3, 3, 3, 3, 0, 1, 1, 3} and {3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3}
func FindMaxElement(in []int) int {

	res := findMax(in, 0, len(in) - 1)

	return res
}

func findMax(arr []int, start int, end int) int {
	if len(arr) == 0 {
		// empty array
		return -1
	}
	if start == end {
		// end of recursion
		return arr[start]
	}

	mid := start + (end - start)/2

	if start < end && arr[mid+1] < arr[mid] {
		// start is max since next element is smaller
		return arr[mid]
	} else {
		// binary search decision, left or right
		if arr[start] > arr[end] {
			// left hand side
			return findMax(arr, start, mid)
		} else {
			// right hand side
			return findMax(arr, mid+1, end)
		}
	}
}
