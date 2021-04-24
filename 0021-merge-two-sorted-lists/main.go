package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	root := &ListNode{}
	p := root
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p.Next = p2
			break
		} else if p2 == nil {
			p.Next = p1
			break
		}
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	return root.Next
}
