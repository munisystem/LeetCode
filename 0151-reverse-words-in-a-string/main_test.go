package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in  string
		out string
	}{
		{in: "the sky is blue", out: "blue is sky the"},
		{in: "  hello world  ", out: "world hello"},
		{in: "a good   example", out: "example good a"},
		{in: "  Bob    Loves  Alice   ", out: "Alice Loves Bob"},
		{in: "Alice does not even like bob", out: "bob like even not does Alice"},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := reverseWords(tt.in)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v mismatch (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
