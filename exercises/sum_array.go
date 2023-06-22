package main

import "fmt"

// The time complexity of the solution is O(n)
func sumArray(data []int) int {
	total := 0

	for _, n := range data {
		total += n
	}
	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Sum of all the values in array:", sumArray(data))
}
