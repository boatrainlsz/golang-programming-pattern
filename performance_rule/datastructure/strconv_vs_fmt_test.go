package datastructure

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkFmt(b *testing.B) {
	// Bad
	for i := 0; i < b.N; i++ {
		fmt.Sprint(rand.Int())
	}

}
func BenchmarkConv(b *testing.B) {
	// Good
	for i := 0; i < b.N; i++ {
		strconv.Itoa(rand.Int())
	}

}
