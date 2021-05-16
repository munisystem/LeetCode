package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := []string{}
	ans = append(ans, binaryTreePaths(root.Left)...)
	ans = append(ans, binaryTreePaths(root.Right)...)
	for i := 0; i < len(ans); i++ {
		ans[i] = fmt.Sprintf("%d->%s", root.Val, ans[i])
	}
	if len(ans) == 0 {
		ans = append(ans, fmt.Sprint(root.Val))
	}
	return ans
}
