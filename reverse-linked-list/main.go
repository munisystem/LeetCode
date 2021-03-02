package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
		t := head
		head = head.Next
		t.Next = p
		p = t
	}
	return p
}
