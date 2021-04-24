package main

func minOperations(nums []int) int {
	ans := 0
	for i := 1; i < len(nums); i++ {
		sub := nums[i-1] + 1 - nums[i]
		if sub > 0 {
			nums[i] += sub
			ans += sub
		}
	}
	return ans
}
