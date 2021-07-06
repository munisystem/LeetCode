package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	ans := make([]int, len(nums1), len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := m[nums1[i]]; ok {
			ans[i] = v
		} else {
			ans[i] = -1
		}
	}
	return ans
}
