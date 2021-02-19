package main

func removeOuterParentheses(S string) string {
	p := 0
	ans := ""
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "(" {
			if p != 0 {
				ans += string(S[i])
			}
			p++
		} else {
			if p != 1 {
				ans += string(S[i])
			}
			p--
		}
	}
	return ans
}
