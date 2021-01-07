package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	root := head
	slow, fast := head, head
	for {
		slow = next(slow)
		fast = next(next(fast))
		if slow == fast {
			break
		}
	}
	if slow == nil {
		return nil
	}
	n := 0
	fast = root
	for ; slow != fast; n++ {
		slow = next(slow)
		fast = next(fast)
	}
	return slow
}

func next(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}
