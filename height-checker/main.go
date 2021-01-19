package main

import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights), len(heights))
	for i := 0; i < len(heights); i++ {
		sorted[i] = heights[i]
	}
	sort.Ints(sorted)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if sorted[i] != heights[i] {
			ans++
		}
	}
	return ans
}
