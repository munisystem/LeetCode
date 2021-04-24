package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out int
	}{
		{in: []int{1, 2, 3, 1, 1, 3}, out: 4},
		{in: []int{1,1,1,1}, out: 6},
		{in: []int{1,2,3}, out: 0},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := numIdenticalPairs(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v  mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
