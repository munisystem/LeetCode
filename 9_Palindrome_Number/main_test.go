package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testcase := []struct {
		in  int
		out bool
	}{
		{in: 121, out: true},
		{in: -121, out: false},
		{in: 10, out: false},
		{in: -101, out: false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			actual := isPalindrome(tt.in)
			if actual != tt.out {
				t.Errorf("input %d, got %t, want %t", tt.in, actual, tt.out)
			}
		})
	}
}
