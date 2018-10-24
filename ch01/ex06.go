package ch01

import "fmt"

// Using AllPermutation function discussed before, write a function, which give only distinct solutions
func DistinctPermutations(in []int) []string {
	var m map[string]struct{}
	m = make(map[string]struct{})

	Permutation(in, 0, len(in), m)

	var res []string
	for k := range m {
		res = append(res, k)
	}

	return res
}

func Permutation(data []int, i int, length int, m map[string]struct{}) {
	//PrintSlice(data)
	if length == i {
		m[fmt.Sprintf("%v", data)] = struct{}{}
		return
	}
	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length, m)
		swap(data, i, j)
	}
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}
