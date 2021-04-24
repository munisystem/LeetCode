package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	root := &ListNode{}
	hp := root
	p := head
	t := &ListNode{}
	tp := t
	for p != nil {
		if p.Val < x {
			hp.Next = p
			hp = hp.Next
			p = p.Next
			continue
		}
		tp.Next = p
		tp = tp.Next
		p = p.Next
	}
	hp.Next = t.Next
	tp.Next = nil
	return root.Next
}
