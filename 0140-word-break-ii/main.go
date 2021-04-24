package main

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool, 0)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	return dfs(s, dict, map[string][]string{})
}

func dfs(s string, dict map[string]bool, cache map[string][]string) []string {
	if strs, ok := cache[s]; ok {
		return strs
	}
	if len(s) == 0 {
		return []string{""}
	}
	strs := []string{}
	for word := range dict {
		if len(word) <= len(s) && word == s[:len(word)] {
			for _, str := range dfs(s[len(word):], dict, cache) {
				if len(str) == 0 {
					strs = append(strs, word)
				} else {
					strs = append(strs, word+" "+str)
				}
			}
		}
	}
	cache[s] = strs
	return strs
}
