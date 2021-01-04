package main

import (
	"math"
)

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	r := 1
	mm, nn := m-1, n-1
	k := int(math.Min(float64(mm), float64(nn)))
	for i := 0; i < k; i++ {
		r = r * (mm + nn - i)
		r = r / (i + 1)
	}
	return r
}
