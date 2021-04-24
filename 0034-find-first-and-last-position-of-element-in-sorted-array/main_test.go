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
		out []int
	}{
		{in1: []int{5, 7, 7, 8, 8, 10}, in2: 8, out: []int{3, 4}},
		{in1: []int{5, 7, 7, 8, 8, 10}, in2: 6, out: []int{-1, -1}},
		{in1: []int{}, in2: 0, out: []int{-1, -1}},
		{in1: []int{1}, in2: 1, out: []int{0, 0}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := searchRange(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
