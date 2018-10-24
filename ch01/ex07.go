package ch01

// Write a method to compute Sum(N) = 1+2+3+...+N

func SumN(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n + SumN(n-1)
}
