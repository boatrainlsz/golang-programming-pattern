package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   float32
}

var employees = []Employee{
	{"Hao", 44, 0, 8000.5},
	{"Bob", 34, 10, 5000.5},
	{"Alice", 23, 5, 9000.0},
	{"Jack", 26, 0, 4000.0},
	{"Tom", 48, 9, 7500.75},
	{"Marry", 29, 0, 6000.0},
	{"Mike", 32, 8, 4000.3},
}

func totalSalary() {
	total_pay := gReduce(employees, 0.0, func(result float32, e Employee) float32 {
		return result + e.Salary
	})
	fmt.Printf("Total Salary: %0.2f\n", total_pay) // Total Salary: 43502.05
}
