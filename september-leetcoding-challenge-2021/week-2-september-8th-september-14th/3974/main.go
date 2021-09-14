package main

func reverseOnlyLetters(s string) string {
	ans := make([]byte, len(s), len(s))
	left, right := 0, len(s)-1
	for left <= right {
		if !isLetter(s[left]) {
			ans[left] = s[left]
			left++
			continue
		}
		if !isLetter(s[right]) {
			ans[right] = s[right]
			right--
			continue
		}
		if left <= right {
			ans[left], ans[right] = s[right], s[left]
			left++
			right--
		}
	}
	return string(ans)
}

func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}
