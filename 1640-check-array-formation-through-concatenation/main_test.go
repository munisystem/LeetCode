package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []int
		in2 [][]int
		out bool
	}{
		{in1: []int{85}, in2: [][]int{{85}}, out: true},
		{in1: []int{15, 88}, in2: [][]int{{88}, {15}}, out: true},
		{in1: []int{49, 18, 16}, in2: [][]int{{16, 18, 49}}, out: false},
		{in1: []int{91, 4, 64, 78}, in2: [][]int{{78}, {4, 64}, {91}}, out: true},
		{in1: []int{1, 3, 5, 7}, in2: [][]int{{2, 4, 6, 8}}, out: false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v and %v", tt.in1, tt.in2), func(t *testing.T) {
			got := canFormArray(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
