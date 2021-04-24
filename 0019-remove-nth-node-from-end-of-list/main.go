package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	root := &ListNode{}
	root.Next = head
	p := root
	for i := 0; i < l-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return root.Next
}
