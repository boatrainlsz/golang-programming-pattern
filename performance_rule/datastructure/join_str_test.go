package datastructure

import (
	"fmt"
	"testing"
)

/**
从性能出发，兼顾易用可读，如果待拼接的变量不涉及类型转换且数量较少（<=5），行内拼接字符串推荐使用运算符 +，反之使用 fmt.Sprintf()。
*/

// Good
func BenchmarkJoinStrWithOperator(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2 + s3
	}
}

// Bad
func BenchmarkJoinStrWithSprintf(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
}
