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
		{in1: 10, in2: 3, out: 3},
		{in1: 7, in2: -3, out: -2},
		{in1: 0, in2: 1, out: 0},
		{in1: 1, in2: 1, out: 1},
		{in1: 1, in2: -1, out: -1},
		{in1: 4, in2: 2, out: 2},
		{in1: 4, in2: -2, out: -2},
		{in1: -4, in2: 2, out: -2},
		{in1: -4, in2: -2, out: 2},
		{in1: -2147483648, in2: -3, out: 715827882},
		{in1: -2147483648, in2: 1, out: -2147483648},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d,%d", tt.in1, tt.in2), func(t *testing.T) {
			got := divide(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in1 %d in2 %d mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
