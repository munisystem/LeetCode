package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if v := x / 10; v == 0 {
		return true
	}
	i := 0
	j := x
	for {
		i = i*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}
	return i == j
}
