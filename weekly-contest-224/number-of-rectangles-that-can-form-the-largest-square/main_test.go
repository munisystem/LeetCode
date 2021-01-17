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
		{in: [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}, out: 3},
		{in: [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}, out: 3},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countGoodRectangles(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
