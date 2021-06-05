package main

import (
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := int(math.Abs(float64(nums[i]))) - 1
		if nums[n] > 0 {
			nums[n] = -nums[n]
		}
	}
	var ans []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
