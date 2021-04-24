package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := [][]int{}
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) != 0 {
		r := []int{}
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
			r = append(r, t.Val)
		}
		res = append(res, r)
	}
	return res
}
