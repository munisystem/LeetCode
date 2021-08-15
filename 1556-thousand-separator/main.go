package main

import (
	"strconv"
)

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	var ans []rune
	for i, v := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			ans = append(ans, '.')
		}
		ans = append(ans, v)
	}
	return string(ans)
}
