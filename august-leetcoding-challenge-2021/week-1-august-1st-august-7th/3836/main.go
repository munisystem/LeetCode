package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	ans := make([]int, 2, 2)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok && i != v {
			ans[0], ans[1] = i, v
			break
		}
	}
	return ans
}
