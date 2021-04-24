package main

func findKthPositive(arr []int, k int) int {
	l, n := 0, len(arr)
	for l < n {
		m := l + (n-l)/2
		if arr[m]-m-1 >= k {
			n = m
		} else {
			l = m + 1
		}
	}
	return k + l
}
