package main

import "math"

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		h := int(math.Min(float64(height[i]), float64(height[j])))
		ans = int(math.Max(float64(ans), float64(h*(j-i))))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return ans
}
