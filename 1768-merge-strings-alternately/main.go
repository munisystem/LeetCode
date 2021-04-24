package main

func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	i, j := 0, 0
	for i != len(word1) && j != len(word2) {
		ans = ans + string(word1[i]) + string(word2[j])
		i++
		j++
	}
	if i != len(word1) {
		ans = ans + word1[i:]
	} else if j != len(word2) {
		ans = ans + word2[j:]
	}
	return ans
}
