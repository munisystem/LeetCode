package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  []int
		out int
	}{
		{in: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, out: 3},
		{in: []int{7}, out: 0},
		{in: []int{7, 6, 9, 6, 9, 6, 9, 7}, out: 1},
		{in: []int{6, 1, 9}, out: 2},
		{in: []int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}, out: 3},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minJumps(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
