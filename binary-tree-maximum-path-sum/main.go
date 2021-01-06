package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	_, max := dfs(root)
	return max
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return math.MinInt32, math.MinInt32
	}
	sum, max := root.Val, root.Val
	ls, lm := dfs(root.Left)
	rs, rm := dfs(root.Right)
	sum = int(math.Max(float64(sum), float64(root.Val+ls)))
	sum = int(math.Max(float64(sum), float64(root.Val+rs)))
	max = int(math.Max(float64(max), float64(lm)))
	max = int(math.Max(float64(max), float64(rm)))
	max = int(math.Max(float64(max), float64(root.Val+ls+rs)))
	max = int(math.Max(float64(max), float64(sum)))
	return sum, max
}
