package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var size int
	p := head
	for p != nil {
		size++
		p = p.Next
	}
	ans := make([]*ListNode, k, k)
	p = head
	for i := 0; i < k; i++ {
		if p == nil {
			ans[i] = nil
			continue
		}
		for j := 0; j < (size/k)-1; j++ {
			head = head.Next
		}
		if size/k > 0 && size%k > i {
			head = head.Next
		}
		tmp := head.Next
		head.Next = nil
		ans[i] = p
		p = tmp
		head = tmp
	}
	return ans
}
