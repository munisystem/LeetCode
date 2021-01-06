package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	r := [][]int{}
	toRight := false
	for len(queue) != 0 {
		l := len(queue)
		rr := []int{}
		for i := 0; i < l; i++ {
			t := queue[0]
			queue = queue[1:]
			if t == nil {
				continue
			}
			queue = append(queue, t.Left)
			queue = append(queue, t.Right)
			rr = append(rr, t.Val)
		}
		if toRight {
			t := []int{}
			for i := len(rr) - 1; i >= 0; i-- {
				t = append(t, rr[i])
			}
			rr = t
		}
		if len(rr) != 0 {
			r = append(r, rr)
		}
		toRight = !toRight
	}
	return r
}
