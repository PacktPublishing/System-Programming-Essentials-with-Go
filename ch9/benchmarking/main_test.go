package benchmark

import "testing"

// Benchmark for the Fibonacci function
func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

// Benchmark for the Sum function
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

// Sub-benchmark for the Sum function with different cases
func BenchmarkSumSub(b *testing.B) {
	cases := []struct {
		name string
		a, b int
	}{
		{"small", 1, 2},
		{"large", 1000, 2000},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(c.a, c.b)
			}
		})
	}
}
