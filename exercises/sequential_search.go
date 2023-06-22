package main

import "fmt"

// The time complexity is O(n)
func sequentialSearch(data []int, value int) bool {
	for _, n := range data {
		if value == n {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("SequentialSearch:", sequentialSearch(arr, 8))
	fmt.Println("SequentialSearch:", sequentialSearch(arr, 9))
}
