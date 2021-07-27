package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	ans := make([]string, len(score), len(score))
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Ints(sorted)
	places := map[int]int{}
	for i := 0; i < len(sorted); i++ {
		places[sorted[i]] = len(sorted) - 1 - i
	}
	for i := 0; i < len(score); i++ {
		p := places[score[i]] + 1
		switch p {
		case 1:
			ans[i] = "Gold Medal"
		case 2:
			ans[i] = "Silver Medal"
		case 3:
			ans[i] = "Bronze Medal"
		default:
			ans[i] = fmt.Sprintf("%d", p)
		}
	}
	return ans
}
