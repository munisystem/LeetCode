package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out bool
	}{
		{in: []int{1, 2, 2, 1, 1, 3}, out: true},
		{in: []int{1, 2}, out: false},
		{in: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, out: true},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := uniqueOccurrences(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
