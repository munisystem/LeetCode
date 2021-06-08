package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	var ans int
	for len(queue) != 0 {
		s := len(queue)
		for i := 0; i < s; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			if i%2 == 0 && node.Left == nil && node.Right == nil {
				ans += node.Val
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return ans
}
