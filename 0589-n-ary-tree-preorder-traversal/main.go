package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	ans := []int{}
	ans = append(ans, root.Val)
	for i := 0; i < len(root.Children); i++ {
		ans = append(ans, preorder(root.Children[i])...)
	}
	return ans
}
