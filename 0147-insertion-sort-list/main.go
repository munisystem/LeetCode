package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	root := &ListNode{}
	root.Next = head
	p := root.Next
	for p != nil && p.Next != nil {
		if p.Val <= p.Next.Val {
			p = p.Next
			continue
		}
		tmp := p.Next
		p.Next = p.Next.Next
		rp := root
		for rp.Next.Val < tmp.Val {
			rp = rp.Next
		}
		tmp.Next = rp.Next
		rp.Next = tmp
	}
	return root.Next
}
