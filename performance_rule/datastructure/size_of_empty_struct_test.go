package datastructure

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeOfEmptyStructTest(t *testing.T) {
	fmt.Println(unsafe.Sizeof(struct{}{}))
}
