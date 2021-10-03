package main

func countKDifference(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] - nums[j]
			if sum < 0 {
				sum *= -1
			}
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
