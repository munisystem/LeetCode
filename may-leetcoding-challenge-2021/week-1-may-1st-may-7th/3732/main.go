package main

func jump(nums []int) int {
	ans, end, max := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
		if i == end {
			ans++
			end = max
		}
	}
	return ans
}
