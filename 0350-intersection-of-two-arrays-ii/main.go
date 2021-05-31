package main

func intersect(nums1 []int, nums2 []int) []int {
	nums := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		nums[nums1[i]]++
	}
	var ans []int
	for i := 0; i < len(nums2); i++ {
		nums[nums2[i]]--
		if nums[nums2[i]] >= 0 {
			ans = append(ans, nums2[i])
		}
	}
	return ans
}
