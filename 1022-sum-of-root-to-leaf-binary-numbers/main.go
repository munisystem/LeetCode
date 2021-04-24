package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return val<<1 + root.Val
	}
	return dfs(root.Left, val<<1+root.Val) + dfs(root.Right, val<<1+root.Val)
}
