package main

type WordFilter struct {
	words []string
}

func Constructor(words []string) WordFilter {
	return WordFilter{words: words}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	for i := len(this.words) - 1; i >= 0; i-- {
		word := this.words[i]
		if word[0:len(prefix)] == prefix && word[len(word)-len(suffix):] == suffix {
			return i
		}
	}
	return -1
}
