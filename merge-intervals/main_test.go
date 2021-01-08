package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  [][]int
		out [][]int
	}{
		{in: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, out: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{in: [][]int{{1, 4}, {4, 5}}, out: [][]int{{1, 5}}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := merge(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
