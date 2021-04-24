package main

import (
	"math"
)

func maxProfit(prices []int, fee int) int {
	p, h := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		p = int(math.Max(float64(p), float64(h+prices[i]-fee)))
		h = int(math.Max(float64(h), float64(p-prices[i])))
	}
	return p
}
