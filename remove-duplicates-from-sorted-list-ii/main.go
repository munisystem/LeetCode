package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			v := p.Val
			for p != nil && p.Val == v {
				p = p.Next
			}
		} else {
			node.Next = p
			node = node.Next
			p = p.Next
		}
	}
	if p != nil {
		node.Next = p
	} else {
		node.Next = nil
	}
	return root.Next
}
