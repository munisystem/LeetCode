package main

import "fmt"

func countAndSay(n int) string {
	dp := make([]string, n, n)
	dp[0] = "1"
	for i := 1; i < n; i++ {
		prev := dp[i-1]
		last := string(prev[0])
		count := 1
		s := ""
		for i := 1; i < len(prev); i++ {
			if string(prev[i]) == last {
				count++
				continue
			} else {
				s = s + fmt.Sprintf("%d%s", count, last)
				last = string(prev[i])
				count = 1
			}
		}
		s = s + fmt.Sprintf("%d%s", count, last)
		dp[i] = s
	}
	return dp[n-1]
}
