package main

import "fmt"

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
func main() {
	nums := []int{1, 2, 3, 4}
	even := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	odd := Filter(nums, func(n int) bool {
		return n%2 != 0
	})
	fmt.Println(even) // [2 4]
	fmt.Println(odd)

}
