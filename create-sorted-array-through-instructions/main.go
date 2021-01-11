package main

import (
	"math"
)

const mod = 1000000007

func createSortedArray(instructions []int) int {
	bit := make([]int, 100001, 100001)
	cost := 0
	for i := 0; i < len(instructions); i++ {
		num := instructions[i]
		cost = (cost + int(math.Min(float64(get(bit, num-1)), float64(i-get(bit, num))))) % mod
		bit = update(bit, num)
	}
	return cost
}

func update(bit []int, num int) []int {
	i := num
	for i < len(bit) {
		bit[i]++
		i += i & -i
	}
	return bit
}

func get(bit []int, num int) int {
	res := 0
	i := num
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}
