package main

import "sort"

func maxNumberOfBalloons(text string) int {
	chars := make([]int, 5, 5)
	for _, r := range text {
		switch r {
		case 'b':
			chars[0]++
		case 'a':
			chars[1]++
		case 'l':
			chars[2]++
		case 'o':
			chars[3]++
		case 'n':
			chars[4]++
		}
	}
	chars[2] /= 2
	chars[3] /= 2

	sort.Ints(chars)

	return chars[0]
}
