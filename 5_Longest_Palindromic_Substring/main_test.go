package main

import "testing"

func TestLongestPalindromic(t *testing.T) {
	testcase := []struct {
		in  string
		out string
	}{
		{in: "babad", out: "bab"},
		{in: "cbbd", out: "bb"},
		{in: "a", out: "a"},
		{in: "ac", out: "a"},
		{in: "bb", out: "bb"},
	}

	for _, tt := range testcase {
		t.Run(tt.in, func(t *testing.T) {
			actual := longestPalindrome(tt.in)
			if actual != tt.out {
				t.Errorf("input %q, got %q, want %q", tt.in, actual, tt.out)
			}
		})
	}
}
