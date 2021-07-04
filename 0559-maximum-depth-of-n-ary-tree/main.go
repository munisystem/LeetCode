package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var ans int
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) != 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			for j := 0; j < len(n.Children); j++ {
				if n.Children[j] != nil {
					queue = append(queue, n.Children[j])
				}
			}
		}
	}
	return ans
}
