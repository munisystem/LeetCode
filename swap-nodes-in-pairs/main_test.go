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
		in  *ListNode
		out *ListNode
	}{
		{in: node([]int{1, 2, 3, 4}), out: node([]int{2, 1, 4, 3})},
		{in: node([]int{}), out: node([]int{})},
		{in: node([]int{1}), out: node([]int{1})},
		{in: node([]int{1, 2, 3}), out: node([]int{2, 1, 3})},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			actual := swapPairs(tt.in)
			if diff := cmp.Diff(tt.out, actual); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
