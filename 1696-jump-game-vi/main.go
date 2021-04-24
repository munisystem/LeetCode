package main

import (
	"math"
)

func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}
	dp[0] = nums[0]
	m := 0
	for i := 1; i < len(dp); i++ {
		if m >= i-k {
			dp[i] = dp[m] + nums[i]
			if dp[i] >= dp[m] {
				m = i
			}
			continue
		}
		m = i
		for j := 1; j <= k && i-j >= 0; j++ {
			dp[i] = int(math.Max(float64(dp[i]), float64(dp[i-j]+nums[i])))
			if dp[m] < dp[i-j] {
				m = i - j
			}
		}
	}
	return dp[len(dp)-1]
}
