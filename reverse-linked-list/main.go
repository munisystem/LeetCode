package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	t := head.Next
	p.Next = nil
	for t != nil {
		tt := t.Next
		t.Next = p
		p = t
		t := tt
	}
	return p
}
