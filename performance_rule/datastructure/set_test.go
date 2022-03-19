package datastructure

import (
	"fmt"
	"testing"
)

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func TestSet(t *testing.T) {
	s := make(Set)
	s.Add("foo")
	s.Add("bar")
	fmt.Println(s.Has("foo"))
	fmt.Println(s.Has("bar"))
}
