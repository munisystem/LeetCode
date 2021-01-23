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
		{in: []int{-5, 1, 5, 0, -7}, out: 1},
		{in: []int{-4, -3, -2, -1, 4, 3, 2}, out: 0},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := largestAltitude(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
