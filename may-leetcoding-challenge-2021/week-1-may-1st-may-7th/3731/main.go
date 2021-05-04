package main

func checkPossibility(nums []int) bool {
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if n++; n > 1 {
				return false
			}
			if i-2 < 0 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return n <= 1
}
