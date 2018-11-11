package ch06

func BubbleSort(arr []int, comp func(int, int) bool) {

	for i := 0; i <  len(arr) - 1 ; i++ {
		for j := 0; j < len(arr) - 1 - i - 1; j++ {
			if comp(arr[j], arr[j+1]) {
				// swap
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
