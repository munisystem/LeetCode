package main

import (
	"math"
)

func maxAscendingSum(nums []int) int {
	var ans int
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans = int(math.Max(float64(ans), float64(sum)))
			sum = 0
		}
		sum += nums[i]
	}
	return int(math.Max(float64(ans), float64(sum)))
}
