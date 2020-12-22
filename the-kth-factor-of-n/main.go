package main

import "sort"

func kthFactor(n int, k int) int {
	f := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			f = append(f, i)
			if n/i != i {
				f = append(f, n/i)
			}
		}
	}
	sort.Ints(f)
	if len(f) >= k {
		return f[k-1]
	}
	return -1
}
