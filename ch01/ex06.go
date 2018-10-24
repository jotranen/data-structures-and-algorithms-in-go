package ch01

// Using AllPermutation function discussed before, write a function, which give only distinct solutions
//func Permutation(data []int, i int, length int) {
//	if length == 1 {
//		PrintSlice(data)
//		return
//	}
//	for j := i; j < length; j++ {
//		swap(data, i, j)
//		Permutation(data, i+1, length)
//		swap(data, i, j)
//	}
//}
//
//func swap(data []int, x int, y int) {
//	data[x], data[y] = data[y], data[x]
//}
//
//func main() {
//	var data [5]int
//
//	for i := 0; i < 5; i++ {
//		data[i] = i
//	}
//
//	Permutation(data[:], 0, 5)
//}
//
//func PrintSlice(data []int) {
//	fmt.Printf("%v :: len=%d cap-=%d \n", data, len(data), cap(data))
//}
