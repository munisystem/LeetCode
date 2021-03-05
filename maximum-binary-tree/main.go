package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	tree := &TreeNode{}
	mi, mv := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mv {
			mv = nums[i]
			mi = i
		}
	}
	tree.Val = mv
	tree.Left = constructMaximumBinaryTree(nums[0:mi])
	tree.Right = constructMaximumBinaryTree(nums[mi+1:])
	return tree
}
