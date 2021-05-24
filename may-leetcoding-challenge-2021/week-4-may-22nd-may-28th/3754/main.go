package main

func toLowerCase(s string) string {
	chars := []byte(s)
	for i := 0; i < len(chars); i++ {
		code := int(chars[i])
		if code >= 65 && code <= 90 {
			chars[i] += 32
		}
	}
	return string(chars)
}
