package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 int
		in2 int
		out int
	}{
		{in1: 12, in2: 3, out: 3},
		{in1: 7, in2: 2, out: 7},
		{in1: 4, in2: 4, out: -1},
		{in1: 1, in2: 1, out: 1},
		{in1: 1000, in2: 3, out: 4},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d, %d", tt.in1, tt.in2), func(t *testing.T) {
			actual := kthFactor(tt.in1, tt.in2)
			if actual != tt.out {
				t.Errorf("input1 %d, input2 %d, got %d, want %d", tt.in1, tt.in2, actual, tt.out)
			}
		})
	}
}
