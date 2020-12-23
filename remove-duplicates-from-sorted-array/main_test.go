package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out int
	}{
		{in: []int{1, 1, 2}, out: 2},
		{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: 5},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			actual := removeDuplicates(tt.in)
			if actual != tt.out {
				t.Errorf("input %v, got %d, want %d", tt.in, actual, tt.out)
			}
		})
	}
}
