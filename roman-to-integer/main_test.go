package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  string
		out int
	}{
		{in: "III", out: 3},
		{in: "IV", out: 4},
		{in: "IX", out: 9},
		{in: "LVIII", out: 58},
		{in: "MCMXCIV", out: 1994},
		{in: "MCCCIV", out: 1304},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%s", tt.in), func(t *testing.T) {
			actual := romanToInt(tt.in)
			if actual != tt.out {
				t.Errorf("input %s, got %d, want %d", tt.in, actual, tt.out)
			}
		})
	}
}
