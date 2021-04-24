package main

import "math"

func minOperations(boxes string) []int {
	ans := make([]int, len(boxes), len(boxes))
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '0' {
			continue
		}
		for j := 0; j < len(ans); j++ {
			ans[j] += int(math.Abs(float64(j - i)))
		}
	}
	return ans
}
