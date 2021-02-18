package main

func maxDepth(s string) int {
	ans, depth := 0, 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			depth++
			if depth > ans {
				ans = depth
			}
		case ")":
			depth--
		}
	}
	return ans
}
