package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var rev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := rev
		rev = slow
		slow = slow.Next
		rev.Next = tmp
	}
	if fast != nil {
		slow = slow.Next
	}
	for rev != nil {
		if rev.Val != slow.Val {
			return false
		}
		slow, rev = slow.Next, rev.Next
	}
	return true
}
