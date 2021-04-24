package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  [][]int
		out int
	}{
		{in: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, out: 11},
		{in: [][]int{{-10}}, out: -10},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minimumTotal(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
