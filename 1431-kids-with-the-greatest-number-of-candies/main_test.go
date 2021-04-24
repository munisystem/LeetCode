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
		out []bool
	}{
		{in1: []int{2, 3, 5, 1, 3}, in2: 3, out: []bool{true, true, true, false, true}},
		{in1: []int{4, 2, 1, 1, 2}, in2: 1, out: []bool{true, false, false, false, false}},
		{in1: []int{12, 1, 12}, in2: 10, out: []bool{true, false, true}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := kidsWithCandies(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
