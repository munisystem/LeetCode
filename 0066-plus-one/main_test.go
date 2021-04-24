package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out []int
	}{
		{in: []int{1, 2, 3}, out: []int{1, 2, 4}},
		{in: []int{4, 3, 2, 1}, out: []int{4, 3, 2, 2}},
		{in: []int{0}, out: []int{1}},
		{in: []int{9, 9, 9}, out: []int{1, 0, 0, 0}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := plusOne(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
