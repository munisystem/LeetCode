package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	sizeA, sizeB := 0, 0
	for pa.Next != nil {
		sizeA++
		pa = pa.Next
	}
	for pb.Next != nil {
		sizeB++
		pb = pb.Next
	}
	if pa != pb {
		return nil
	}
	if sizeA > sizeB {
		for i := 0; i < sizeA-sizeB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < sizeB-sizeA; i++ {
			headB = headB.Next
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
