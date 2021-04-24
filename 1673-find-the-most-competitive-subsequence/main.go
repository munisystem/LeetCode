package main

func mostCompetitive(nums []int, k int) []int {
	ans := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		for len(ans) > 0 && nums[i] < ans[len(ans)-1] && len(ans)+len(nums)-i > k {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, nums[i])
		}
	}
	return ans
}
