package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l, r := dfs(root1), dfs(root2)
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

func dfs(tree *TreeNode) []int {
	if tree == nil {
		return []int{}
	}
	if tree.Left == nil && tree.Right == nil {
		return []int{tree.Val}
	}
	leaves := []int{}
	if tree.Left != nil {
		leaves = append(leaves, dfs(tree.Left)...)
	}
	if tree.Right != nil {
		leaves = append(leaves, dfs(tree.Right)...)
	}
	return leaves
}
