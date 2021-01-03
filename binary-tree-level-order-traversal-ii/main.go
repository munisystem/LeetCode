package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	r := make([][]int, 0, len(res))
	for i := len(res) - 1; i >= 0; i-- {
		r = append(r, res[i])
	}
	return r
}

func dfs(tree *TreeNode, level int, res *[][]int) {
	if tree == nil {
		return
	}
	if len(*res) < level+1 {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], tree.Val)
	dfs(tree.Left, level+1, res)
	dfs(tree.Right, level+1, res)
	return
}
