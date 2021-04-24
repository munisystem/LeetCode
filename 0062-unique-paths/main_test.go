package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 int
		in2 int
		out int
	}{
		{in1: 3, in2: 7, out: 28},
		{in1: 3, in2: 2, out: 3},
		{in1: 1, in2: 100, out: 1},
		{in1: 2, in2: 100, out: 100},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := uniquePaths(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
