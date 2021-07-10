package main

import "math"

func shortestToChar(S string, C byte) []int {
	pos := -len(S)
	ans := make([]int, len(S), len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pos = i
		}
		ans[i] = i - pos
	}
	for i := pos - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		ans[i] = int(math.Min(float64(ans[i]), float64(pos-i)))
	}
	return ans
}
