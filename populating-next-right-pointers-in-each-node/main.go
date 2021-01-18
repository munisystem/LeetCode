package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		var prev, next *Node
		for i := 0; i < l; i++ {
			next = queue[0]
			queue = queue[1:]
			if next.Left != nil {
				queue = append(queue, next.Left)
			}
			if next.Right != nil {
				queue = append(queue, next.Right)
			}
			if prev != nil {
				prev.Next = next
			}
			prev = next
		}
	}
	return root
}
