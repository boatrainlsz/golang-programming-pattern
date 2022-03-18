package datastructure

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// Ntohl 将网络字节序的 uint32 转为主机字节序。
func Ntohl(bys []byte) uint32 {
	r := bytes.NewReader(bys)
	var v uint32
	_ = binary.Read(r, binary.BigEndian, &v)
	return v
}

func BenchmarkNtohl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Ntohl([]byte{0x7f, 0, 0, 0x1})
	}
}

func BenchmarkNtohlNotUseBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NtohlNotUseBinary([]byte{0x7f, 0, 0, 0x1})
	}
}

func NtohlNotUseBinary(bys []byte) uint32 {
	return uint32(bys[3]) | uint32(bys[2])<<8 | uint32(bys[1])<<16 | uint32(bys[0])<<24
}
