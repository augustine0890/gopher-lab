package main

import (
	"math"
	"testing"
)

// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted. O(n)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum // keep track of maxSum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i] // sliding window
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{0, 4, 0, 3, 2}, 1, 4.0},
		{[]int{-1, -2, -3, -4, -5}, 3, -2.0},
		{[]int{2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 9.5},
	}

	for _, tc := range testCases {
		result := findMaxAverage(tc.nums, tc.k)
		if math.Abs(result-tc.expected) > 1e-5 {
			t.Errorf("findMaxAverage(%v, %v) = %v; expected %v", tc.nums, tc.k, result, tc.expected)
		}
	}
}
