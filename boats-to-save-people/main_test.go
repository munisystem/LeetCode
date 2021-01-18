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
		{in1: []int{1, 2}, in2: 3, out: 1},
		{in1: []int{3, 2, 2, 1}, in2: 3, out: 3},
		{in1: []int{3, 5, 3, 4}, in2: 5, out: 4},
		{in1: []int{3, 2, 3, 2, 2}, in2: 6, out: 3},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := numRescueBoats(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
