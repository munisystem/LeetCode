package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var mi, smi int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[smi] || i <= 1 {
			smi = i
		}
		if nums[smi] > nums[mi] {
			smi, mi = mi, smi
		}
	}
	if nums[mi] >= nums[smi]*2 {
		return mi
	}
	return -1
}
