package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []string
		in2 []string
		out bool
	}{
		{in1: []string{"ab", "c"}, in2: []string{"a", "bc"}, out: true},
		{in1: []string{"a", "cb"}, in2: []string{"ab", "c"}, out: false},
		{in1: []string{"abc", "d", "defg"}, in2: []string{"abcddefg"}, out: true},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := arrayStringsAreEqual(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v and %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
