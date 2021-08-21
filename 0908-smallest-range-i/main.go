package main

import (
	"math"
)

func smallestRangeI(nums []int, k int) int {
	n, m := nums[0], nums[0]
	for _, num := range nums {
		if n > num {
			n = num
		}
		if m < num {
			m = num
		}
	}
	return int(math.Max(0, float64(m-n-k*2)))
}
