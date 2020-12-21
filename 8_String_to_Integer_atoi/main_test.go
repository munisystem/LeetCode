package main

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	testcase := []struct {
		in  string
		out int
	}{
		{in: "42", out: 42},
		{in: "   -42", out: -42},
		{in: "4193 with words", out: 4193},
		{in: "words and 987", out: 0},
		{in: "-91283472332", out: -2147483648},
		{in: "", out: 0},
		{in: "2147483648", out: 2147483647},
		{in: "9223372036854775808", out: 2147483647},
		{in: "21474836460", out: 2147483647},
	}

	for _, tt := range testcase {
		t.Run(tt.in, func(t *testing.T) {
			actual := myAtoi(tt.in)
			if actual != tt.out {
				t.Errorf("input %q, got %d, want %d", tt.in, actual, tt.out)
			}
		})
	}
}
