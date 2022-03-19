package datastructure

import "testing"

type Item struct {
	id  int
	val [1024]byte
}

func BenchmarkIndexStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := 0; j < len(items); j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
