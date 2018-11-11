package ch06

import "testing"

func BenchmarkBubbleSort1(b *testing.B) {

	in := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		in = append(in, i)
	}
	benchmarkBubbleSort(b, in)
}

func BenchmarkBubbleSort2(b *testing.B) {

	in := make([]int, 1000)
	for i := 1000; i > 0; i-- {
		in = append(in, i)
	}
	benchmarkBubbleSort(b, in)
}

func benchmarkBubbleSort(b *testing.B, in []int) {

	for n := 0; n < b.N; n++ {

		BubbleSort(in, less)
	}
}
