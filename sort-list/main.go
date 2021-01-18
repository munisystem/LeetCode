package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	var mid *ListNode
	for p != nil && p.Next != nil {
		if mid == nil {
			mid = head
		} else {
			mid = mid.Next
		}
		p = p.Next.Next
	}
	t := mid.Next
	mid.Next = nil
	l1, l2 := mergeSort(head), mergeSort(t)
	root := &ListNode{}
	p = root
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return root.Next
}
