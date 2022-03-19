package datastructure

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkJoinStrWithStringsJoin(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2, s3}, "")
	}
}

func BenchmarkJoinStrWithStringsBuilder(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		_, _ = builder.WriteString(s1)
		_, _ = builder.WriteString(s2)
		_, _ = builder.WriteString(s3)
	}
}

func BenchmarkJoinStrWithBytesBuffer(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		_, _ = buffer.WriteString(s1)
		_, _ = buffer.WriteString(s2)
		_, _ = buffer.WriteString(s3)
	}
}

func BenchmarkJoinStrWithByteSlice(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var bys []byte
		bys = append(bys, s1...)
		bys = append(bys, s2...)
		_ = append(bys, s3...)
	}
}

func BenchmarkJoinStrWithByteSlicePreAlloc(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		bys := make([]byte, 0, 9)
		bys = append(bys, s1...)
		bys = append(bys, s2...)
		_ = append(bys, s3...)
	}
}
