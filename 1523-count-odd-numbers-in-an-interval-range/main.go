package main

func countOdds(low int, high int) int {
	var ans int
	ans = (high - low) / 2
	if low%2 == 1 || high%2 == 1 {
		ans++
	}
	return ans
}
