package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		s, d := points[i-1], points[i]
		ans += int(math.Max(math.Abs(float64(d[0]-s[0])), math.Abs(float64(d[1]-s[1]))))
	}
	return ans
}
