package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := []float64{}
	for len(queue) != 0 {
		l := len(queue)
		sum := 0
		for i := 0; i < l; i++ {
			tree := queue[0]
			queue = queue[1:]
			sum += tree.Val
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(l))
	}
	return ans
}
