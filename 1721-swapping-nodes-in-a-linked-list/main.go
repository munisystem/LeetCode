package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	p := head
	size := 0
	for p != nil {
		size++
		p = p.Next
	}
	l, r := head, head
	for i := 0; i < k-1; i++ {
		l = l.Next
	}
	for i := 0; i < size-k; i++ {
		r = r.Next
	}
	tmp := r.Val
	r.Val = l.Val
	l.Val = tmp
	return head
}
