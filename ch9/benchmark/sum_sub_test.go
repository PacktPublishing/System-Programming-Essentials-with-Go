package benchmark

import "testing"

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
