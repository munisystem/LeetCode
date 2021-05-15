package main

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	patterns := map[string]int{}
	chars := make([]bool, 26)
	if len(pattern) != len(strs) {
		return false
	}
	for i := 0; i < len(strs); i++ {
		p := int(pattern[i]) - 'a'
		if v, ok := patterns[strs[i]]; ok {
			if v != p {
				return false
			}
		} else if chars[p] {
			return false
		} else {
			patterns[strs[i]] = p
			chars[p] = true
		}
	}
	return true
}
