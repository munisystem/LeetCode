package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 string
		in2 string
		out string
	}{
		{in1: "11", in2: "1", out: "100"},
		{in1: "1010", in2: "1011", out: "10101"},
		{in1: "11111", in2: "11", out: "100010"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := addBinary(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
