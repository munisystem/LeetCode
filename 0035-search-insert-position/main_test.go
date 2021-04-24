package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []int
		in2 int
		out int
	}{
		{in1: []int{1, 3, 5, 6}, in2: 5, out: 2},
		{in1: []int{1, 3, 5, 6}, in2: 2, out: 1},
		{in1: []int{1, 3, 5, 6}, in2: 7, out: 4},
		{in1: []int{1, 3, 5, 6}, in2: 0, out: 0},
		{in1: []int{1}, in2: 0, out: 0},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := searchInsert(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
