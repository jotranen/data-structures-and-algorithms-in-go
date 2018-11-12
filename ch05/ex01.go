package ch05

import "errors"

var errFindFirstRepeated = errors.New("not found")

// Given a list of n elements, find the first repeated element
func FindFirstRepeated(in []int) (int, error) {

	hmap := make(map[int]bool)

	for _, val := range in {
		if hmap[val] == true {
			return val, nil
		} else {
			hmap[val] = true
		}
	}

	return 0, errFindFirstRepeated
}
