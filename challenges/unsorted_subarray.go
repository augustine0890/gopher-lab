package main

import (
	"fmt"
)

// Time: O(N) and Space: O(1)
func findUnsortedSubarray(nums []int) int {
	left, right, max := 1, 1, nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > max {
			max = nums[i]
		}

		// increasing until num[:left] is correctly sorted
		if left == i && nums[i] > nums[i-1] {
			left++
			right++
		} else if nums[i] < max {
			// subsequent element should be equal to or greater the current maximum so far,
			// if not the marked as unsorted elements
			right = i + 1
		}

		// check current element fits into the sorted portion num[:left]
		for left > 0 && nums[left-1] > nums[i] {
			left--
		}
	}

	return right - left

}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))

}
