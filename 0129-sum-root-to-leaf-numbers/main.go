package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}
	return sum(root.Left, num) + sum(root.Right, num)
}
