package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	root := &ListNode{}
	p := root
	for l1 != nil || l2 != nil || carry != 0 {
		p.Next = &ListNode{}
		p = p.Next
		l1v, l2v := 0, 0
		if l1 != nil {
			l1v = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
		}
		p.Val, carry = (l1v+l2v+carry)%10, (l1v+l2v+carry)/10
	}
	return root.Next
}
