package main

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	g := math.MinInt32
	for i := 0; i < len(candies); i++ {
		g = int(math.Max(float64(g), float64(candies[i])))
	}
	b := make([]bool, len(candies), len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= g {
			b[i] = true
		} else {
			b[i] = false
		}
	}
	return b
}
