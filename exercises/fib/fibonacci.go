package fib

func Fibonacci(n int) int {
	var sequence = []int{0, 1}
	for i := 2; i <= n; i++ {
		sequence = append(sequence, sequence[i-2]+sequence[i-1])
	}
	return sequence[n]
}

// way worse performance
//func Fibonacci(n int) int {
//	if n < 2 {
//		return n
//	}
//	return Fibonacci(n-2) + Fibonacci(n-1)
//}
