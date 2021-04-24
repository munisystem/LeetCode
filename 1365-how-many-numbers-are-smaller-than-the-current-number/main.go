package main

func smallerNumbersThanCurrent(nums []int) []int {
	cnt := make([]int, 101, 101)
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}
	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ans[i] = cnt[nums[i]-1]
		}
	}
	return ans
}
