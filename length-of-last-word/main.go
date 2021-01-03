package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	if len(ss) == 0 {
		return 0
	}
	return len(ss[len(ss)-1])
}
