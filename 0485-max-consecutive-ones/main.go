package main

func findMaxConsecutiveOnes(nums []int) int {
	max, counter := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter = counter + 1
		} else {
			if counter > max {
				max = counter
			}
			counter = 0
		}
	}
	if counter > max {
		max = counter
	}
	return max
}
