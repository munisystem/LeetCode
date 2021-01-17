package main

func generateMatrix(n int) [][]int {
	rl, cl, rh, ch := 0, 0, n-1, n-1
	r := make([][]int, n, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n, n)
	}
	num := 1
	for {
		for i := cl; i <= ch; i++ {
			r[rl][i] = num
			num++
		}
		if num > n*n {
			break
		}
		for i := rl + 1; i <= rh; i++ {
			r[i][ch] = num
			num++
		}
		if num > n*n {
			break
		}
		for i := ch - 1; i >= cl; i-- {
			r[rh][i] = num
			num++
		}
		if num > n*n {
			break
		}
		for i := rh - 1; i >= rl+1; i-- {
			r[i][rl] = num
			num++
		}
		if num > n*n {
			break
		}
		rl++
		cl++
		rh--
		ch--
	}
	return r
}
