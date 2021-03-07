package main

func balancedStringSplit(s string) int {
	cond, count := 0, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "L" {
			cond--
		} else {
			cond++
		}
		if cond == 0 {
			count++
		}
	}
	return count
}
