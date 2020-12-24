package main

func generateParenthesis(n int) []string {
	return p("", n, n)
}

func p(s string, o, c int) []string {
	if o == 0 && c == 0 {
		return []string{s}
	}
	parenthesis := []string{}
	if o > 0 {
		parenthesis = append(parenthesis, p(s+"(", o-1, c)...)
	}
	if o < c && c > 0 {
		parenthesis = append(parenthesis, p(s+")", o, c-1)...)
	}
	return parenthesis
}
