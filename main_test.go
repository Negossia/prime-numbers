package main

import (
	"testing"
)

func Benchmark_sieve1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		sieve(100)
	}
}
func BenchmarkGetPrimes1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		GetPrimes1(100)
	}
}

func BenchmarkGetPrimes2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		GetPrimes2(100)
	}
}
