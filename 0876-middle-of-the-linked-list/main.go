package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p, f := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		p = p.Next
	}
	return p
}
