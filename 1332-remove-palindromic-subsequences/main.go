package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("abba"))
	fmt.Println(removePalindromeSub(""))
	fmt.Println(removePalindromeSub("bbaabaaa"))
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if palindrome(s) {
		return 1
	} else {
		return 2
	}
}

func palindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
