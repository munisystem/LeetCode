package main

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j > 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j-1])))
			}
			if j < len(triangle[i])-1 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j])))
			}
			dp[i][j] = dp[i][j] + triangle[i][j]
		}
	}
	r := math.MaxInt32
	for i := 0; i < n; i++ {
		r = int(math.Min(float64(r), float64(dp[n-1][i])))
	}
	return r
}
