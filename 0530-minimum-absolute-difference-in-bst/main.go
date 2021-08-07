package main

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	nums := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			continue
		}
		nums = append(nums, n.Val)
		queue = append(queue, n.Left, n.Right)
	}
	sort.Ints(nums)
	ans := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < 0 {
			diff *= -1
		}
		if ans > diff {
			ans = diff
		}
	}
	return ans
}
