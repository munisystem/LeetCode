package main

import (
	"sort"
)

func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	refs := []string{}
	size := 0
	for i := 0; i < len(words); i++ {
		isMatch := false
		for j := 0; j < len(refs); j++ {
			if words[i] == refs[j][len(refs[j])-len(words[i]):] {
				isMatch = true
				break
			}
		}
		if !isMatch {
			size += len(words[i]) + 1
			refs = append(refs, words[i])
		}
	}
	return size
}
