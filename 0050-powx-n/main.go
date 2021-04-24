package main

import "fmt"

func myPow(x float64, n int) float64 {
	fmt.Println(x, n)
	if n == 0 {
		return 1
	}
	m := x
	signed := false
	if n < 0 {
		n = -n
		signed = true
	}
	x, _ = pow(x, m, 1, n)
	if signed {
		x = 1 / x
	}
	return x
}

func pow(x, m float64, i, n int) (float64, int) {
	if (i * 2) < n {
		x, n = pow(x, m*m, i*2, n)
	}
	for n > 0 && n > i {
		n = n - i
		x = x * m
	}
	return x, n
}
