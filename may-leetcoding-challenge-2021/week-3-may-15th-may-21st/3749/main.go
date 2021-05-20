package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		nums := make([]int, size, size)
		for i := 0; i < size; i++ {
			leaf := queue[0]
			queue = queue[1:]
			nums[i] = leaf.Val
			if leaf.Left != nil {
				queue = append(queue, leaf.Left)
			}
			if leaf.Right != nil {
				queue = append(queue, leaf.Right)
			}
		}
		ans = append(ans, nums)
	}
	return ans
}
