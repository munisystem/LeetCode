package main

func hammingDistance(x int, y int) int {
	ans := 0
	n := x ^ y
	for n > 0 {
		ans++
		n &= n - 1
	}
	return ans
}
