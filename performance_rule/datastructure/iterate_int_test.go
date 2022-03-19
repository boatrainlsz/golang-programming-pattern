package datastructure

import (
	"math/rand"
	"testing"
	"time"
)

// genRandomIntSlice 生成指定长度的随机 []int 切片
func genRandomIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkIndexIntSlice(b *testing.B) {
	nums := genRandomIntSlice(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(nums); k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := genRandomIntSlice(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}
