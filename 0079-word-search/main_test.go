package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	testcase := []struct {
		in1 [][]byte
		in2 string
		out bool
	}{
		{in1: [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, in2: "ABCCED", out: true},
		{in1: [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, in2: "SEE", out: true},
		{in1: [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, in2: "ABCD", out: false},
		{in1: [][]byte{[]byte("CAA"), []byte("AAA"), []byte("BCD")}, in2: "AAB", out: true},
	}

	for _, tt := range testcase {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := exist(tt.in1, tt.in2)
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("in %v, %v mismatch (-want +got):\n%s", tt.in1, tt.in2, diff)
			}
		})
	}
}
