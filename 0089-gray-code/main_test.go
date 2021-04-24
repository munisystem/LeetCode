package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  int
		out []int
	}{
		{in: 2, out: []int{0, 1, 3, 2}},
		{in: 1, out: []int{0, 1}},
		{in: 3, out: []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := grayCode(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
