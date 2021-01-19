package main

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]interface{}{}
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = struct{}{}
	}
	dp := make([]bool, len(s)+1, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dict[s[j:i]]; ok && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
