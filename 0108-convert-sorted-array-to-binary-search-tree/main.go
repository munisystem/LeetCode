package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return bst(nums)
}

func bst(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := (len(nums) - 1) / 2
	root := &TreeNode{Val: nums[i]}
	root.Left, root.Right = bst(nums[:i]), bst(nums[i+1:])
	return root
}
