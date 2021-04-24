package main

import "fmt"

func main() {
	fmt.Println(combine(10, 10))
}

func combine(n int, k int) [][]int {
	return backtrack(n, k, 1, []int{})
}

func backtrack(n, k, index int, nums []int) [][]int {
	if k == 0 {
		return [][]int{append([]int{}, nums...)}
	}
	r := [][]int{}
	for i := index; i <= n; i++ {
		r = append(r, backtrack(n, k-1, i+1, append(append([]int{}, nums...), i))...)
	}
	return r
}
