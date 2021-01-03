package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out [][]int
	}{
		{in: []int{-1, 0, 1, 2, -1, -4}, out: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{in: []int{}, out: [][]int{}},
		{in: []int{0}, out: [][]int{}},
		{in: []int{0, 0, 0, 0}, out: [][]int{{0, 0, 0}}},
		{in: []int{0, 0, 0, 0, 0}, out: [][]int{{0, 0, 0}}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := threeSum(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
