package main

func totalMoney(n int) int {
	r := 0
	for i := 0; i < n; i++ {
		r = r + 1 + (i % 7) + (i / 7)
	}
	return r
}
