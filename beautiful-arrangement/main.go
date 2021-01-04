package main

import "fmt"

func main() {
	fmt.Println(countArrangement(15))
}

func countArrangement(n int) int {
	nums := make([]int, n, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return backtrack(nums, 1)
}

func backtrack(nums []int, index int) int {
	if len(nums) == 0 {
		return 1
	}
	r := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num%index == 0 || index%num == 0 {
			n := append([]int{}, nums[0:i]...)
			n = append(n, nums[i+1:]...)
			r = r + backtrack(n, index+1)
		}
	}
	return r
}
