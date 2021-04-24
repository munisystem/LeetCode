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
		{in: []int{0, 1, 0}, out: 1},
		{in: []int{0, 2, 1, 0}, out: 1},
		{in: []int{0, 10, 5, 2}, out: 1},
		{in: []int{3, 4, 5, 1}, out: 2},
		{in: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, out: 2},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := peakIndexInMountainArray(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
