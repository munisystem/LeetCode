package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}
	op := head
	dummy := &Node{}
	p := dummy
	for i := 0; op != nil; i++ {
		p.Next = &Node{Val: op.Val}
		m[op] = p.Next
		p, op = p.Next, op.Next
	}
	op = head
	p = dummy
	for op != nil {
		if op.Random != nil {
			p.Next.Random = m[op.Random]
		}
		p, op = p.Next, op.Next
	}
	return dummy.Next
}
