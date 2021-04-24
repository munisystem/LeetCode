package main

func toLowerCase(str string) string {
	ans := make([]byte, len(str), len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			ans[i] = byte(int(str[i]) + 32)
		} else {
			ans[i] = str[i]
		}
	}
	return string(ans)
}
