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
		{in1: "2", in2: "3", out: "6"},
		{in1: "123", in2: "456", out: "56088"},
		{in1: "0", in2: "1234", out: "0"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v,%v", tt.in1, tt.in2), func(t *testing.T) {
			got := multiply(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in1 %v and in2 %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
