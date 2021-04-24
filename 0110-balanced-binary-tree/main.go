package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := backtrack(root)
	return b
}

func backtrack(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 1, true
	}
	ld, lb := backtrack(root.Left)
	rd, rb := backtrack(root.Right)
	if !(lb && rb) {
		return 0, false
	}
	if int(math.Abs(float64(ld-rd))) <= 1 {
		return int(math.Max(float64(ld), float64(rd))) + 1, true
	}
	return 0, false
}
