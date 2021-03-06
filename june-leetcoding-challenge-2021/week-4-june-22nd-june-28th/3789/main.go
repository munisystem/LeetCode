package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	root := &ListNode{}
	root.Next = head
	p := root
	for i := 0; i < m-1; i++ {
		p = p.Next
	}
	pp := p.Next
	for i := 0; i < n-m; i++ {
		tmp := p.Next
		p.Next = pp.Next
		pp.Next = pp.Next.Next
		p.Next.Next = tmp
	}
	return root.Next
}
