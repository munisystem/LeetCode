package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(r, inorderTraversal(root.Left)...)
	r = append(r, root.Val)
	r = append(r, inorderTraversal(root.Right)...)
	return r
}
