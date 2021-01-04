package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	r := [][]int{}
	r = append(r, []int{})
	for i := 0; i < len(nums); i++ {
		l := len(r)
		r = append(r, []int{nums[i]})
		for j := 1; j < l; j++ {
			r = append(r, append(append([]int{}, r[j]...), nums[i]))
		}
	}
	return r
}
