package main

func countSegments(s string) int {
	var ans, seg int
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			seg++
			continue
		}
		if seg != 0 {
			ans++
			seg = 0
		}
	}
	if seg != 0 {
		ans++
	}
	return ans
}
