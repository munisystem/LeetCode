package main

func isPrefixString(s string, words []string) bool {
	start := 0
	for i := 0; i < len(words) && start < len(s); i++ {
		if (len(s)-start) < len(words[i]) || s[start:start+len(words[i])] != words[i] {
			return false
		}
		start += len(words[i])
	}
	return start == len(s)
}
