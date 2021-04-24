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
		{in1: []int{2, 5, 1, 3, 4, 7}, in2: 3, out: []int{2, 3, 5, 4, 1, 7}},
		{in1: []int{1, 2, 3, 4, 4, 3, 2, 1}, in2: 4, out: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{in1: []int{1, 1, 2, 2}, in2: 2, out: []int{1, 2, 1, 2}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := shuffle(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
