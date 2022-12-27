package main

import (
	"testing"
)

var num = 40

func BenchmarkBasicRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicRecursion(num)
	}
}

func BenchmarkExtraCaseRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtraCaseRecursion(num)
	}
}

func BenchmarkTCORecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TCORecursion(num)
	}
}
