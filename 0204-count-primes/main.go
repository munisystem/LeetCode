package main

func countPrimes(n int) int {
	m := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		m[i] = true
	}
	for i := 2; i*i < n; i++ {
		if m[i] == false {
			continue
		}
		for j := 2; i*j < n; j++ {
			m[i*j] = false
		}
	}
	r := 0
	for _, v := range m {
		if v == true {
			r++
		}
	}
	return r
}
