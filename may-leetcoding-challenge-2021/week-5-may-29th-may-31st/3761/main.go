package main

import "sort"

func maximumGap(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 1; i < len(nums); i++ {
		if ans < nums[i]-nums[i-1] {
			ans = nums[i] - nums[i-1]
		}
	}
	return ans
}
