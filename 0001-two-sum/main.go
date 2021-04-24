package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
