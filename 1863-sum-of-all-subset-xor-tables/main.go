package main

func subsetXORSum(nums []int) int {
	var ans int
	for _, v := range nums {
		ans |= v
	}
	return ans * (1 << (len(nums) - 1))
}
