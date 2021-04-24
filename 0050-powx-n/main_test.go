package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 float64
		in2 int
		out float64
	}{
		{in1: 2.00000, in2: 10, out: 1024.00000},
		{in1: 2.10000, in2: 3, out: 9.26100},
		{in1: 2.00000, in2: -2, out: 0.25000},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%f, %d", tt.in1, tt.in2), func(t *testing.T) {
			got := myPow(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in1 %f, in2 %d mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
