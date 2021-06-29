package main

var morsecodes []string = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	w := map[string]struct{}{}
	for i := 0; i < len(words); i++ {
		var s string
		for j := 0; j < len(words[i]); j++ {
			s = s + morsecodes[words[i][j]-'a']
		}
		w[s] = struct{}{}
	}
	return len(w)
}
