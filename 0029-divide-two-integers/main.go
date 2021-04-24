package main

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}
	signed := false
	switch {
	case dividend < 0 && divisor < 0:
		dividend = -dividend
		divisor = -divisor
	case dividend > 0 && divisor < 0:
		divisor = -divisor
		signed = true
	case dividend < 0 && divisor > 0:
		dividend = -dividend
		signed = true
	}
	i, _ := div(dividend, divisor, 1)
	if signed {
		i = -i
	}
	if i > math.MaxInt32 || i < math.MinInt32 {
		return math.MaxInt32
	}
	return i
}

func div(dividend, divisor, quotient int) (int, int) {
	q := 0
	if dividend >= divisor+divisor {
		q, dividend = div(dividend, divisor+divisor, quotient+quotient)
	}
	for dividend >= divisor {
		dividend = dividend - divisor
		q = q + quotient
	}
	return q, dividend
}
