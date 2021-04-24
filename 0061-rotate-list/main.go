package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	} else if k == 0 {
		return head
	}

	root := head
	p := root
	size := 1
	for p.Next != nil {
		p = p.Next
		size++
	}
	p.Next = head
	p = root
	rotate := size - (k % size) - 1
	for i := 0; i < rotate; i++ {
		p = p.Next
	}
	r := p.Next
	p.Next = nil
	return r
}
