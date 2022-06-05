package array

import "math"

func maxProfit(prices []int) int {
	profit, min := 0, math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			profit = max(profit, prices[i]-min)
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
