package main

func sumZero(n int) []int {
	ans := make([]int, n, n)
	l, r := 0, n-1
	num := n / 2
	for l <= r {
		ans[l], ans[r] = -num, num
		num--
		l++
		r--
	}
	return ans
}
