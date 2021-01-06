package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	return dfs(root, sum)
}

func dfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return true
		} else {
			return false
		}
	}
	return dfs(root.Left, sum) || dfs(root.Right, sum)
}
