package main

import "testing"

// Given an array of integers, return the indices of the two numbers that add up to a given target.
func twoSum(nums []int, target int) (int, int) {
	mapIndex := make(map[int]int)
	for i, num := range nums {
		subNum := target - num
		if j, ok := mapIndex[subNum]; ok {
			return i, j
		} else {
			mapIndex[num] = i
		}
	}
	return -1, -1
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	i, j := twoSum(nums, target)
	if !(i == 1 && j == 0) {
		t.Errorf("Expected indices 1, 0, but got %d, %d", i, j)
	}

	nums = []int{3, 2, 4}
	target = 6
	i, j = twoSum(nums, target)
	if !(i == 2 && j == 1) {
		t.Errorf("Expected indices 2, 1, but got %d, %d", i, j)
	}
}
