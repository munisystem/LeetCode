package main

func numOfStrings(patterns []string, word string) int {
	var ans int
	for _, pattern := range patterns {
		for i := 0; i < len(word)-len(pattern)+1; i++ {
			var sub string
			if i+len(pattern) == len(word) {
				sub = word[i:]
			} else {
				sub = word[i : i+len(pattern)]
			}
			if pattern == sub {
				ans++
				break
			}
		}
	}
	return ans
}
