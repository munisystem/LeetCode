package main

func removeDuplicates(s string) string {
	var chars []byte
	for i := 0; i < len(s); i++ {
		if len(chars) == 0 || chars[len(chars)-1] != s[i] {
			chars = append(chars, s[i])
		} else {
			chars = chars[:len(chars)-1]
		}
	}
	return string(chars)
}
