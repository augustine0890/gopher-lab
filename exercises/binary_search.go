package main

import "fmt"

// Time complexity: O(log(n))
func binarySearch(data []int, value int) bool {
	var mid int
	low := 0
	high := len(data) - 1
	for low <= high {
		mid = (low + high) / 2
		if value == data[mid] {
			return true
		} else {
			if value < data[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("BinarySearch:", binarySearch(arr, 8))
	fmt.Println("BinarySearch:", binarySearch(arr, 3))
}
