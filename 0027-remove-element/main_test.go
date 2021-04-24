package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []int
		in2 int
		out int
	}{
		{in1: []int{3, 2, 2, 3}, in2: 3, out: 2},
		{in1: []int{0, 1, 2, 2, 3, 0, 4, 2}, in2: 2, out: 5},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %d", tt.in1, tt.in2), func(t *testing.T) {
			actual := removeElement(tt.in1, tt.in2)
			if actual != tt.out {
				t.Errorf("in1 %v, in2 %d, got %d, want %d", tt.in1, tt.in2, actual, tt.out)
			}
		})
	}
}
