package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func node(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	root := &ListNode{}
	p := root
	for i := range vals {
		p.Val = vals[i]
		if i != len(vals)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return root
}

func Test(t *testing.T) {
	testcase := []struct {
		in1 *ListNode
		in2 int
		out *ListNode
	}{
		{in1: node([]int{1, 2, 3, 4, 5}), in2: 2, out: node([]int{1, 2, 3, 5})},
		{in1: node([]int{1}), in2: 1, out: nil},
		{in1: node([]int{1, 2}), in2: 1, out: node([]int{1})},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %d", tt.in1, tt.in2), func(t *testing.T) {
			actual := removeNthFromEnd(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, actual); diff != "" {
				t.Errorf("%v and %d mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
