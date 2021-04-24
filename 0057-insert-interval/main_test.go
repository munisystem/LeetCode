package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 [][]int
		in2 []int
		out [][]int
	}{
		{in1: [][]int{{1, 3}, {6, 9}}, in2: []int{2, 5}, out: [][]int{{1, 5}, {6, 9}}},
		{in1: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, in2: []int{4, 8}, out: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{in1: [][]int{{1, 5}}, in2: []int{2, 3}, out: [][]int{{1, 5}}},
		{in1: [][]int{{1, 5}}, in2: []int{2, 7}, out: [][]int{{1, 7}}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := insert(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
