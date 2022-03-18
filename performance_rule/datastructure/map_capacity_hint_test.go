package datastructure

import (
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkMapWithoutCapacityHint(b *testing.B) {
	// Bad
	m := make(map[string]os.FileInfo)

	files, _ := ioutil.ReadDir("C:\\Windows\\WinSxS")
	for _, f := range files {
		m[f.Name()] = f
	}
	// m 是在没有大小提示的情况下创建的； 在运行时可能会有更多分配。
}

func BenchmarkMapWithCapacityHint(b *testing.B) {
	// Good
	files, _ := ioutil.ReadDir("C:\\Windows\\WinSxS")

	m := make(map[string]os.FileInfo, len(files))
	for _, f := range files {
		m[f.Name()] = f
	}
	// m 是有大小提示创建的；在运行时可能会有更少的分配。
}
