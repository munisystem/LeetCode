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
		in3 []string
		out int
	}{
		{in1: "hit", in2: "cog", in3: []string{"hot", "dot", "dog", "lot", "log", "cog"}, out: 5},
		{in1: "hit", in2: "cog", in3: []string{"hot", "dot", "dog", "lot", "log"}, out: 0},
		{in1: "hit", in2: "coq", in3: []string{"hot", "dot", "dog", "lot", "log", "cog", "coq"}, out: 6},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v, %v", tt.in1, tt.in2, tt.in3), func(t *testing.T) {
			got := ladderLength(tt.in1, tt.in2, tt.in3)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, tt.in3, diff)
			}
		})
	}
}
