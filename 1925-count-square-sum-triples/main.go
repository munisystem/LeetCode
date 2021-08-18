package main

func countTriples(n int) int {
	m := map[int]struct{}{}
	for i := 1; i <= n; i++ {
		m[i*i] = struct{}{}
	}
	var ans int
	for i := 1; i*i < n*n; i++ {
		for j := i + 1; i*i+j*j <= n*n; j++ {
			if _, ok := m[i*i+j*j]; ok {
				ans += 2
			}
		}
	}
	return ans
}
