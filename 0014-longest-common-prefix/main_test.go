package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []string
		out string
	}{
		{in: []string{"flower", "flow", "flight"}, out: "fl"},
		{in: []string{"dog", "racecar", "car"}, out: ""},
		{in: []string{"", "dog", "racecar"}, out: ""},
		{in: []string{}, out: ""},
		{in: []string{"a"}, out: "a"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			actual := longestCommonPrefix(tt.in)
			if actual != tt.out {
				t.Errorf("input %v, got %s, want %s", tt.in, actual, tt.out)
			}
		})
	}
}
