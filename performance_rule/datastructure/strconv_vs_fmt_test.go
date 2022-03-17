package datastructure

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestStrconvVSFmt(b *testing.B) {
	// Bad
	for i := 0; i < b.N; i++ {
		s := fmt.Sprint(rand.Int())
	}

	// Good
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Int())
	}

}
