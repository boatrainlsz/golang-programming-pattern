package datastructure

import "testing"

type Item1 struct {
	id  int
	val [1024]byte
}

// genItems 生成指定长度 []*Item1 切片
func genItems(n int) []*Item1 {
	items := make([]*Item1, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item1{id: i})
	}
	return items
}

func BenchmarkIndexPointer(b *testing.B) {
	items := genItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(items); k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := genItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
