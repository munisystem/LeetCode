package main

func sumOfUnique(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	ans := 0
	for k, v := range m {
		if v == 1 {
			ans += k
		}
	}
	return ans
}
