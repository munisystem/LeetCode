package main

func arraySign(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			cnt++
		}
	}
	if cnt%2 == 0 {
		return 1
	} else {
		return -1
	}
}
