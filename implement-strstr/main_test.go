package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 string
		in2 string
		out int
	}{
		{in1: "hello", in2: "ll", out: 2},
		{in1: "aaaaa", in2: "baa", out: -1},
		{in1: "", in2: "", out: 0},
		{in1: "aaaa", in2: "", out: 0},
		{in1: "", in2: "aaaa", out: -1},
		{in1: "a", in2: "a", out: 0},
		{in1: "hello", in2: "lo", out: 3},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%s, %s", tt.in1, tt.in2), func(t *testing.T) {
			actual := strStr(tt.in1, tt.in2)
			if actual != tt.out {
				t.Errorf("in1 %s in2 %s, got %d, want %d", tt.in1, tt.in2, actual, tt.out)
			}
		})
	}
}
