package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		si, sj := s[i], s[j]
		if !isAlphanumeric(si) {
			i++
			continue
		}
		if !isAlphanumeric(sj) {
			j--
			continue
		}
		if strings.ToLower(string(si)) != strings.ToLower(string(sj)) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}
