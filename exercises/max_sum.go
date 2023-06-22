package main

import "fmt"

// Time complexity: O(n)
func maxSubArraySum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	currentMax, maxSum := data[0], data[0]
	for _, num := range data[1:] {
		if currentMax+num < num {
			currentMax = num
		} else {
			currentMax += num
		}

		if currentMax > maxSum {
			maxSum = currentMax
		}
	}
	return maxSum
}

// Given an array of positive and negative integers, find a contiguous subarray whose sum (sum of its elements) is the maximum.
func main() {
	data := []int{1, -2, 4, 4, -4, -6, -14, 1, 3}
	fmt.Println("Max sub array sum :", maxSubArraySum(data))
	fmt.Println(maxSubArraySum([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // Expected output: 7
	fmt.Println(maxSubArraySum([]int{1, 2, 3, 4, 5}))               // Expected output: 15
	fmt.Println(maxSubArraySum([]int{-1, -2, -3, -4, -5}))          // Expected output: -1
	fmt.Println(maxSubArraySum([]int{3, -2, 5, -1}))                // Expected output: 6
	fmt.Println(maxSubArraySum([]int{}))                            // Expected output: 0
}
