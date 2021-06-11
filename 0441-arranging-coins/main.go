package main

func arrangeCoins(n int) int {
	row := 1
	for n-row >= 0 {
		n -= row
		row++
	}
	return row - 1
}
