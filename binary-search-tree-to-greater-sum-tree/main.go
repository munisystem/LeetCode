package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	tree, _ := gst(root, 0)
	return tree
}

func gst(root *TreeNode, base int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	var val int
	if root.Right != nil {
		root.Right, val = gst(root.Right, base)
		root.Val += int(math.Max(float64(root.Right.Val), float64(val)))
	} else {
		root.Val += base
	}
	root.Left, val = gst(root.Left, root.Val)
	return root, int(math.Max(float64(root.Val), float64(val)))
}
