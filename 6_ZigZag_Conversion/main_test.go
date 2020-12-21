package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	testcase := []struct {
		in1 string
		in2 int
		out string
	}{
		{in1: "PAYPALISHIRING", in2: 3, out: "PAHNAPLSIIGYIR"},
		{in1: "PAYPALISHIRING", in2: 4, out: "PINALSIGYAHRPI"},
		{in1: "A", in2: 1, out: "A"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%q and %d", tt.in1, tt.in2), func(t *testing.T) {
			actual := convert(tt.in1, tt.in2)
			if actual != tt.out {
				t.Errorf("input1 %q, intput2 %d, got %q, want %q", tt.in1, tt.in2, actual, tt.out)
			}
		})
	}
}
