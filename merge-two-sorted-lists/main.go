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
	for !(p1 == nil && p2 == nil) {
		if p1 == nil {
			p.Val = p2.Val
			p2 = p2.Next
		} else if p2 == nil {
			p.Val = p1.Val
			p1 = p1.Next
		} else if v1, v2 := p1.Val, p2.Val; v1 <= v2 {
			p.Val = v1
			p1 = p1.Next
		} else {
			p.Val = p2.Val
			p2 = p2.Next
		}
		if !(p1 == nil && p2 == nil) {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return root
}
