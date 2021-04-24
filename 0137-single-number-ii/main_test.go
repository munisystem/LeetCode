package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out int
	}{
		{in: []int{2, 2, 3, 2}, out: 3},
		{in: []int{0, 1, 0, 1, 0, 1, 99}, out: 99},
		{in: []int{0, 1, 0, 2, 0, 2, 2}, out: 1},
		{in: []int{0, 1, 1, 1}, out: 0},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := singleNumber(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
