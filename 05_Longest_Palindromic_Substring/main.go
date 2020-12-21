package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	for i := range s {
		for j := i; j < len(s); j++ {
			buf := s[i : j+1]
			if len(buf) == 1 && len(buf) > len(res) {
				res = buf
				continue
			}
			isPalindrome := true
			for k := 0; k < len(buf)/2; k++ {
				if buf[k:k+1] != buf[len(buf)-1-k:len(buf)-k] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome && len(buf) > len(res) {
				res = buf
			}
		}
	}
	return res
}
