package main

func buildArray(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
