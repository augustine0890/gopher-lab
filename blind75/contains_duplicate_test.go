package main

import "testing"

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct. O(n)
func containsDuplicate(nums []int) bool {
	encountered := make(map[int]bool)
	for _, num := range nums {
		if encountered[num] {
			return true
		}
		encountered[num] = true
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			nums: []int{1, 2, 3, 2},
			want: true,
		},
		{
			nums: []int{0},
			want: false,
		},
		{
			nums: []int{},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 1, 1},
			want: true,
		},
	}

	for _, tc := range testCases {
		got := containsDuplicate(tc.nums)
		if got != tc.want {
			t.Errorf("containsDuplicate(%v) = %t; want %t", tc.nums, got, tc.want)
		}
	}
}
