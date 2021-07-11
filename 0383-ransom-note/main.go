package main

func canConstruct(ransomNote string, magazine string) bool {
	chars := make([]int, 26, 26)
	for i := 0; i < len(magazine); i++ {
		chars[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if chars[ransomNote[i]-'a'] == 0 {
			return false
		}
		chars[ransomNote[i]-'a']--
	}
	return true
}
