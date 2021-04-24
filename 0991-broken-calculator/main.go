package main

func brokenCalc(X int, Y int) int {
	ans := 0
	for Y > X {
		if Y%2 == 1 {
			Y++
		} else {
			Y /= 2
		}
		ans++
	}
	return ans + X - Y
}
