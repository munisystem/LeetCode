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
		{in: "Hello World", out: 5},
		{in: " ", out: 0},
		{in: "a ", out: 1},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := lengthOfLastWord(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
