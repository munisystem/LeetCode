package main

import "math"

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
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		s := []int{}
		for i := 0; i < l; i++ {
			t := queue[0]
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue, t.Left)
				s = append(s, t.Left.Val)
			} else {
				s = append(s, math.MaxInt32)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
				s = append(s, t.Right.Val)
			} else {
				s = append(s, math.MaxInt32)
			}
		}
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}
