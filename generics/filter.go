package main

func gFilter[T any](arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}

func gFilterIn[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, true, f)
}

func gFilterOut[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, false, f)
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	odds := gFilterIn(nums, func(elem int) bool {
		return elem%2 == 1
	})
	print(odds)
}
