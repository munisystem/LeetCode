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
		{in: []int{3, 4, 5, 1, 2}, out: 1},
		{in: []int{4, 5, 6, 7, 0, 1, 2}, out: 0},
		{in: []int{11, 13, 15, 17}, out: 11},
		{in: []int{5, 1, 3}, out: 1},
		{in: []int{1, 3, 5}, out: 1},
		{in: []int{2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2}, out: 1},
		{in: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2}, out: 1},
		{in: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, out: 2},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findMin(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
