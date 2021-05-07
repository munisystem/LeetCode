package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := m[nums[i]]; ok {
			if i-j <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}
