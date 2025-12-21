package main

import (
	"runtime"
	"testing"
)

func BenchmarkBadLinkedList(b *testing.B) {
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		badLinkedList(1000) // 1000 stack objects!
	}
}

func BenchmarkGoodArray(b *testing.B) {
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goodArray(1000) // 1 stack object!
	}
}
