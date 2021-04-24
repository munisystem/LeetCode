package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, root)

	for len(queue) != 0 {
		l := queue[0]
		r := queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}
