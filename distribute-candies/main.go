package main

func distributeCandies(candyType []int) int {
	m := map[int]struct{}{}
	types := 0
	for i := 0; i < len(candyType); i++ {
		if _, ok := m[candyType[i]]; ok {
			continue
		}
		m[candyType[i]] = struct{}{}
		types++
	}
	if types > len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return types
	}
}
