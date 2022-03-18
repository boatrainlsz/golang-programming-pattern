package datastructure

import (
	"reflect"
	"testing"
)

// DeleteSliceElms 从切片中过滤指定元素。注意：不修改原切片。
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	// 构建 map set。
	m := make(map[interface{}]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	v := reflect.ValueOf(i)
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface()
}

// DeleteU64liceElms 从 []uint64 过滤指定元素。注意：不修改原切片。
func DeleteU64liceElms(i []uint64, elms ...uint64) []uint64 {
	// 构建 map set。
	m := make(map[uint64]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	t := make([]uint64, 0, len(i))
	for _, v := range i {
		if _, ok := m[v]; !ok {
			t = append(t, v)
		}
	}
	return t
}

func BenchmarkDeleteSliceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []interface{}{uint64(1), uint64(3), uint64(5), uint64(7), uint64(9)}
	for i := 0; i < b.N; i++ {
		_ = DeleteSliceElms(slice, elms...)
	}
}

func BenchmarkDeleteU64liceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []uint64{1, 3, 5, 7, 9}
	for i := 0; i < b.N; i++ {
		_ = DeleteU64liceElms(slice, elms...)
	}
}
