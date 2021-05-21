package main

var m = map[byte]interface{}{}

func init() {
	vowels := "aiueoAIUEO"
	for i := 0; i < len(vowels); i++ {
		m[vowels[i]] = struct{}{}
	}
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	ans := []byte(s)
	for l < r {
		for l < r {
			if _, ok := m[ans[l]]; ok {
				break
			}
			l++
		}
		for l < r {
			if _, ok := m[ans[r]]; ok {
				break
			}
			r--
		}
		ans[l], ans[r] = s[r], s[l]
		l++
		r--
	}
	return string(ans)
}
