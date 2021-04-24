package main

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	prev := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if prev != -1 && i-prev-k <= 0 {
				return false
			}
			prev = i
		}
	}
	return true
}
