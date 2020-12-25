package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{}
	root.Next = head
	p := root
	for p.Next != nil && p.Next.Next != nil {
		nnn := p.Next.Next.Next
		n := p.Next
		p.Next = p.Next.Next
		p.Next.Next = n
		p.Next.Next.Next = nnn
		p = p.Next.Next
	}
	return root.Next
}
