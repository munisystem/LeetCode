package main

var memoize = make([]int, 38, 38)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}
	if memoize[n] != 0 {
		return memoize[n]
	}
	memoize[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)

	return memoize[n]
}
