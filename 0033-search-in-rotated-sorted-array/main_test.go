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
		{in1: []int{4, 5, 6, 7, 0, 1, 2}, in2: 0, out: 4},
		{in1: []int{4, 5, 6, 7, 0, 1, 2}, in2: 4, out: 0},
		{in1: []int{4, 5, 6, 7, 0, 1, 2}, in2: 3, out: -1},
		{in1: []int{1}, in2: 0, out: -1},
		{in1: []int{0, 1}, in2: 1, out: 1},
		{in1: []int{0, 1}, in2: 0, out: 0},
		{in1: []int{3, 1}, in2: 1, out: 1},
		{in1: []int{5, 1, 3}, in2: 5, out: 0},
		{in1: []int{3, 1}, in2: 4, out: -1},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := search(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
