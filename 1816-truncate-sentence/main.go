package main

import "strings"

func truncateSentence(s string, k int) string {
	strs := strings.Split(s, " ")
	return strings.Join(strs[:k], " ")
}
