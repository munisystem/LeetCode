package main

func intersection(nums1 []int, nums2 []int) []int {
	nums := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		if nums[nums1[i]] == 0 {
			nums[nums1[i]]++
		}
	}
	var ans []int
	for i := 0; i < len(nums2); i++ {
		if nums[nums2[i]] == 1 {
			ans = append(ans, nums2[i])
			nums[nums2[i]]++
		}
	}
	return ans
}
