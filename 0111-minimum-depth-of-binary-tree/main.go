package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	n := 1

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			tree := queue[0]
			queue = queue[1:]
			if tree.Left == nil && tree.Right == nil {
				return n
			}
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
		}
		n++
	}
	return n
}
