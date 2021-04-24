package main

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	r := 0
	i := 0
	for j := 0; j < len(s); j++ {
		for strings.Contains(s[i:j], string(s[j])) {
			i++
		}
		r = int(math.Max(float64(r), float64(len(s[i:j+1]))))
	}
	return r
}
