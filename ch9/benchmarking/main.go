package benchmark

// Fib returns the nth number in the Fibonacci sequence
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// Sum returns the sum of two integers
func Sum(a, b int) int {
	return a + b
}
