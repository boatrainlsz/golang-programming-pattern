package datastructure

import (
	"bytes"
	"testing"
)

func BenchmarkInitiateRepeatedly(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Write([]byte("Hello world"))
	}
}

func BenchmarkInitiateOnce(b *testing.B) {
	var buf bytes.Buffer
	data := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		buf.Write(data)
	}
}
