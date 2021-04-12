package main

import "math"

func findShortestSubArray(nums []int) int {
	m, l, r := map[int]int{}, map[int]int{}, map[int]int{}
	degree := 0
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		r[nums[i]] = i
		if m[nums[i]] == 1 {
			l[nums[i]] = i
		}
		if m[nums[i]] > degree {
			degree = m[nums[i]]
		}
	}
	ans := len(nums)
	for k, v := range m {
		if v == degree {
			ans = int(math.Min(float64(ans), float64(r[k]-l[k]+1)))
		}
	}
	return ans
}
