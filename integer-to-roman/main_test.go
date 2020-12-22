package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  int
		out string
	}{
		{in: 3, out: "III"},
		{in: 4, out: "IV"},
		{in: 9, out: "IX"},
		{in: 58, out: "LVIII"},
		{in: 1994, out: "MCMXCIV"},
		{in: 1304, out: "MCCCIV"},
		{in: 14, out: "XIV"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			actual := intToRoman(tt.in)
			if actual != tt.out {
				t.Errorf("input %d, got %s, want %s", tt.in, actual, tt.out)
			}
		})
	}
}
