package main

func checkIfPangram(sentence string) bool {
	chars := make([]int, 26, 26)
	for i := 0; i < len(sentence); i++ {
		chars[sentence[i]-'a']++
	}
	for i := 0; i < len(chars); i++ {
		if chars[i] == 0 {
			return false
		}
	}
	return true
}
