package main

import "math"

func maximumProduct(nums []int) int {
	h1, h2, h3 := math.MinInt32, math.MinInt32, math.MinInt32
	l1, l2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num < l1 {
			l1, l2 = num, l1
		} else if num < l2 {
			l2 = num
		}
		if num > h1 {
			h1, h2, h3 = num, h1, h2
		} else if num > h2 {
			h2, h3 = num, h2
		} else if num > h3 {
			h3 = num
		}
	}
	return int(math.Max(float64(l1*l2*h1), float64(h1*h2*h3)))
}
