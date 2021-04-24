package main

func lengthOfLastWord(s string) int {
	l := 0
	l, t := 0, len(s)-1
	for t >= 0 && string(s[t]) == " " {
		t--
	}
	for t >= 0 && string(s[t]) != " " {
		l++
		t--
	}
	return l
}
