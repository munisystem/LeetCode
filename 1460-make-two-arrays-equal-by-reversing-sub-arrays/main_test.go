package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []int
		in2 []int
		out bool
	}{
		{in1: []int{1, 2, 3, 4}, in2: []int{2, 4, 1, 3}, out: true},
		{in1: []int{7}, in2: []int{7}, out: true},
		{in1: []int{1, 12}, in2: []int{12, 1}, out: true},
		{in1: []int{3, 7, 9}, in2: []int{3, 7, 11}, out: false},
		{in1: []int{1, 1, 1, 1, 1}, in2: []int{1, 1, 1, 1, 1}, out: true},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := canBeEqual(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
