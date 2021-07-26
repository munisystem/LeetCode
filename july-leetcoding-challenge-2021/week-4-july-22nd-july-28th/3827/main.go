package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := (len(nums) - 1) / 2
	tree := &TreeNode{Val: nums[i]}
	tree.Left, tree.Right = sortedArrayToBST(nums[:i]), sortedArrayToBST(nums[i+1:])
	return tree
}
