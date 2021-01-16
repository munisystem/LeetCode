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
		{in1: []int{3, 2, 1, 5, 6, 4}, in2: 2, out: 5},
		{in1: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, in2: 4, out: 4},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := findKthLargest(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
