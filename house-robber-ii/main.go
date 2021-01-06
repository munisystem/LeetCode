package main

import (
	"math"
)

func rob(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n+1, n+1)
	dp[1][0] = nums[0]
	for i := 2; i < n+1; i++ {
		for j := 0; j <= 1; j++ {
			if i == n && j == 0 {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-2][j])))
				continue
			}
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-2][j]+nums[i-1])))
		}
	}
	return int(math.Max(float64(dp[n][0]), float64(dp[n][1])))
}
