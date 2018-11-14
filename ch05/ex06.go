package ch05

import "errors"

// Give a list of elements in monotonically increasing with both negative and positive numbers.
// Write an algorithm to find the point at which list becomes positive.

var errFindMonotonicalList0Point = errors.New("not found")

func FindMonotonicalList0Point(in []int) (int, error) {
	for i := range in {
		if in[i] > 0 {
			return i, nil
		}
	}

	return 0, errFindMonotonicalList0Point
}
