package main

import "math"

func findPoisonedDuration(timeSeries []int, duration int) int {
	var ans, timer int
	for i := 0; i < len(timeSeries); i++ {
		if i > 0 {
			ans += int(math.Min(float64(timeSeries[i]-timeSeries[i-1]), float64(timer)))
		}
		timer = duration
	}
	return ans + timer
}
