package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p := list1
	for i := 0; i < a-1; i++ {
		p = p.Next
	}
	t := p
	p = p.Next
	t.Next = list2
	for t.Next != nil {
		t = t.Next
	}
	for i := 0; i < b-a; i++ {
		p = p.Next
	}
	t.Next = p.Next
	return list1
}
