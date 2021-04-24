package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	v := -1
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		if t == nil {
			continue
		}
		if v == -1 {
			v = t.Val
		}
		if t.Val != v {
			return false
		}
		queue = append(queue, t.Left)
		queue = append(queue, t.Right)
	}
	return true
}
