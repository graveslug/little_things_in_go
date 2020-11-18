package main

import "testing"

//Benchmark just testing how long it takes to finish
func Benchmark(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1Echo()
	}
}
