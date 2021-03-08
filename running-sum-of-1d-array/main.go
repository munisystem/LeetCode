package main

func runningSum(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		if i > 0 {
			ans[i] += ans[i-1]
		}
	}
	return ans
}
