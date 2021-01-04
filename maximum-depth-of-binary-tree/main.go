package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	n := 1
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			t := queue[0]
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		if len(queue) != 0 {
			n++
		}
	}
	return n
}
