package main

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		odd, even := palindrome(s, i, i), palindrome(s, i, i+1)
		if len(ans) >= len(odd) && len(ans) >= len(even) {
			continue
		}
		if len(odd) > len(even) {
			ans = odd
		} else {
			ans = even
		}
	}
	return ans
}

func palindrome(s string, i, j int) string {
	str := ""
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		str = s[i : j+1]
		i--
		j++
	}
	return str
}
