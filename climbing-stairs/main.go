package main

func climbStairs(n int) int {
	dp := make([]int, n+2, n+2)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n+1]
}
