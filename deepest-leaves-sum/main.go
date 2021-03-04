package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var sum int
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		sum = 0
		l := len(queue)
		for i := 0; i < l; i++ {
			p := queue[0]
			queue = queue[1:]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			sum += p.Val
		}
	}
	return sum
}
