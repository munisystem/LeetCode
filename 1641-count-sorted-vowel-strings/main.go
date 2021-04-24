package main

func countVowelStrings(n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r = r * ((5 + n - 1) - i) / (i + 1)
	}
	return r
}
