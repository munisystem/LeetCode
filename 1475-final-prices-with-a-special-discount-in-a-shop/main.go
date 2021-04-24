package main

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices), len(prices))
	for i := 0; i < len(prices); i++ {
		ans[i] = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				ans[i] -= prices[j]
				break
			}
		}
	}
	return ans
}
