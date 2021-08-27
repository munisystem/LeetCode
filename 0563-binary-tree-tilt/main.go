package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var ans int
	_ = tilt(root, &ans)
	return ans
}

func tilt(root *TreeNode, tSum *int) int {
	if root == nil {
		return 0
	}
	left, right := tilt(root.Left, tSum), tilt(root.Right, tSum)
	var val int
	if left > right {
		val = left - right
	} else {
		val = right - left
	}
	*tSum += val

	return root.Val + left + right
}
