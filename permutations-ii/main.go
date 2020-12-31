package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{-1, -1, -1, 1, 1, 2, 2, 2}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return dfs([]int{}, nums)
}

func dfs(order, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{append([]int{}, order...)}
	}
	m := map[int]interface{}{}
	r := [][]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		n := []int{}
		if i > 0 {
			n = append(n, nums[:i]...)
		}
		if i < len(nums)-1 {
			n = append(n, nums[i+1:]...)
		}
		a := dfs(append(order, nums[i]), n)
		r = append(r, a...)
		m[nums[i]] = nil
	}
	return r
}
