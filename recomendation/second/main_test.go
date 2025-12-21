package main

import (
	"testing"
)

const N = 100_000_000

func BenchmarkBigWithPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = useBigWithPtr(N)
	}
}

func BenchmarkBigNoPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = useBigNoPtr(N)
	}
}
