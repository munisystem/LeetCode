package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[int]int{}
	for i := 0; i < len(s); i++ {
		m[int(s[i])-'a']++
		m[int(t[i])-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
