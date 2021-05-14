package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if left != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		left.Right = tmp
		if right == nil {
			return left
		}
	}
	return right
}
