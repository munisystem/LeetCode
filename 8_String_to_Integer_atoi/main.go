package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	sign, n := 1, 0
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	}
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)
		switch {
		case sign == 1 && n > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && n*sign < math.MinInt32:
			return math.MinInt32
		}
	}

	n = n * sign
	return n
}
