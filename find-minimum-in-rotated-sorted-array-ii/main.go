package main

import "math"

func findMin(nums []int) int {
	return bs(nums, 0, len(nums)-1)
}

func bs(nums []int, start, end int) int {
	if start >= end {
		return nums[start]
	}
	m := (start + end) / 2
	p := nums[m]
	if p > nums[end] {
		return bs(nums, m+1, end)
	} else {
		return int(math.Min(float64(bs(nums, start, m)), float64(bs(nums, m+1, end))))
	}
}
