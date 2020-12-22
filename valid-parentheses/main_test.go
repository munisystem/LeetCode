package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  string
		out bool
	}{
		{in: "()", out: true},
		{in: "()[]{}", out: true},
		{in: "(]", out: false},
		{in: "([)]", out: false},
		{in: "{[]}", out: true},
		{in: "((", out: false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%s", tt.in), func(t *testing.T) {
			actual := isValid(tt.in)
			if actual != tt.out {
				t.Errorf("input %s, got %v, want %v", tt.in, actual, tt.out)
			}
		})
	}
}
