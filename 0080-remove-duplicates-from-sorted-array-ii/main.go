package main

func removeDuplicates(nums []int) int {
	lastNum, count, p := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			lastNum = nums[i]
			count++
			p++
			continue
		}
		if nums[i] == lastNum {
			if count == 2 {
				continue
			} else {
				count++
				nums[p] = nums[i]
				p++
			}
		} else {
			lastNum, nums[p] = nums[i], nums[i]
			count = 1
			p++
		}
	}
	return p
}
