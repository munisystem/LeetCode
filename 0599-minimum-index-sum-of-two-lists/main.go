package main

import (
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i := 0; i < len(list1); i++ {
		m[list1[i]] = i
	}
	min := math.MaxInt32
	var ans []string
	for i := 0; i < len(list2); i++ {
		j, ok := m[list2[i]]
		if ok && i+j <= min {
			if i+j < min {
				min = i + j
				ans = []string{}
			}
			ans = append(ans, list2[i])
		}
	}
	return ans
}
