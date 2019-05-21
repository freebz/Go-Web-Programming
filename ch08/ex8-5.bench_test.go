// Listing 8.5  Benchmark testing

package main

import (
	"testing"
)

// benchmarking the decode function
func BenchmarkingDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}
