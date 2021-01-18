package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return bst(head)
}

func bst(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	p, mid := head, head
	var prev *ListNode
	for p != nil && p.Next != nil {
		prev = mid
		mid = mid.Next
		p = p.Next.Next
	}
	root := &TreeNode{Val: mid.Val}
	root.Right = bst(mid.Next)
	if prev != nil {
		prev.Next = nil
		root.Left = bst(head)
	}
	return root
}
