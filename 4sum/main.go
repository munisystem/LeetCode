package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return backtrack(nums, target, 4, []int{})
}

func backtrack(nums []int, target, n int, selected []int) [][]int {
	if n == 0 {
		if target == 0 {
			return [][]int{append([]int{}, selected...)}
		}
		return [][]int{}
	}
	r := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if target-nums[i] < 0 && nums[i] > 0 {
			break
		}
		r = append(r, backtrack(nums[i+1:], target-nums[i], n-1, append(append([]int{}, selected...), nums[i]))...)
	}
	return r
}
