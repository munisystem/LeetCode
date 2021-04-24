package main

import "math"

func countGoodRectangles(rectangles [][]int) int {
	maxLen, count := 0, 0
	for i := 0; i < len(rectangles); i++ {
		min := int(math.Min(float64(rectangles[i][0]), float64(rectangles[i][1])))
		if min == maxLen {
			count++
		} else if min > maxLen {
			maxLen = min
			count = 1
		}
	}
	return count
}
