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
		{in1: []int{2, 7, 11, 15}, in2: 9, out: []int{1, 2}},
		{in1: []int{2, 3, 4}, in2: 6, out: []int{1, 3}},
		{in1: []int{-1, 0}, in2: -1, out: []int{1, 2}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := twoSum(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
