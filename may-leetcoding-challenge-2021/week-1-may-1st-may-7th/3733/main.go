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
	if head == nil {
		return nil
	}

	prev, mid, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = mid
		mid = mid.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: mid.Val}
	root.Right = sortedListToBST(mid.Next)
	if prev != mid {
		prev.Next = nil
		root.Left = sortedListToBST(head)
	}
	return root
}
