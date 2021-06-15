package main

import "math"

func maxCount(m int, n int, ops [][]int) int {
	for i := 0; i < len(ops); i++ {
		m = int(math.Min(float64(m), float64(ops[i][0])))
		n = int(math.Min(float64(n), float64(ops[i][1])))
	}
	return m * n
}
