package main

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i > 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j])))
			}
			if j > 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j-1])))
			}
			dp[i][j] += grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
