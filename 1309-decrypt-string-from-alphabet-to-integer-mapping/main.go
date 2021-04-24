package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	ans := ""
	i := len(s) - 1
	for i >= 0 {
		var code int
		if s[i] == '#' {
			code, _ = strconv.Atoi(s[i-2 : i])
			i -= 3
		} else {
			code, _ = strconv.Atoi(s[i : i+1])
			i--
		}
		ans = fmt.Sprintf("%s", string(rune(code+96))) + ans
	}
	return ans
}
