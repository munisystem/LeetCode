package main

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		m := (start + end) / 2
		pivot := nums[m]
		if pivot > nums[end] {
			start = m + 1
		} else {
			end = m
		}
	}
	return nums[start]
}
