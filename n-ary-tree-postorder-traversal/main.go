package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	ans := []int{}
	for i := 0; i < len(root.Children); i++ {
		ans = append(ans, postorder(root.Children[i])...)
	}
	ans = append(ans, root.Val)
	return ans
}
