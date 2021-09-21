package main

func findMaxConsecutiveOnes(nums []int) int {
	var ans, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
			continue
		}
		count++
		if ans <= count {
			ans = count
		}
	}
	return ans
}
