package main

// Floyd's cycle-finding algorithm
// c.f. https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare
func sumDigits(n int) int {
	r := 0
	for x := n; x > 0; x = x / 10 {
		r = r + (x%10)*(x%10)
	}
	return r
}

func isHappy(n int) bool {
	fast, slow := n, n
	for slow != 1 {
		fast = sumDigits(sumDigits(fast))
		slow = sumDigits(slow)
		if slow != 1 && slow == fast {
			return false
		}
	}
	return true
}
