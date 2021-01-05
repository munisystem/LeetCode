package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  string
		out int
	}{
		{in: "32", out: 3},
		{in: "82734", out: 8},
		{in: "27346209830709182346", out: 9},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minPartitions(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
