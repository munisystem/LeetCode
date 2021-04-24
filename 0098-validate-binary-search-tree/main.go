package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	b, _, _ := validate(root)
	return b
}

func validate(root *TreeNode) (bool, int, int) {
	var b bool
	ll, lr, rl, rr := root.Val, root.Val, root.Val, root.Val
	if root.Left != nil {
		b, ll, lr = validate(root.Left)
		if !b || ll >= root.Val {
			return false, -1, -1
		}
	}
	if root.Right != nil {
		b, rl, rr = validate(root.Right)
		if !b || rr <= root.Val {
			return false, -1, -1
		}
	}
	return true, rl, lr
}
