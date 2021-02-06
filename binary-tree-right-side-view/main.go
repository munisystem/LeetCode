package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	ans := []int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			t := queue[0]
			queue = queue[1:]
			if i == 0 {
				ans = append(ans, t.Val)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
		}
	}
	return ans
}
