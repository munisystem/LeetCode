package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		min := i
		for j := 1; j*j <= i; j++ {
			if num := 1 + dp[i-j*j]; min > num {
				min = num
			}
		}
		dp[i] = min
	}

	return dp[n]
}
