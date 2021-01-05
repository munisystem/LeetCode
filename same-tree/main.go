package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, p, q)
	for len(queue) != 0 {
		t1, t2 := queue[0], queue[1]
		queue = queue[2:]
		if t1 == nil && t2 == nil {
			continue
		} else if t1 == nil || t2 == nil {
			return false
		} else if t1.Val != t2.Val {
			return false
		}
		queue = append(queue, t1.Left, t2.Left, t1.Right, t2.Right)
	}
	return true
}
