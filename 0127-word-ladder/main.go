package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := map[string][]string{}
	for i := 0; i < len(wordList); i++ {
		word := wordList[i]
		for j := 0; j < len(beginWord); j++ {
			s := word[0:j] + "*" + word[j+1:]
			m[s] = append(m[s], word)
		}
	}
	queue := []string{}
	queue = append(queue, beginWord)
	p := map[string]interface{}{}
	r := 1
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return r
			}
			for j := 0; j < len(word); j++ {
				s := word[0:j] + "*" + word[j+1:]
				if matched, ok := m[s]; ok {
					for k := 0; k < len(matched); k++ {
						if _, ok := p[matched[k]]; ok {
							continue
						}
						queue = append(queue, matched[k])
						p[matched[k]] = nil
					}
				}
			}
		}
		r++
	}
	return 0
}
