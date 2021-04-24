package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 string
		in2 []string
		out []string
	}{
		{in1: "catsanddog", in2: []string{"cat", "cats", "and", "sand", "dog"}, out: []string{"cat sand dog", "cats and dog"}},
		{in1: "pineapplepenapple", in2: []string{"apple", "pen", "applepen", "pine", "pineapple"}, out: []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"}},
		{in1: "catsandog", in2: []string{"cats", "dog", "sand", "and", "cat"}, out: []string{}},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := wordBreak(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
