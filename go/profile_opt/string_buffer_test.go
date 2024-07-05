package main

import "testing"

func BenchmarkNewBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewString()
	}
}
