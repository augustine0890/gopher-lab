package main

import (
	"testing"
)

// Given prices and maiximize the profit: O(n)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   4,
		},
		{
			prices: []int{},
			want:   0,
		},
	}

	for _, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.want {
			t.Errorf("maxProfit(%v) = %d; want %d", tc.prices, got, tc.want)
		}
	}
}
