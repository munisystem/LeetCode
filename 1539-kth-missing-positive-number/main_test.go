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
		{in1: []int{2, 3, 4, 7, 11}, in2: 5, out: 9},
		{in1: []int{1, 2, 3, 4}, in2: 2, out: 6},
		{in1: []int{1, 3, 5, 8, 11}, in2: 6, out: 10},
		{in1: []int{1, 3, 5, 8, 100}, in2: 50, out: 54},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := findKthPositive(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
