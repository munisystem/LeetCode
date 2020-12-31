package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target, []int{})
}

func dfs(candidates []int, target int, nums []int) [][]int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if len(candidates) == 0 {
		if sum == target {
			return [][]int{append([]int{}, nums...)}
		}
		return [][]int{}
	}
	r := [][]int{}
	r = append(r, dfs(candidates[1:], target, append([]int{}, nums...))...)
	for i := 1; sum+candidates[0]*i <= target; i++ {
		nums = append(nums, candidates[0])
		r = append(r, dfs(candidates[1:], target, append([]int{}, nums...))...)
	}
	return r
}
