package datastructure

import (
	"testing"
)

const size = 1000000

func BenchmarkSliceWithoutCapacity(b *testing.B) {
	// Bad
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkSliceWithCapacity(b *testing.B) {
	// Good
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}
