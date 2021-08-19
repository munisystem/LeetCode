package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	var ans int
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			ans++
		}
	}
	return ans
}
