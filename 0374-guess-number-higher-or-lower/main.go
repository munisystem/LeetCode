package main

func guessNumber(n int) int {
	l, r := 0, n
	var mid int
	for {
		mid := (l + r) / 2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			l = mid + 1
		case -1:
			r = mid
		}
	}
	return mid
}
