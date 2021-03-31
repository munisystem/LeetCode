package main

func findComplement(num int) int {
	m := ^0
	for m&num > 0 {
		m <<= 1
	}
	return ^m ^ num
}
