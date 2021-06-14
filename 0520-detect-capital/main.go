package main

func detectCapitalUse(word string) bool {
	var count int
	for i := 0; i < len(word); i++ {
		if 'Z'-int(word[i]) >= 0 {
			count++
		}
	}
	return count == 0 || count == len(word) || (count == 1 && 'Z'-int(word[0]) >= 0)
}
