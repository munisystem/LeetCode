package main

import (
	"strconv"
)

func numDecodings(s string) int {
	dp := make([]int, len(s)+1, len(s)+1)
	dp[0] = 1
	for i := 1; i < len(s)+1; i++ {
		dp[i] = 0
		n, _ := strconv.Atoi(s[i-1 : i])
		if n != 0 {
			dp[i] = dp[i] + dp[i-1]
		}
		if i > 1 {
			nn, _ := strconv.Atoi(s[i-2 : i])
			if nn <= 26 && nn >= 10 {
				dp[i] = dp[i] + dp[i-2]
			}
		}
	}
	return dp[len(s)]
}
