package main

func plusOne(digits []int) []int {
	carry := true
	i := len(digits) - 1
	for i >= 0 && carry {
		digits[i] += 1
		if digits[i] > 9 {
			digits[i] = 0
			i--
		} else {
			carry = false
		}
	}
	if carry {
		return append([]int{1}, digits...)
	}
	return digits
}
