package ch01

// Given a list of intervals, merge all overlapping intervals Input{[1,4], [3,6],[8,10]} Output{[1,6], [8,10]}
func MergeIntervals(in [][]int) [][]int {

	var max int

	var res [][]int

	// find end point for intervals
	for _, array := range in {
		if array[len(array)-1] > max {
			max = array[len(array)-1]
		}
	}

	// default for bool is false
	intervals := make([]bool, max+1)

	// create array and flag active entries in the who range
	for _, array := range in {
		start := array[0]
		end := array[len(array)-1]
		for start <= end {
			intervals[start] = true
			start++
		}
	}

	var active bool
	var start int
	var end int

	// split intervals based on active flag
	for i := 1; i <= max; i++ {
		if !active {
			if intervals[i] {
				active = true
				start = i
				end = i
			}
		} else {
			if intervals[i] && i == len(intervals)-1 {
				res = append(res, []int{start, end + 1})
				active = false
			} else if intervals[i] {
				end = i
			} else {
				active = false
				res = append(res, []int{start, end})
			}
		}
	}

	return res
}
