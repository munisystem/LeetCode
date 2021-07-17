package main

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for n > 0 {
		m := n % 10
		p *= m
		s += m
		n = n / 10
	}
	return p - s
}
