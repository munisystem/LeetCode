package main

func intersect(nums1 []int, nums2 []int) []int {
	nums := make([]int, 1001, 1001)
	ans := []int{}
	for _, num := range nums1 {
		nums[num]++
	}
	for _, num := range nums2 {
		if nums[num] > 0 {
			nums[num]--
			ans = append(ans, num)
		}
	}
	return ans
}
