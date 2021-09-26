package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	prev = head
	traversal(root)
	return head.Right
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	traversal(root.Left)
	prev.Right, prev = root, root
	root.Left = nil
	traversal(root.Right)
}
