package main

import (
	"math"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums), len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := len(nums), 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] {
			l = int(math.Min(float64(l), float64(i)))
			r = int(math.Max(float64(r), float64(i)))
		}
	}
	if r-l >= 0 {
		return r - l + 1
	}
	return 0
}
