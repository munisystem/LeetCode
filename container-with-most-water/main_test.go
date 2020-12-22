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
		{in: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, out: 49},
		{in: []int{1, 1}, out: 1},
		{in: []int{1, 2, 1}, out: 2},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			actual := maxArea(tt.in)
			if actual != tt.out {
				t.Errorf("input %v, got %d, want %d", tt.in, actual, tt.out)
			}
		})
	}
}
