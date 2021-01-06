package main

import (
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			profit := (prices[i-1] - prices[i-j-1]) + dp[i-j-1]
			dp[i] = int(math.Max(float64(dp[i]), float64(profit)))

		}
	}
	return dp[n]
}
