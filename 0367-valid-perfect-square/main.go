package main

func isPerfectSquare(num int) bool {
	n := num
	for n*n > num {
		n = (n + num/n) / 2
	}
	return n*n == num
}
