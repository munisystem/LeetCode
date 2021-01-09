package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  [][]int
		out []int
	}{
		{in: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{in: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, out: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{in: [][]int{{1}, {2}, {3}}, out: []int{1, 2, 3}},
		{in: [][]int{{1, 2, 3}}, out: []int{1, 2, 3}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := spiralOrder(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
