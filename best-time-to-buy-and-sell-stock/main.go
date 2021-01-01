package main

import (
	"math"
)

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			max = int(math.Max(float64(max), float64(prices[j]-prices[i])))
		}
	}
	return max
}
