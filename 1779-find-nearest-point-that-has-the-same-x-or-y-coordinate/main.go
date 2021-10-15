package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	ans, tmp := -1, math.MaxInt32
	for i, p := range points {
		a, b := p[0], p[1]
		if a == x || b == y {
			if int(math.Abs(a - x)) {
				tmp = abs(a-x) + abs(b-y)
				ans = i
			}
		}
	}
	return ans
}
