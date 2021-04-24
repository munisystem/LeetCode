package main

func shuffle(nums []int, n int) []int {
	r := make([]int, len(nums), len(nums))
	for i := 0; i < n; i++ {
		r[i*2], r[i*2+1] = nums[i], nums[i+n]
	}
	return r
}
