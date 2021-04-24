package main

func longestPalindrome(s string) int {
	m := make(map[string]int, 0)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	r := 0
	odds := 0
	for _, v := range m {
		r += 2 * (v / 2)
		if v%2 == 1 {
			odds++
		}
	}
	if odds > 0 {
		r++
	}
	return r
}
