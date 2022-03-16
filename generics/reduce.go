package main

import "fmt"

func gReduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

func gCountIf[T any](arr []T, f func(T) bool) int {
	cnt := 0
	for _, elem := range arr {
		if f(elem) {
			cnt += 1
		}
	}
	return cnt
}

//https://bitfieldconsulting.com/golang/generics
//~是类型近似符号
type Sumable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func gSum[T any, U Sumable](arr []T, f func(T) U) U {
	var sum U
	for _, elem := range arr {
		sum += f(elem)
	}
	return sum
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := gReduce(nums, 0, func(result, elem int) int {
		return result + elem
	})
	fmt.Printf("Sum = %d \n", sum)

	old := gCountIf(employees, func(e Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people(>40): %d\n", old)
	// ld people(>40): 2

	high_pay := gCountIf(employees, func(e Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people(>6k): %d\n", high_pay)
	//High Salary people(>6k): 4

	younger_pay := gSum(employees, func(e Employee) float32 {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("Total Salary of Young People: %0.2f\n", younger_pay)
	//Total Salary of Young People: 19000.00

	total_vacation := gSum(employees, func(e Employee) int {
		return e.Vacation
	})
	fmt.Printf("Total Vacation: %d day(s)\n", total_vacation)
	//Total Vacation: 32 day(s)

	no_vacation := gFilterIn(employees, func(e Employee) bool {
		return e.Vacation == 0
	})
	print(no_vacation)
	//{Hao 44 0 8000.5} {Jack 26 0 4000} {Marry 29 0 6000}
}
