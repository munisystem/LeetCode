package main

func findMaxAverage(nums []int, k int) float64 {
	var prev, max int
	for i := 0; i < len(nums); i++ {
		prev += nums[i]
		if i < k {
			max = prev
			continue
		}
		prev -= nums[i-k]
		if prev > max {
			max = prev
		}
	}
	return float64(max) / float64(k)
}
