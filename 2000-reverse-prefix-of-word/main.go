package main

func reversePrefix(word string, ch byte) string {
	ans := []byte(word)
	right, left := 0, 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			left = i
			break
		}
	}
	for right < left {
		ans[right], ans[left] = ans[left], ans[right]
		right++
		left--
	}
	return string(ans)
}
