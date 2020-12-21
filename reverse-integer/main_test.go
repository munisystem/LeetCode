package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  int
		out int
	}{
		{in: 123, out: 321},
		{in: -123, out: -321},
		{in: 120, out: 21},
		{in: 0, out: 0},
		{in: 1534236469, out: 0},
		{in: -2147483648, out: 0},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			actual := reverse(tt.in)
			if actual != tt.out {
				t.Errorf("input %d, got %d, want %d", tt.in, actual, tt.out)
			}
		})
	}
}
