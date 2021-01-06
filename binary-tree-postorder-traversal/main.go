package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(r, postorderTraversal(root.Left)...)
	r = append(r, postorderTraversal(root.Right)...)
	r = append(r, root.Val)
	return r
}
