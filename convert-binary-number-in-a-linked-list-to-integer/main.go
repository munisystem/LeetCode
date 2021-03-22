package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans<<1 + head.Val
		head = head.Next
	}
	return ans
}
