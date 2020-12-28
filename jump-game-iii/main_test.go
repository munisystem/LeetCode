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
		out bool
	}{
		{in1: []int{4, 2, 3, 0, 3, 1, 2}, in2: 5, out: true},
		{in1: []int{4, 2, 3, 0, 3, 1, 2}, in2: 0, out: true},
		{in1: []int{3, 0, 2, 1, 2}, in2: 2, out: false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v,%v", tt.in1, tt.in2), func(t *testing.T) {
			got := canReach(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in1 %v and in2 %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
