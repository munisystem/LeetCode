package main

import (
	"strings"
)

func findLongestWord(s string, d []string) string {
	ans := ""
	for i := 0; i < len(d); i++ {
		c := 0
		for j := 0; j < len(s); j++ {
			if c < len(d[i]) && s[j] == d[i][c] {
				c++
			}
		}
		if c == len(d[i]) && (len(ans) < len(d[i]) || (len(ans) == len(d[i]) && strings.Compare(ans, d[i]) > 0)) {
			ans = d[i]
		}
	}
	return ans
}
