package datastructure

import (
	"fmt"
	"testing"
)

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
}

func TestChannelWithoutData(t *testing.T) {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
	close(ch)
}
