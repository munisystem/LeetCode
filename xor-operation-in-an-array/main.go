package main

func xorOperation(n int, start int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum ^= i*2 + start
	}
	return sum
}
