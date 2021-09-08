package main

func findMiddleIndex(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v

	}
	var left int
	for i, v := range nums {
		sum -= v
		if left == sum {
			return i
		}
		left += v
	}
	return -1
}
