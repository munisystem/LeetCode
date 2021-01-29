package main

import "math"

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for i := 0; i < len(position); i++ {
		if position[i]%2 == 1 {
			odd++
		} else {
			even++
		}
	}
	return int(math.Min(float64(odd), float64(even)))
}
