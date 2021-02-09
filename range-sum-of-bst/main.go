package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		if t == nil {
			continue
		}
		if t.Val < low {
			queue = append(queue, t.Right)
		} else if t.Val > high {
			queue = append(queue, t.Left)
		} else {
			ans += t.Val
			queue = append(queue, t.Left)
			queue = append(queue, t.Right)
		}
	}
	return ans
}
