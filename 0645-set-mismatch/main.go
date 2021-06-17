package main

func findErrorNums(nums []int) []int {
	n := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		n[nums[i]-1]++
	}
	ans := make([]int, 2, 2)
	for i := 0; i < len(n); i++ {
		switch n[i] {
		case 2:
			ans[0] = i + 1
		case 0:
			ans[1] = i + 1
		}
	}
	return ans
}
