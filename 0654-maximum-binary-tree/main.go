package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mi, mv := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mv {
			mv = nums[i]
			mi = i
		}
	}
	return &TreeNode{
		Val:   nums[mi],
		Left:  constructMaximumBinaryTree(nums[0:mi]),
		Right: constructMaximumBinaryTree(nums[mi+1:]),
	}
}
