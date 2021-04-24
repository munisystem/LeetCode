package main

import "math"

func largestAltitude(gain []int) int {
	ans := 0
	last := 0
	for i := 0; i < len(gain); i++ {
		last += gain[i]
		ans = int(math.Max(float64(ans), float64(last)))
	}
	return ans
}
