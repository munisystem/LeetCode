package main

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (1 + n) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}
