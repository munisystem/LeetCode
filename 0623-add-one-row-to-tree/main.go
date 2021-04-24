package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for depth := 1; depth < d-1; depth++ {
		l := len(queue)
		for i := 0; i < l; i++ {
			tree := queue[0]
			queue = queue[1:]
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
		}
	}
	for i := 0; i < len(queue); i++ {
		tree := queue[i]
		l, r := tree.Left, tree.Right
		tree.Left, tree.Right = &TreeNode{Val: v, Left: l}, &TreeNode{Val: v, Right: r}
	}
	return root
}
