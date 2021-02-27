package main

func countConsistentStrings(allowed string, words []string) int {
	m := map[int]interface{}{}
	for i := 0; i < len(allowed); i++ {
		m[int(allowed[i])-'a'] = struct{}{}
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		isConsistent := true
		for j := 0; j < len(words[i]); j++ {
			if _, ok := m[int(words[i][j])-'a']; !ok {
				isConsistent = false
				break
			}
		}
		if isConsistent {
			ans++
		}
	}
	return ans
}
