package main

import "math"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost), len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) + cost[i]
	}
	return int(math.Min(float64(dp[len(cost)-1]), float64(dp[len(cost)-2])))
}
