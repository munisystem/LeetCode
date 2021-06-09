package main

import "math"

func thirdMax(nums []int) int {
	f, s, t := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < t {
			continue
		}
		if nums[i] > f {
			t = s
			s = f
			f = nums[i]
		} else if nums[i] < f && nums[i] > s {
			t = s
			s = nums[i]
		} else if nums[i] < s && nums[i] > t {
			t = nums[i]
		}
	}
	if s == math.MinInt64 || t == math.MinInt64 {
		return f
	} else {
		return t
	}
}
