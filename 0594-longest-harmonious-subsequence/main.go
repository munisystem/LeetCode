package main

import (
	"math"
)

func findLHS(nums []int) int {
	m := map[int]int{}
	max := 0
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if v, ok := m[nums[i]+1]; ok {
			max = int(math.Max(float64(max), float64(m[nums[i]]+v)))
		}
		if v, ok := m[nums[i]-1]; ok {
			max = int(math.Max(float64(max), float64(m[nums[i]]+v)))
		}
	}
	return max
}
