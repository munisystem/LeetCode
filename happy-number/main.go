package main

func isHappy(n int) bool {
	for {
		if n/10 == 0 {
			if n == 1 || n == 7 {
				return true
			} else {
				return false
			}
		}
		r := 0
		for x := n; x > 0; x = x / 10 {
			r = r + (x%10)*(x%10)
		}
		n = r
	}
}
