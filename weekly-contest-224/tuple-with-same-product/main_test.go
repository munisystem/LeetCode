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
		{in: []int{2, 3, 4, 6}, out: 8},
		{in: []int{1, 2, 4, 5, 10}, out: 16},
		{in: []int{2, 3, 4, 6, 8, 12}, out: 40},
		{in: []int{2, 3, 5, 7}, out: 0},
		{in: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}, out: 1288},
		{in: []int{20, 10, 6, 24, 15, 5, 4, 30}, out: 48},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := tupleSameProduct(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
