package main

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(index); i++ {
		for j := 0; j < i; j++ {
			if index[i] <= index[j] {
				index[j]++
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		ans[index[i]] = nums[i]
	}
	return ans
}
