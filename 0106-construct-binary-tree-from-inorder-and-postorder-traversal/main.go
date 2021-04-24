package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	i := len(inorder) - 1
	for ; inorder[i] != rootVal; i-- {
	}
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	root.Left = buildTree(inorder[:i], postorder[:i])
	return root
}
