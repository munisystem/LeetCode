package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
