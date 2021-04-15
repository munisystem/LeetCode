package main

var mem = map[int]int{}

func fib(n int) int {
	if v, ok := mem[n]; ok {
		return v
	}
	if n <= 1 {
		mem[n] = n
		return mem[n]
	}
	mem[n] = fib(n-1) + fib(n-2)
	return mem[n]
}
