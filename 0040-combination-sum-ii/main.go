package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return backtrack(candidates, target, []int{}, 0)
}

func backtrack(candidates []int, target int, current []int, index int) [][]int {
	if target == 0 {
		return [][]int{append([]int{}, current...)}
	}
	if target < 0 {
		return [][]int{}
	}
	res := [][]int{}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		res = append(res, backtrack(candidates, target-candidates[i], append(current, candidates[i]), i+1)...)
	}
	return res
}
