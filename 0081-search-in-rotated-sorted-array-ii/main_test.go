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
		out bool
	}{
		{in1: []int{2, 5, 6, 0, 0, 1, 2}, in2: 0, out: true},
		{in1: []int{2, 5, 6, 0, 0, 1, 2}, in2: 3, out: false},
		{in1: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, in2: 2, out: true},
		{in1: []int{3, 3, 0, 1, 3}, in2: 1, out: true},
		{in1: []int{3, 1}, in2: 1, out: true},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := search(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
