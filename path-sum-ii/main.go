package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	return dfs(root, sum, []int{})
}

func dfs(root *TreeNode, sum int, selected []int) [][]int {
	if root == nil {
		return [][]int{}
	}
	sum = sum - root.Val
	selected = append(selected, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return [][]int{append([]int{}, selected...)}
		} else {
			return [][]int{}
		}
	}
	r := [][]int{}
	r = append(r, dfs(root.Left, sum, append([]int{}, selected...))...)
	r = append(r, dfs(root.Right, sum, append([]int{}, selected...))...)
	return r
}
