package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	r, p, i, isBreak := "", "", 1, false
	for {
		for _, str := range strs {
			if len(str) == 0 {
				isBreak = true
				break
			}
			if p == "" {
				if i > len(str) {
					isBreak = true
					break
				}
				p = str[0:i]
			}
			if len(p) > len(str) {
				isBreak = true
				break
			}
			if p != str[0:i] {
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
		r, p, i = p, "", i+1
	}
	return r
}
