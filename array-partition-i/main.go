package main

import (
	"math"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)/2; i++ {
		ans += int(math.Min(float64(nums[i*2]), float64(nums[i*2+1])))
	}
	return ans
}
