package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(r, root.Val)
	r = append(r, preorderTraversal(root.Left)...)
	r = append(r, preorderTraversal(root.Right)...)
	return r
}
