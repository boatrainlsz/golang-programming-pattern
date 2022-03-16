package main

import "fmt"

func gReduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := gReduce(nums, 0, func(result, elem int) int {
		return result + elem
	})
	fmt.Printf("Sum = %d \n", sum)
}
