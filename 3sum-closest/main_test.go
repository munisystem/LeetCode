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
		{in1: []int{-1, 2, 1, -4}, in2: 1, out: 2},
		{in1: []int{-1, 1, 1, -4}, in2: 1, out: 1},
		{in1: []int{-1, 2, 2, -4}, in2: 0, out: 0},
		{in1: []int{-1, 2, 2, -3}, in2: 0, out: 1},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := threeSumClosest(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
