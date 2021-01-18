package main

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1
	r := 0
	for i <= j {
		r++
		if people[i]+people[j] <= limit {
			i++
		}
		j--
	}
	return r
}
