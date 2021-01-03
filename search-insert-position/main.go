package main

func searchInsert(nums []int, target int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
		r = i + 1
	}
	return r
}
