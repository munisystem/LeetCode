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
		out bool
	}{
		{in1: "abc", in2: "bca", out: true},
		{in1: "a", in2: "aa", out: false},
		{in1: "cabbba", in2: "abbccc", out: true},
		{in1: "cabbba", in2: "aabbss", out: false},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := closeStrings(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
