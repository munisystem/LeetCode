package main

func replaceDigits(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			ans = append(ans, s[i])
		} else {
			ans = append(ans, s[i-1]+s[i]-'0')
		}
	}
	return string(ans)
}
