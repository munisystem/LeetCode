package main

func longestValidParentheses(s string) int {
	max := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ')' {
			continue
		}
		p := 1
		for j := i + 1; j < len(s); j++ {
			switch s[j] {
			case ')':
				p--
			case '(':
				p++
			}
			if p == 0 && max < j-i+1 {
				max = j - i + 1
			}
			if p < 0 {
				break
			}
		}
	}
	return max
}
