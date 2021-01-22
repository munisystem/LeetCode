package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 []int
		in2 int
		out []int
	}{
		{in1: []int{3, 5, 2, 6}, in2: 2, out: []int{2, 6}},
		{in1: []int{2, 4, 3, 3, 5, 4, 9, 6}, in2: 4, out: []int{2, 3, 3, 4}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := mostCompetitive(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
