package main

func mySqrt(x int) int {
	// Newton's method
	// ref. https://en.wikipedia.org/wiki/Integer_square_root#Using_only_integer_division
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
