package main

import (
	"fmt"
	"testing"
)

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all
// the elements of nums except nums[i]. O(n)
/**
Compute the products of all elements to the left and all elements to the right of each index index in the  input array.
The space complexity is O(1). The answer array is not considered when calculating the extra space complexity
**/
func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	answer := make([]int, length)
	// answer[i] contains the product of all the numbers to the left of i.
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	fmt.Println("answer: ", answer)

	// Now for each i, compute product of all numbers to the right and multiply it with answer[i].
	right := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}

func arrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			nums: []int{2, 2, 2, 2},
			want: []int{8, 8, 8, 8},
		},
		{
			nums: []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			nums: []int{1, 0},
			want: []int{0, 1},
		},
		{
			nums: []int{},
			want: []int{},
		},
	}

	for _, tc := range testCases {
		got := productExceptSelf(tc.nums)
		if !arrayEqual(got, tc.want) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", tc.nums, got, tc.want)
		}
	}
}
