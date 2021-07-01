package main

func sumBase(n int, k int) int {
	var ans int
	for n > 0 {
		ans += n % k
		n /= k
	}
	return ans
}
