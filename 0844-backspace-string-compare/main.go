package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	var char int
	for {
		char = 0
		for i >= 0 && (char > 0 || s[i] == '#') {
			if s[i] == '#' {
				char++
			} else {
				char--
			}
			i--
		}
		char = 0
		for j >= 0 && (char > 0 || t[j] == '#') {
			if t[j] == '#' {
				char++
			} else {
				char--
			}
			j--
		}
		if i >= 0 && j >= 0 && s[i] == t[j] {
			i--
			j--
		} else {
			break
		}
	}
	return i == -1 && j == -1
}
