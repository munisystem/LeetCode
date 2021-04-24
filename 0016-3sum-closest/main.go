package main

import "math"

func threeSumClosest(nums []int, target int) int {
	min := math.MaxInt32
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := int(math.Abs(float64(target - sum)))
				if diff == 0 {
					return target
				}
				if diff < min {
					min = diff
					r = sum
				}
			}
		}
	}
	return r
}
