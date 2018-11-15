package ch05

// Give a sorted list, find a given number. If found return the index if not, find the index of that number if it is
// inserted into the list.
func FindNumberOrSlot(in []int, want int) (int, bool) {

	i, j := 0, len(in)-1

	var mid int

	for i < j {
		mid = (i + j) / 2
		if in[mid] == want {
			break
		} else if in[mid] < want {
			i = mid + 1
		} else {
			j = mid
		}
	}

	if in[mid] == want {
		return mid, true
	} else {
		if i > 0 && in[mid] < want {
			return mid + 1, false
		} else {
			return mid, false
		}
	}
}
