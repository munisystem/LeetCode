package main

import "math"

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2, n+2)
	dp[0], dp[1] = 0, 0
	for i := 2; i < n+2; i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+nums[i-2]), float64(dp[i-1])))
	}
	return dp[n+1]
}
