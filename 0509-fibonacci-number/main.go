package main

import (
	"math"
)

func fib(n int) int {
	golden := (1.0 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(golden, float64(n)) / math.Sqrt(5)))
}
