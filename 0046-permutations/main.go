package main

func permute(nums []int) [][]int {
	return dfs([]int{}, nums)
}

func dfs(order, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{order}
	}
	r := [][]int{}
	for i := 0; i < len(nums); i++ {
		n := []int{}
		if i > 0 {
			n = append(n, nums[:i]...)
		}
		if i < len(nums)-1 {
			n = append(n, nums[i+1:]...)
		}
		r = append(r, dfs(append(order, nums[i]), n)...)
	}
	return r
}
