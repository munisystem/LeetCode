package main

func reverse(x int) int {
	n := 0
	for {
		var mod int
		x, mod = x/10, x%10
		n = 10*n + mod
		if x == 0 {
			break
		}
	}
	if n > 2147483647 || n < -2147483647 {
		return 0
	}
	return n
}
