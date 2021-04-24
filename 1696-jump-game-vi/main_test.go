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
		{in1: []int{1, -1, -2, 4, -7, 3}, in2: 2, out: 7},
		{in1: []int{10, -5, -2, 4, 0, 3}, in2: 3, out: 17},
		{in1: []int{1, -5, -20, 4, -1, 3, -6, -3}, in2: 2, out: 0},
		{in1: []int{0}, in2: 1, out: 0},
		{in1: []int{0, -1, -2, -3, 1}, in2: 2, out: -1},
		{in1: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, in2: 3, out: -4},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v,%v", tt.in1, tt.in2), func(t *testing.T) {
			got := maxResult(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in1 %v, in2 %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
