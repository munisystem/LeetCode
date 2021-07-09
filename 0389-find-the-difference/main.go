package main

func findTheDifference(s string, t string) byte {
	var b int
	for i := 0; i < len(s); i++ {
		b += int(t[i])
		b -= int(s[i])
	}
	b += int(t[len(t)-1])
	return byte(b)
}
