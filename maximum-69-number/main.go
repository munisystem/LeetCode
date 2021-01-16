package main

func maximum69Number(num int) int {
	digits := 1
	for i := num; i/10 != 0; i /= 10 {
		digits *= 10
	}
	r := 0
	for i := digits; i > 0; i /= 10 {
		if num/i == 6 {
			r += 9*i + num%i
			break
		}
		r += 9 * i
		num %= i
	}
	return r
}
