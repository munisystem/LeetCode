package main

import (
	"math"
)

func grayCode(n int) []int {
	r := []int{}
	for i := 0; i < n; i++ {
		if i == 0 {
			r = []int{0, 1}
			continue
		}
		for j := len(r) - 1; j >= 0; j-- {
			r = append(r, int(math.Pow(2, float64(i)))+r[j])
		}
	}
	return r
}
