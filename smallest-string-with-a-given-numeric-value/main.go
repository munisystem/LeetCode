package main

func getSmallestString(n int, k int) string {
	ans := make([]rune, n, n)
	for i := n; i >= 1; i-- {
		if i == 1 {
			ans[i-1] = rune(96 + k)
			break
		}
		for j := 26; j >= 1; j-- {
			if k-j >= 0 && k-j >= i-1 {
				ans[i-1] = rune(96 + j)
				k -= j
				break
			}
		}
	}
	return string(ans)
}
