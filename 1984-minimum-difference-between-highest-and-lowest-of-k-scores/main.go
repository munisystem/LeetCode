package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		if ans > nums[i+k-1]-nums[i] {
			ans = nums[i+k-1] - nums[i]
		}
	}
	return ans
}
