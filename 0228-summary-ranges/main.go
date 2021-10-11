package main

import "fmt"

func summaryRanges(nums []int) []string {
	var ans []string
	if len(nums) == 0 {
		return ans
	}
	var left int
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[left]-(i-left) == 0 {
			continue
		}
		if i-left == 1 {
			ans = append(ans, fmt.Sprintf("%d", nums[left]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[left], nums[i-1]))
		}
		left = i
	}
	if left == len(nums)-1 {
		ans = append(ans, fmt.Sprintf("%d", nums[left]))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[left], nums[len(nums)-1]))
	}
	return ans
}
