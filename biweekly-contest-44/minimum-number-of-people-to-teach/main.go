package main

import (
	"math"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := make(map[int]map[int]bool, len(languages))
	for i := 1; i <= len(languages); i++ {
		m[i] = make(map[int]bool, n)
		for j := 0; j < len(languages[i-1]); j++ {
			m[i][languages[i-1][j]] = true
		}
	}
	cannotCommunicates := [][]int{}
	for j := 0; j < len(friendships); j++ {
		a, b := friendships[j][0], friendships[j][1]
		canCommunicate := false
		for k := range m[a] {
			if m[b][k] == true {
				canCommunicate = true
				break
			}
		}
		if canCommunicate {
			continue
		}
		cannotCommunicates = append(cannotCommunicates, []int{a, b})
	}
	if len(cannotCommunicates) == 0 {
		return 0
	}
	teaches := make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		alreadyTeaches := map[int]bool{}
		for j := 0; j < len(cannotCommunicates); j++ {
			a, b := cannotCommunicates[j][0], cannotCommunicates[j][1]
			if !m[a][i] && !alreadyTeaches[a] {
				teaches[i]++
				alreadyTeaches[a] = true
			}
			if !m[b][i] && !alreadyTeaches[b] {
				teaches[i]++
				alreadyTeaches[b] = true
			}
		}
	}
	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		ans = int(math.Min(float64(ans), float64(teaches[i])))
	}
	return ans
}
