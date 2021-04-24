package main

func halvesAreAlike(s string) bool {
	first, second := 0, 0
	for i := 0; i < len(s)/2; i++ {
		if isVowels(s[i]) {
			first++
		}
		if isVowels(s[i+len(s)/2]) {
			second++
		}
	}
	return first == second
}

func isVowels(b byte) bool {
	switch b {
	case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
		return true
	default:
		return false
	}
}
