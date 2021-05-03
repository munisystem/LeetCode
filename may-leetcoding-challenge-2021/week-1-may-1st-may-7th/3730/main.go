package main

func runningSum(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = nums[i]
		} else {
			ans[i] = nums[i] + ans[i-1]
		}
	}
	return ans
}
