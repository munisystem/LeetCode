package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l, r := dfs(root.Left, res), dfs(root.Right, res)
	*res = int(math.Max(float64(*res), float64(l+r)))
	return int(math.Max(float64(l), float64(r))) + 1
}
