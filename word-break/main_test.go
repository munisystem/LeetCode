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
		out bool
	}{
		{in1: "leetcode", in2: []string{"leet", "code"}, out: true},
		{in1: "applepenapple", in2: []string{"apple", "pen"}, out: true},
		{in1: "catsandog", in2: []string{"cats", "dog", "sand", "and", "cat"}, out: false},
		{in1: "bb", in2: []string{"a", "b", "bbb", "bbbb"}, out: true},
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
